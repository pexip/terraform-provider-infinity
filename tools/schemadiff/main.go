/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

// Command schemadiff compares the provider's resource schemas against the
// Infinity API configuration schemas shipped in go-infinity-sdk, and reports
// the provider changes needed for the latest schema version. Each finding is
// tagged as new in the latest version or pre-existing, using the schema
// version the provider currently depends on as the baseline.
//
// Modes:
//
//   - report (default): every difference between the provider and the latest
//     schema, including validator checks.
//   - upgrade: per-resource change lists for moving the provider to the latest
//     schema version, with file references and API definitions, intended for
//     creating one PR per resource. Use -out-dir to write one file per resource.
//
// Usage:
//
//	go run ./tools/schemadiff [-mode report|upgrade] [-latest config-v43] [-baseline config-v42]
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

const (
	defaultSDKDir        = "../go-infinity-sdk"
	defaultSchemaRootDir = defaultSDKDir + "/schema"
)

var sdkModulePattern = regexp.MustCompile(`github\.com/pexip/go-infinity-sdk/v(\d+)\s`)

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "schemadiff:", err)
		os.Exit(1)
	}
}

func run(args []string, w io.Writer) error {
	fs := flag.NewFlagSet("schemadiff", flag.ContinueOnError)
	schemaDir := fs.String("schema-dir", defaultSchemaRootDir, "directory containing config-vNN schema directories")
	latestName := fs.String("latest", "", "latest schema directory name (default: highest version found)")
	baselineName := fs.String("baseline", "", "baseline schema directory name (default: SDK version in go.mod)")
	filter := fs.String("resource", "", "only compare this resource (provider suffix, e.g. conference)")
	format := fs.String("format", "markdown", "output format: markdown or json")
	includeReadOnly := fs.Bool("include-readonly", false, "report read-only API fields missing from the provider")
	failOnFindings := fs.Bool("fail-on-findings", false, "exit non-zero if any findings are reported")
	mode := fs.String("mode", "report", "report: all differences; upgrade: per-resource changes for the latest version")
	scope := fs.String("scope", scopeNew, "upgrade mode: new (changes in the latest version) or all (also pre-existing gaps)")
	outDir := fs.String("out-dir", "", "upgrade mode: write index.md and one markdown file per resource to this directory")
	sdkDir := fs.String("sdk-dir", defaultSDKDir, "go-infinity-sdk checkout, used to check SDK models in upgrade mode")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *mode != "report" && *mode != "upgrade" {
		return fmt.Errorf("unknown mode %q", *mode)
	}
	if *scope != scopeNew && *scope != scopeAll {
		return fmt.Errorf("unknown scope %q", *scope)
	}

	dirs, err := findSchemaDirs(*schemaDir)
	if err != nil {
		return err
	}
	latestDir, err := resolveDir(*schemaDir, *latestName, dirs, latestVersion(dirs))
	if err != nil {
		return fmt.Errorf("latest: %w", err)
	}
	baseVersion, err := goModSDKVersion("go.mod")
	if err != nil && *baselineName == "" {
		return fmt.Errorf("baseline: %w (use -baseline)", err)
	}
	baselineDir, err := resolveDir(*schemaDir, *baselineName, dirs, baseVersion)
	if err != nil {
		return fmt.Errorf("baseline: %w", err)
	}

	latest, err := loadSchemaSet(latestDir)
	if err != nil {
		return err
	}
	baseline, err := loadSchemaSet(baselineDir)
	if err != nil {
		return err
	}

	resources := loadProviderResources(context.Background())
	rep := buildReport(resources, latest, baseline, Options{IncludeReadOnly: *includeReadOnly}, *filter)

	if *mode == "upgrade" {
		plan := buildUpgradePlan(rep, resources, latest, baseline, *scope, PlanPaths{RepoDir: ".", SDKDir: *sdkDir})
		if err := writeUpgrade(w, plan, *format, *outDir); err != nil {
			return err
		}
		if *failOnFindings && (len(plan.Resources) > 0 || len(plan.NewResources) > 0) {
			return fmt.Errorf("%d resources need updating", len(plan.Resources))
		}
		return nil
	}

	switch *format {
	case "json":
		if err := writeJSON(w, rep); err != nil {
			return err
		}
	case "markdown":
		writeMarkdown(w, rep)
	default:
		return fmt.Errorf("unknown format %q", *format)
	}

	if *failOnFindings && (len(rep.Resources) > 0 || len(rep.Unmapped) > 0) {
		return fmt.Errorf("%d resources with findings", len(rep.Resources))
	}
	return nil
}

func writeUpgrade(w io.Writer, plan *UpgradePlan, format, outDir string) error {
	if outDir != "" {
		if err := writeUpgradeDir(outDir, plan); err != nil {
			return err
		}
		fmt.Fprintf(w, "Wrote %d resource files and index.md to %s\n", len(plan.Resources), outDir)
		return nil
	}
	switch format {
	case "json":
		return writeJSON(w, plan)
	case "markdown":
		writeUpgradeMarkdown(w, plan)
		return nil
	default:
		return fmt.Errorf("unknown format %q", format)
	}
}

func writeJSON(w io.Writer, v any) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}

func latestVersion(dirs map[int]string) int {
	v := -1
	for k := range dirs {
		v = max(v, k)
	}
	return v
}

// resolveDir returns the named schema directory, or the one for version if no
// name was given.
func resolveDir(root, name string, dirs map[int]string, version int) (string, error) {
	if name != "" {
		return filepath.Join(root, name), nil
	}
	dir, ok := dirs[version]
	if !ok {
		return "", fmt.Errorf("no schema directory for v%d in %s", version, root)
	}
	return dir, nil
}

// goModSDKVersion returns the major version of go-infinity-sdk required in go.mod.
func goModSDKVersion(path string) (int, error) {
	data, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return 0, err
	}
	m := sdkModulePattern.FindSubmatch(data)
	if m == nil {
		return 0, fmt.Errorf("go-infinity-sdk not found in %s", path)
	}
	return strconv.Atoi(string(m[1]))
}

func writeMarkdown(w io.Writer, rep *Report) {
	var newCount, preCount int
	for _, r := range rep.Resources {
		for _, f := range r.Findings {
			if f.Tag == tagPreExisting {
				preCount++
			} else {
				newCount++
			}
		}
	}

	fmt.Fprintf(w, "# Schema diff: provider vs SDK %s (baseline %s)\n\n", rep.Latest, rep.Baseline)
	fmt.Fprintf(w, "Summary: %d resources with findings, %d new, %d pre-existing, %d unmapped schemas/resources\n\n",
		len(rep.Resources), newCount, preCount, len(rep.Unmapped))

	if len(rep.Unmapped) > 0 {
		fmt.Fprintln(w, "## Unmapped resources")
		fmt.Fprintln(w)
		fmt.Fprintln(w, "| Resource | Issue | Detail | Tag |")
		fmt.Fprintln(w, "|---|---|---|---|")
		unmapped := slices.Clone(rep.Unmapped)
		// Show new-in-latest entries first.
		slices.SortStableFunc(unmapped, func(a, b Finding) int {
			return boolRank(a.Tag == tagPreExisting) - boolRank(b.Tag == tagPreExisting)
		})
		for _, f := range unmapped {
			fmt.Fprintf(w, "| %s | %s | %s | %s |\n", f.Resource, f.Issue, mdEscape(f.Detail), f.Tag)
		}
		fmt.Fprintln(w)
	}

	for _, r := range rep.Resources {
		fmt.Fprintf(w, "## %s (%s.json)\n\n", r.TypeName, r.Schema)
		fmt.Fprintln(w, "| Field | Issue | Detail | Tag |")
		fmt.Fprintln(w, "|---|---|---|---|")
		for _, f := range r.Findings {
			tag := f.Tag
			if f.Baseline != "" {
				tag += " (was: " + mdEscape(f.Baseline) + ")"
			}
			fmt.Fprintf(w, "| %s | %s | %s | %s |\n", f.Field, f.Issue, mdEscape(f.Detail), tag)
		}
		fmt.Fprintln(w)
	}

	if len(rep.Changes) > 0 {
		fmt.Fprintf(w, "## API field changes %s → %s\n\n", rep.Baseline, rep.Latest)
		fmt.Fprintln(w, "Fields present in both versions whose definition changed (help_text changes may need doc updates).")
		fmt.Fprintln(w)
		fmt.Fprintln(w, "| Schema | Field | Changed | Detail |")
		fmt.Fprintln(w, "|---|---|---|---|")
		for _, c := range rep.Changes {
			fmt.Fprintf(w, "| %s | %s | %s | %s |\n", c.Schema, c.Field, strings.Join(c.Keys, ", "), mdEscape(c.Detail))
		}
		fmt.Fprintln(w)
	}
}

func boolRank(b bool) int {
	if b {
		return 1
	}
	return 0
}

func mdEscape(s string) string {
	return strings.ReplaceAll(s, "|", `\|`)
}
