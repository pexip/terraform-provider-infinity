//go:build unit

/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testResource() *TFResource {
	return newTFResource(context.Background(), "infinity_widget", schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true},
			"resource_id": schema.Int32Attribute{Computed: true},
			"name": schema.StringAttribute{
				Required:   true,
				Validators: []validator.String{stringvalidator.LengthAtMost(250)},
			},
			"mode": schema.StringAttribute{
				Optional:   true,
				Computed:   true,
				Default:    stringdefault.StaticString("a"),
				Validators: []validator.String{stringvalidator.OneOf("a", "b", "old")},
			},
			"template": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("line1\r\nline2"),
			},
			"enabled": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"count":     schema.StringAttribute{Optional: true},
			"hash":      schema.StringAttribute{Computed: true},
			"password":  schema.StringAttribute{Optional: true},
			"leftover":  schema.StringAttribute{Optional: true},
			"timestamp": schema.StringAttribute{Optional: true, Computed: true, Default: stringdefault.StaticString("")},
		},
	})
}

func apiField(typ string, mods ...func(*APIField)) *APIField {
	f := &APIField{Type: typ, Blank: true}
	for _, m := range mods {
		m(f)
	}
	return f
}

func withDefault(v any) func(*APIField) {
	return func(f *APIField) { f.Default, f.HasDefault = v, true }
}

func requiredField(f *APIField) { f.Blank = false }

func testAPISchema() *APISchema {
	return &APISchema{Name: "widget", Fields: map[string]*APIField{
		"id":           apiField("integer"),
		"resource_uri": apiField("string", func(f *APIField) { f.ReadOnly = true }),
		"name": apiField("string", requiredField, func(f *APIField) {
			f.HelpText = "The name. Maximum length: 100 characters."
		}),
		"mode": apiField("string", withDefault("a"), func(f *APIField) {
			f.ValidChoices = []any{"a", "b", "c"}
		}),
		"template":  apiField("string", withDefault("line1\nline2")),
		"enabled":   apiField("boolean", withDefault(true)),
		"count":     apiField("integer"),
		"hash":      apiField("string", requiredField),
		"password":  apiField("string"),
		"timestamp": apiField("datetime", withDefault("2026-11-14T16:49:52.050565")),
		"new_field": apiField("boolean", withDefault(false)),
		"server_ro": apiField("string", func(f *APIField) { f.ReadOnly = true }),
		"created":   apiField("datetime", withDefault("2026-11-14T16:49:52.050565")),
	}}
}

func findingsByField(findings []Finding) map[string][]string {
	out := map[string][]string{}
	for _, f := range findings {
		out[f.Field] = append(out[f.Field], f.Issue+": "+f.Detail)
	}
	return out
}

func TestCompareResource(t *testing.T) {
	got := findingsByField(compareResource(testResource(), testAPISchema(), Options{}))

	assert.Equal(t, []string{"max length mismatch: api max 100, provider max 250"}, got["name"])
	assert.Equal(t, []string{"choices mismatch: add c; remove old"}, got["mode"])
	assert.Equal(t, []string{"default mismatch: api true, provider false"}, got["enabled"])
	assert.Equal(t, []string{"type mismatch: api integer, provider string (expected int32|int64)"}, got["count"])
	assert.Equal(t, []string{"not sensitive: field name suggests a secret; consider Sensitive: true"}, got["password"])
	assert.Equal(t, []string{"extra attribute: string attribute not in API schema (removed or renamed?)"}, got["leftover"])
	assert.Equal(t, []string{"missing attribute: boolean, optional, default false"}, got["new_field"])

	// CRLF-only default differences, generated defaults, computed-only
	// attributes, ignored fields, and read-only API fields are not reported.
	for _, field := range []string{"template", "timestamp", "hash", "id", "resource_id", "resource_uri", "server_ro", "created"} {
		assert.Empty(t, got[field], field)
	}
}

func TestCompareResourceIncludeReadOnly(t *testing.T) {
	got := findingsByField(compareResource(testResource(), testAPISchema(), Options{IncludeReadOnly: true}))
	assert.Equal(t, []string{"missing attribute: string, read-only"}, got["server_ro"])
	assert.Empty(t, got["resource_uri"])
}

func TestCompareResourceRequiredMismatch(t *testing.T) {
	api := testAPISchema()
	api.Fields["password"] = apiField("string", requiredField)
	api.Fields["name"].Blank = true
	got := findingsByField(compareResource(testResource(), api, Options{}))
	assert.Contains(t, got["password"], "required mismatch: api required, provider optional")
	assert.Contains(t, got["name"], "required mismatch: api optional, provider required")
}

func TestTagFindings(t *testing.T) {
	latest := []Finding{
		{Resource: "r", Field: "a", Issue: IssueMissingAttr, Detail: "boolean"},
		{Resource: "r", Field: "b", Issue: IssueDefault, Detail: "api 2, provider 1"},
		{Resource: "r", Field: "c", Issue: IssueChoices, Detail: "add x"},
	}
	baseline := []Finding{
		{Resource: "r", Field: "b", Issue: IssueDefault, Detail: "api 1, provider 0"},
		{Resource: "r", Field: "c", Issue: IssueChoices, Detail: "add x"},
	}
	tagFindings(latest, baseline, 43)

	assert.Equal(t, "new in v43", latest[0].Tag)
	assert.Empty(t, latest[0].Baseline)
	assert.Equal(t, "new in v43", latest[1].Tag)
	assert.Equal(t, "api 1, provider 0", latest[1].Baseline)
	assert.Equal(t, tagPreExisting, latest[2].Tag)
}

func TestBuildReport(t *testing.T) {
	base := testAPISchema()
	delete(base.Fields, "new_field")
	latest := &SchemaSet{Dir: "config-v43", Version: 43, Schemas: map[string]*APISchema{
		"widget": testAPISchema(),
		"gadget": {Name: "gadget", Fields: map[string]*APIField{"name": apiField("string")}},
	}}
	baseline := &SchemaSet{Dir: "config-v42", Version: 42, Schemas: map[string]*APISchema{"widget": base}}

	rep := buildReport([]*TFResource{testResource()}, latest, baseline, Options{}, "")

	require.Len(t, rep.Resources, 1)
	tags := map[string]string{}
	for _, f := range rep.Resources[0].Findings {
		tags[f.Field] = f.Tag
	}
	assert.Equal(t, "new in v43", tags["new_field"])
	assert.Equal(t, tagPreExisting, tags["enabled"])

	require.Len(t, rep.Unmapped, 1)
	assert.Equal(t, "gadget", rep.Unmapped[0].Resource)
	assert.Equal(t, "new in v43", rep.Unmapped[0].Tag)
}

func TestSchemaChanges(t *testing.T) {
	base := &APISchema{Name: "widget", Fields: map[string]*APIField{
		"a": {Raw: map[string]any{"type": "string", "help_text": "old"}},
		"b": {Raw: map[string]any{"type": "string", "default": "2026-09-17T15:08:57.306828"}},
		"c": {Raw: map[string]any{"type": "integer", "default": float64(1)}},
	}}
	latest := &APISchema{Name: "widget", Fields: map[string]*APIField{
		"a": {Raw: map[string]any{"type": "string", "help_text": "new"}},
		"b": {Raw: map[string]any{"type": "string", "default": "2026-11-14T16:49:52.050565"}},
		"c": {Raw: map[string]any{"type": "integer", "default": float64(2)}},
	}}
	changes := schemaChanges(base, latest)

	require.Len(t, changes, 2)
	assert.Equal(t, Change{Schema: "widget", Field: "a", Keys: []string{"help_text"}}, changes[0])
	assert.Equal(t, Change{Schema: "widget", Field: "c", Keys: []string{"default"}, Detail: "default: 1 → 2"}, changes[1])
}

func TestProviderResourcesLoad(t *testing.T) {
	resources := loadProviderResources(context.Background())
	require.NotEmpty(t, resources)
	for _, r := range resources {
		assert.NotEmpty(t, r.Suffix, r.TypeName)
		assert.NotEmpty(t, r.Attrs, r.TypeName)
	}
}
