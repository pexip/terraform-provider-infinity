//go:build unit

/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// upgradeFixture returns baseline/latest schema sets where the latest version
// adds a field, changes choices, help text, a default, and deprecates fields.
func upgradeFixture(t *testing.T) (*Report, []*TFResource, *SchemaSet, *SchemaSet) {
	t.Helper()
	parse := func(raw map[string]map[string]any) *APISchema {
		s := &APISchema{Name: "widget", Fields: map[string]*APIField{}}
		for name, def := range raw {
			f := &APIField{Raw: def}
			f.Type, _ = def["type"].(string)
			f.Blank, _ = def["blank"].(bool)
			f.Deprecated, _ = def["deprecated"].(bool)
			f.HelpText, _ = def["help_text"].(string)
			f.Default, f.HasDefault = def["default"]
			if c, ok := def["valid_choices"].([]any); ok {
				f.ValidChoices = c
			}
			s.Fields[name] = f
		}
		return s
	}
	common := func() map[string]map[string]any {
		return map[string]map[string]any{
			"name":      {"type": "string", "blank": false, "help_text": "The name. Maximum length: 250 characters."},
			"mode":      {"type": "string", "blank": true, "default": "a", "valid_choices": []any{"a", "b", "old"}},
			"template":  {"type": "string", "blank": true, "default": "x", "help_text": "old help"},
			"enabled":   {"type": "boolean", "blank": true, "default": false, "help_text": "Enabled."},
			"password":  {"type": "string", "blank": true, "default": "", "help_text": "Maximum length: 100 characters."},
			"legacy":    {"type": "boolean", "blank": true, "default": false},
			"untracked": {"type": "string", "blank": true, "help_text": "old"},
			"leftover":  {"type": "string", "blank": true},
		}
	}
	base, latest := common(), common()
	latest["new_field"] = map[string]any{"type": "boolean", "blank": true, "default": true}
	latest["mode"]["valid_choices"] = []any{"a", "b", "c"}
	latest["template"]["help_text"] = "new help"
	latest["enabled"]["default"] = nil
	latest["password"]["help_text"] = "Maximum length: 200 characters."
	latest["legacy"]["deprecated"] = true
	latest["untracked"]["help_text"] = "new"
	latest["leftover"]["deprecated"] = true

	baseSet := &SchemaSet{Dir: "config-v42", Version: 42, Schemas: map[string]*APISchema{"widget": parse(base)}}
	latestSet := &SchemaSet{Dir: "config-v43", Version: 43, Schemas: map[string]*APISchema{
		"widget": parse(latest),
		"gadget": {Name: "gadget", Fields: map[string]*APIField{"name": apiField("string")}},
	}}
	resources := []*TFResource{testResource()}
	rep := buildReport(resources, latestSet, baseSet, Options{}, "")
	return rep, resources, latestSet, baseSet
}

func changesByField(plan *UpgradePlan) map[string][]string {
	out := map[string][]string{}
	for _, r := range plan.Resources {
		for _, c := range r.Changes {
			out[c.Field] = append(out[c.Field], c.Action+": "+c.Detail)
		}
	}
	return out
}

func TestBuildUpgradePlanNewScope(t *testing.T) {
	rep, resources, latest, baseline := upgradeFixture(t)
	plan := buildUpgradePlan(rep, resources, latest, baseline, scopeNew, PlanPaths{RepoDir: t.TempDir(), SDKDir: t.TempDir()})

	require.Len(t, plan.Resources, 1)
	got := changesByField(plan)
	assert.Equal(t, []string{"Add attribute: boolean, optional, default true"}, got["new_field"])
	assert.Equal(t, []string{"Update OneOf validator: add c; remove old"}, got["mode"])
	assert.Equal(t, []string{"Update description: help_text changed"}, got["template"])
	assert.Equal(t, []string{"Review API change: default: false → null"}, got["enabled"])
	assert.Equal(t, []string{"Deprecate attribute: api field is deprecated"}, got["leftover"])

	// The max length change has no length item because the provider has no
	// validator, but the help text change still updates the description.
	assert.Equal(t, []string{"Update description: help_text changed"}, got["password"])

	// Pre-existing gaps, deprecated fields the provider never had, and changes
	// to fields the provider doesn't expose are skipped.
	for _, field := range []string{"name", "legacy", "untracked", "count"} {
		assert.Empty(t, got[field], field)
	}

	require.Len(t, plan.NewResources, 1)
	assert.Equal(t, "gadget", plan.NewResources[0].Schema)
	assert.Equal(t, "new in v43", plan.NewResources[0].Tag)
}

func TestBuildUpgradePlanAllScope(t *testing.T) {
	rep, resources, latest, baseline := upgradeFixture(t)
	plan := buildUpgradePlan(rep, resources, latest, baseline, scopeAll, PlanPaths{RepoDir: t.TempDir(), SDKDir: t.TempDir()})
	got := changesByField(plan)

	// Pre-existing gaps are included, validator-only checks are not.
	assert.Equal(t, []string{"Remove attribute (not in API): string attribute not in API schema (removed or renamed?)"}, got["count"])
	assert.Equal(t, []string{"Update description: help_text changed"}, got["password"])
	assert.Equal(t, "all", plan.Scope)
}

func TestResourceFiles(t *testing.T) {
	repo := t.TempDir()
	for _, d := range []string{
		"testdata/resource_infinity_conference_min",
		"testdata/resource_infinity_conference_full_integration",
		"testdata/resource_infinity_conference_alias_min",
		"internal/provider",
	} {
		require.NoError(t, os.MkdirAll(filepath.Join(repo, d), 0o750))
	}
	require.NoError(t, os.WriteFile(filepath.Join(repo, "internal/provider/resource_infinity_conference.go"), nil, 0o600))

	files := map[string]bool{}
	for _, f := range resourceFiles(repo, "conference") {
		files[f.Path] = f.Exists
	}
	assert.True(t, files["internal/provider/resource_infinity_conference.go"])
	assert.False(t, files["internal/provider/resource_infinity_conference_test.go"])
	assert.True(t, files["testdata/resource_infinity_conference_min"])
	assert.True(t, files["testdata/resource_infinity_conference_full_integration"])
	assert.NotContains(t, files, "testdata/resource_infinity_conference_alias_min")
}

func TestSDKModelGaps(t *testing.T) {
	sdk := t.TempDir()
	require.NoError(t, os.MkdirAll(filepath.Join(sdk, "config"), 0o750))
	model := "type Widget struct {\n\tName string `json:\"name\"`\n\tMode string `json:\"mode,omitempty\"`\n}\n"
	require.NoError(t, os.WriteFile(filepath.Join(sdk, "config", "widget_model.go"), []byte(model), 0o600))

	r := &UpgradeResource{TypeName: "infinity_widget", Schema: "widget", Changes: []UpgradeChange{
		{Field: "new_field", Issue: IssueMissingAttr},
		{Field: "mode", Issue: IssueTypeMismatch},
		{Field: "other", Issue: IssueDefault},
	}}
	path, missing := sdkModelGaps(sdk, r)
	assert.Equal(t, filepath.Join(sdk, "config", "widget_model.go"), path)
	assert.Equal(t, []string{"new_field"}, missing)
}

func TestWriteUpgradeOutputs(t *testing.T) {
	rep, resources, latest, baseline := upgradeFixture(t)
	plan := buildUpgradePlan(rep, resources, latest, baseline, scopeNew, PlanPaths{RepoDir: t.TempDir(), SDKDir: t.TempDir()})

	var buf bytes.Buffer
	writeUpgradeMarkdown(&buf, plan)
	assert.Contains(t, buf.String(), "# Provider upgrade plan: config-v42 → config-v43")
	assert.Contains(t, buf.String(), "## infinity_widget")
	assert.Contains(t, buf.String(), "### infinity_gadget")

	dir := t.TempDir()
	require.NoError(t, writeUpgradeDir(dir, plan))
	assert.FileExists(t, filepath.Join(dir, "index.md"))
	assert.FileExists(t, filepath.Join(dir, "infinity_widget.md"))
	index, err := os.ReadFile(filepath.Join(dir, "index.md"))
	require.NoError(t, err)
	assert.Contains(t, string(index), "[infinity_widget](infinity_widget.md)")
}
