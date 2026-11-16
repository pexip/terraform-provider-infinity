/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"sort"
	"strings"
)

// Scope values for upgrade mode.
const (
	scopeNew = "new" // only changes introduced in the latest schema version
	scopeAll = "all" // also pre-existing gaps, excluding validator-only checks
)

// validatorOnlyIssues are excluded from upgrade plans unless they were
// introduced in the latest version (e.g. a max length that changed).
var validatorOnlyIssues = map[string]bool{
	IssueMaxLength:    true,
	IssueRange:        true,
	IssueNotSensitive: true,
}

// issueActions maps finding issues to the change to make in the provider.
var issueActions = map[string]string{
	IssueMissingAttr:  "Add attribute",
	IssueExtraAttr:    "Remove attribute (not in API)",
	IssueTypeMismatch: "Change attribute type",
	IssueRequired:     "Change required/optional",
	IssueReadOnly:     "Make attribute computed-only",
	IssueDefault:      "Update default",
	IssueChoices:      "Update OneOf validator",
	IssueMaxLength:    "Update length validator",
	IssueRange:        "Update range validator",
	IssueDeprecated:   "Deprecate attribute",
	IssueNotSensitive: "Mark attribute sensitive",
}

const (
	actionDescription = "Update description"
	actionReview      = "Review API change"
)

// UpgradeChange is a single change to make to a provider resource.
type UpgradeChange struct {
	Action      string         `json:"action"`
	Field       string         `json:"field"`
	Issue       string         `json:"issue,omitempty"`
	Detail      string         `json:"detail"`
	Tag         string         `json:"tag"`
	Baseline    string         `json:"baseline,omitempty"`
	API         map[string]any `json:"api,omitempty"`
	BaselineAPI map[string]any `json:"baseline_api,omitempty"`
}

// FileRef is a file or directory related to a resource.
type FileRef struct {
	Path   string `json:"path"`
	Exists bool   `json:"exists"`
}

// UpgradeResource is the set of changes for one provider resource, sized for
// a single PR.
type UpgradeResource struct {
	TypeName         string          `json:"type_name"`
	Schema           string          `json:"schema"`
	Files            []FileRef       `json:"files"`
	SDKModel         string          `json:"sdk_model,omitempty"`
	SDKMissingFields []string        `json:"sdk_missing_fields,omitempty"`
	Changes          []UpgradeChange `json:"changes"`
}

// NewResource is an API resource with no provider resource yet.
type NewResource struct {
	Schema   string            `json:"schema"`
	Tag      string            `json:"tag"`
	SDKModel string            `json:"sdk_model,omitempty"`
	Fields   map[string]string `json:"fields"`
}

// UpgradePlan is the upgrade-mode output.
type UpgradePlan struct {
	Baseline     string            `json:"baseline"`
	Latest       string            `json:"latest"`
	Scope        string            `json:"scope"`
	Resources    []UpgradeResource `json:"resources"`
	NewResources []NewResource     `json:"new_resources"`
}

// PlanPaths locates the provider repo and SDK checkout for file references.
type PlanPaths struct {
	RepoDir string
	SDKDir  string
}

func inScope(tag, issue, scope string) bool {
	if tag != tagPreExisting {
		return true
	}
	return scope == scopeAll && !validatorOnlyIssues[issue]
}

// isMissingValidator reports validator-only findings where the provider has
// no validator at all; adding validators is out of scope for upgrades.
func isMissingValidator(f Finding) bool {
	return validatorOnlyIssues[f.Issue] && strings.Contains(f.Detail, noValidatorMarker)
}

// buildUpgradePlan turns a report into per-resource change lists.
func buildUpgradePlan(rep *Report, resources []*TFResource, latest, baseline *SchemaSet, scope string, paths PlanPaths) *UpgradePlan {
	plan := &UpgradePlan{Baseline: rep.Baseline, Latest: rep.Latest, Scope: scope}
	tfByType := make(map[string]*TFResource, len(resources))
	for _, r := range resources {
		tfByType[r.TypeName] = r
	}
	byType := map[string]*UpgradeResource{}
	get := func(typeName, schemaName string) *UpgradeResource {
		if r, ok := byType[typeName]; ok {
			return r
		}
		r := &UpgradeResource{TypeName: typeName, Schema: schemaName}
		byType[typeName] = r
		return r
	}
	fieldDef := func(set *SchemaSet, schemaName, field string) map[string]any {
		if s, ok := set.Schemas[schemaName]; ok {
			if f, ok := s.Fields[field]; ok {
				return f.Raw
			}
		}
		return nil
	}

	// Fields with a version-introduced finding; used to avoid duplicating
	// them as "review" items below.
	covered := map[string]bool{}
	for _, rr := range rep.Resources {
		for _, f := range rr.Findings {
			if !inScope(f.Tag, f.Issue, scope) || isMissingValidator(f) || isDeprecatedMissing(latest, rr.Schema, f) {
				continue
			}
			r := get(rr.TypeName, rr.Schema)
			r.Changes = append(r.Changes, UpgradeChange{
				Action: issueActions[f.Issue], Field: f.Field, Issue: f.Issue,
				Detail: f.Detail, Tag: f.Tag, Baseline: f.Baseline,
				API:         fieldDef(latest, rr.Schema, f.Field),
				BaselineAPI: fieldDef(baseline, rr.Schema, f.Field),
			})
			if f.Tag != tagPreExisting {
				covered[rr.TypeName+"|"+f.Field] = true
			}
		}
	}

	// API definition changes between versions that produced no finding, e.g.
	// help_text updates or unique changes. Only fields the provider exposes
	// are relevant, and keys already verified by the per-field checks are
	// dropped since no finding means the provider already matches.
	for _, c := range rep.Changes {
		if c.Resource == "" || covered[c.Resource+"|"+c.Field] {
			continue
		}
		if tf, ok := tfByType[c.Resource]; !ok || tf.Attrs[c.Field] == nil {
			continue
		}
		base := fieldDef(baseline, c.Schema, c.Field)
		api := fieldDef(latest, c.Schema, c.Field)
		keys := unverifiedKeys(c.Keys, api)
		if len(keys) == 0 {
			continue
		}
		r := get(c.Resource, c.Schema)
		tag := fmt.Sprintf(tagNewFormat, latest.Version)
		if slices.Contains(keys, "help_text") {
			r.Changes = append(r.Changes, UpgradeChange{
				Action: actionDescription, Field: c.Field, Detail: "help_text changed",
				Tag: tag, API: api, BaselineAPI: base,
			})
			keys = slices.DeleteFunc(keys, func(k string) bool { return k == "help_text" })
		}
		if len(keys) > 0 {
			details := make([]string, 0, len(keys))
			for _, k := range keys {
				bv, bok := base[k]
				lv, lok := api[k]
				details = append(details, fmt.Sprintf("%s: %s → %s", k, rawValue(bv, bok), rawValue(lv, lok)))
			}
			r.Changes = append(r.Changes, UpgradeChange{
				Action: actionReview, Field: c.Field, Detail: strings.Join(details, "; "),
				Tag: tag, API: api, BaselineAPI: base,
			})
		}
	}

	for _, typeName := range sortedKeys(byType) {
		r := byType[typeName]
		r.Files = resourceFiles(paths.RepoDir, strings.TrimPrefix(typeName, providerTypeName+"_"))
		r.SDKModel, r.SDKMissingFields = sdkModelGaps(paths.SDKDir, r)
		plan.Resources = append(plan.Resources, *r)
	}

	for _, f := range rep.Unmapped {
		if f.Issue != IssueNoProviderRes || !inScope(f.Tag, f.Issue, scope) {
			continue
		}
		nr := NewResource{Schema: f.Schema, Tag: f.Tag, Fields: map[string]string{}}
		nr.SDKModel = sdkModelPath(paths.SDKDir, f.Schema)
		for name, field := range latest.Schemas[f.Schema].Fields {
			nr.Fields[name] = field.describe()
		}
		plan.NewResources = append(plan.NewResources, nr)
	}
	return plan
}

// verifiedKeys are API field keys whose changes are fully checked by
// compareField, so a change without a finding needs no provider update.
var verifiedKeys = map[string]bool{
	"type": true, "related_type": true, "readonly": true, "valid_choices": true,
	"blank": true, "nullable": true,
	"deprecated": true, "deprecated_use_field": true, "deprecated_use_resource": true,
}

// unverifiedKeys filters changed keys down to those needing manual review.
// A default is only verified when the latest default is a fixed non-null
// value; removing a default is not detected by compareField.
func unverifiedKeys(keys []string, api map[string]any) []string {
	var out []string
	for _, k := range keys {
		if verifiedKeys[k] {
			continue
		}
		if k == "default" {
			if d, ok := api["default"]; ok && d != nil && !isDynamic(d) {
				continue
			}
		}
		out = append(out, k)
	}
	return out
}

// isDeprecatedMissing reports a missing attribute whose API field is
// deprecated; the provider should not add it.
func isDeprecatedMissing(latest *SchemaSet, schemaName string, f Finding) bool {
	if f.Issue != IssueMissingAttr {
		return false
	}
	s, ok := latest.Schemas[schemaName]
	if !ok {
		return false
	}
	field, ok := s.Fields[f.Field]
	return ok && field.Deprecated
}

var testdataSuffixPattern = regexp.MustCompile(`^(min|full|basic)(_[a-z_]+)?$`)

// resourceFiles lists the provider files typically touched when updating a resource.
func resourceFiles(repoDir, suffix string) []FileRef {
	paths := []string{
		filepath.Join("internal", "provider", "resource_infinity_"+suffix+".go"),
		filepath.Join("internal", "provider", "resource_infinity_"+suffix+"_test.go"),
		filepath.Join("internal", "provider", "resource_infinity_"+suffix+"_integration_test.go"),
	}
	prefix := "resource_infinity_" + suffix + "_"
	if entries, err := os.ReadDir(filepath.Join(repoDir, "testdata")); err == nil {
		for _, e := range entries {
			// Skip other resources sharing the prefix, e.g. conference_alias for conference.
			if e.IsDir() && strings.HasPrefix(e.Name(), prefix) && testdataSuffixPattern.MatchString(strings.TrimPrefix(e.Name(), prefix)) {
				paths = append(paths, filepath.Join("testdata", e.Name()))
			}
		}
	}
	paths = append(paths,
		filepath.Join("templates", "resources", suffix+".md.tmpl"),
		filepath.Join("examples", "resources", "infinity_"+suffix),
		filepath.Join("docs", "resources", suffix+".md"),
	)
	out := make([]FileRef, 0, len(paths))
	for _, p := range paths {
		_, err := os.Stat(filepath.Join(repoDir, p))
		out = append(out, FileRef{Path: p, Exists: err == nil})
	}
	return out
}

func sdkModelPath(sdkDir string, names ...string) string {
	for _, n := range names {
		p := filepath.Join(sdkDir, "config", n+"_model.go")
		if _, err := os.Stat(p); err == nil {
			return p
		}
	}
	return ""
}

var jsonTagPattern = regexp.MustCompile("json:\"([a-z0-9_]+)")

// sdkModelGaps returns the SDK model file for a resource and the fields being
// added or changed that the model does not declare.
func sdkModelGaps(sdkDir string, r *UpgradeResource) (model string, missing []string) {
	model = sdkModelPath(sdkDir, r.Schema, strings.TrimPrefix(r.TypeName, providerTypeName+"_"))
	if model == "" {
		return "", nil
	}
	data, err := os.ReadFile(filepath.Clean(model))
	if err != nil {
		return model, nil
	}
	tags := map[string]bool{}
	for _, m := range jsonTagPattern.FindAllSubmatch(data, -1) {
		tags[string(m[1])] = true
	}
	seen := map[string]bool{}
	for _, c := range r.Changes {
		if c.Issue != IssueMissingAttr && c.Issue != IssueTypeMismatch {
			continue
		}
		if !tags[c.Field] && !seen[c.Field] {
			seen[c.Field] = true
			missing = append(missing, c.Field)
		}
	}
	sort.Strings(missing)
	return model, missing
}

const upgradeInstructions = `## How to apply

Create one PR per resource section below. For each resource:

1. If **SDK model missing fields** is listed, add those fields to the go-infinity-sdk model and its create/update request structs first, then bump the SDK dependency in this repo.
2. Update the resource file: schema attribute (type, Required/Optional/Computed, Default, validators, description from help_text), the resource model struct, and the create/read/update mappings.
3. Update the unit test, integration test and testdata min/full configs following the test pattern in .claude/CLAUDE.md.
4. Update the docs template and examples if they reference changed attributes, then regenerate docs with ` + "`tfplugindocs generate --provider-name infinity`" + `.
5. Run ` + "`go test -tags unit ./internal/provider/ -run <ResourceTest>`" + ` and ` + "`make lint`" + `.

The API definitions under each change are the raw field definitions from the latest schema (and the baseline, where the field existed).
`

func writeUpgradeMarkdown(w io.Writer, plan *UpgradePlan) {
	writeUpgradeHeader(w, plan)
	fmt.Fprint(w, upgradeInstructions)
	fmt.Fprintln(w)
	writeUpgradeIndex(w, plan, false)
	for i := range plan.Resources {
		writeUpgradeResource(w, plan, &plan.Resources[i])
	}
	writeNewResources(w, plan)
}

func writeUpgradeHeader(w io.Writer, plan *UpgradePlan) {
	fmt.Fprintf(w, "# Provider upgrade plan: %s → %s\n\n", plan.Baseline, plan.Latest)
	scope := "changes introduced in " + plan.Latest
	if plan.Scope == scopeAll {
		scope += ", plus pre-existing gaps (excluding validator-only checks)"
	}
	fmt.Fprintf(w, "Scope: %s. Summary: %d resources to update, %d new resources.\n\n",
		scope, len(plan.Resources), len(plan.NewResources))
}

func writeUpgradeIndex(w io.Writer, plan *UpgradePlan, links bool) {
	if len(plan.Resources) == 0 {
		fmt.Fprintln(w, "No provider resources need updating.")
		fmt.Fprintln(w)
		return
	}
	fmt.Fprintln(w, "## Resources")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "| Resource | Changes | SDK fields missing |")
	fmt.Fprintln(w, "|---|---|---|")
	for _, r := range plan.Resources {
		name := r.TypeName
		if links {
			name = fmt.Sprintf("[%s](%s.md)", r.TypeName, r.TypeName)
		}
		fmt.Fprintf(w, "| %s | %d | %s |\n", name, len(r.Changes), strings.Join(r.SDKMissingFields, ", "))
	}
	fmt.Fprintln(w)
}

func writeUpgradeResource(w io.Writer, plan *UpgradePlan, r *UpgradeResource) {
	fmt.Fprintf(w, "## %s\n\n", r.TypeName)
	fmt.Fprintf(w, "- **API schema:** `%s.json` (%s)\n", r.Schema, plan.Latest)
	if r.SDKModel != "" {
		fmt.Fprintf(w, "- **SDK model:** `%s`\n", r.SDKModel)
	}
	if len(r.SDKMissingFields) > 0 {
		fmt.Fprintf(w, "- **SDK model missing fields:** %s\n", strings.Join(r.SDKMissingFields, ", "))
	}
	fmt.Fprintln(w, "- **Files:**")
	for _, f := range r.Files {
		suffix := ""
		if !f.Exists {
			suffix = " (does not exist)"
		}
		fmt.Fprintf(w, "  - `%s`%s\n", f.Path, suffix)
	}
	fmt.Fprintln(w)

	fmt.Fprintln(w, "| # | Action | Field | Detail | Tag |")
	fmt.Fprintln(w, "|---|---|---|---|---|")
	for i, c := range r.Changes {
		tag := c.Tag
		if c.Baseline != "" {
			tag += " (was: " + mdEscape(c.Baseline) + ")"
		}
		fmt.Fprintf(w, "| %d | %s | `%s` | %s | %s |\n", i+1, c.Action, c.Field, mdEscape(c.Detail), tag)
	}
	fmt.Fprintln(w)

	fmt.Fprintln(w, "<details><summary>API definitions</summary>")
	fmt.Fprintln(w)
	for i, c := range r.Changes {
		if c.API == nil && c.BaselineAPI == nil {
			continue
		}
		fmt.Fprintf(w, "%d. `%s`\n\n", i+1, c.Field)
		if c.API != nil {
			fmt.Fprintf(w, "   %s:\n   ```json\n   %s\n   ```\n", plan.Latest, compactJSON(c.API))
		}
		if c.BaselineAPI != nil && !jsonEqual(c.API, c.BaselineAPI) {
			fmt.Fprintf(w, "   %s:\n   ```json\n   %s\n   ```\n", plan.Baseline, compactJSON(c.BaselineAPI))
		}
		fmt.Fprintln(w)
	}
	fmt.Fprintln(w, "</details>")
	fmt.Fprintln(w)
}

func writeNewResources(w io.Writer, plan *UpgradePlan) {
	if len(plan.NewResources) == 0 {
		return
	}
	fmt.Fprintln(w, "## New resources to create")
	fmt.Fprintln(w)
	for _, nr := range plan.NewResources {
		fmt.Fprintf(w, "### %s_%s (`%s.json`, %s)\n\n", providerTypeName, nr.Schema, nr.Schema, nr.Tag)
		if nr.SDKModel != "" {
			fmt.Fprintf(w, "SDK model: `%s`\n\n", nr.SDKModel)
		} else {
			fmt.Fprintln(w, "SDK model: none, add SDK support first.")
			fmt.Fprintln(w)
		}
		fmt.Fprintln(w, "| Field | Definition |")
		fmt.Fprintln(w, "|---|---|")
		for _, name := range sortedKeys(nr.Fields) {
			fmt.Fprintf(w, "| `%s` | %s |\n", name, mdEscape(nr.Fields[name]))
		}
		fmt.Fprintln(w)
	}
}

// writeUpgradeDir writes index.md plus one file per resource, so each PR can
// be driven from a single file.
func writeUpgradeDir(dir string, plan *UpgradePlan) error {
	if err := os.MkdirAll(dir, 0o750); err != nil {
		return err
	}
	write := func(name string, fn func(io.Writer)) error {
		f, err := os.Create(filepath.Join(dir, name))
		if err != nil {
			return err
		}
		fn(f)
		return f.Close()
	}
	err := write("index.md", func(w io.Writer) {
		writeUpgradeHeader(w, plan)
		writeUpgradeIndex(w, plan, true)
		writeNewResources(w, plan)
	})
	if err != nil {
		return err
	}
	for i := range plan.Resources {
		r := &plan.Resources[i]
		err := write(r.TypeName+".md", func(w io.Writer) {
			fmt.Fprintf(w, "# Upgrade %s: %s → %s\n\n", r.TypeName, plan.Baseline, plan.Latest)
			fmt.Fprint(w, upgradeInstructions)
			fmt.Fprintln(w)
			writeUpgradeResource(w, plan, r)
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func compactJSON(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprint(v)
	}
	return string(b)
}

func jsonEqual(a, b any) bool {
	return compactJSON(a) == compactJSON(b)
}
