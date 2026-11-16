/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"fmt"
	"path/filepath"
	"reflect"
	"regexp"
	"slices"
	"sort"
	"strings"
)

// Issue types reported by the comparison.
const (
	IssueMissingAttr   = "missing attribute"
	IssueExtraAttr     = "extra attribute"
	IssueTypeMismatch  = "type mismatch"
	IssueRequired      = "required mismatch"
	IssueReadOnly      = "read-only in API"
	IssueDefault       = "default mismatch"
	IssueChoices       = "choices mismatch"
	IssueMaxLength     = "max length mismatch"
	IssueRange         = "range mismatch"
	IssueDeprecated    = "deprecated in API"
	IssueNotSensitive  = "not sensitive"
	IssueNoProviderRes = "no provider resource"
	IssueNoAPISchema   = "no API schema"
	tagPreExisting     = "pre-existing"
	tagNewFormat       = "new in v%d"
	noValidatorMarker  = "provider has no"
)

// Finding is a single difference between the API schema and the provider.
type Finding struct {
	Resource string `json:"resource"`
	Schema   string `json:"schema,omitempty"`
	Field    string `json:"field,omitempty"`
	Issue    string `json:"issue"`
	Detail   string `json:"detail"`
	Tag      string `json:"tag"`
	Baseline string `json:"baseline,omitempty"`
}

func (f Finding) key() string { return f.Resource + "|" + f.Field + "|" + f.Issue }

// Change is an API field change between the baseline and latest schemas.
type Change struct {
	Resource string   `json:"resource,omitempty"`
	Schema   string   `json:"schema"`
	Field    string   `json:"field"`
	Keys     []string `json:"keys"`
	Detail   string   `json:"detail"`
}

// ResourceReport groups findings for one provider resource.
type ResourceReport struct {
	TypeName string    `json:"type_name"`
	Schema   string    `json:"schema"`
	Findings []Finding `json:"findings"`
}

// Report is the full comparison output.
type Report struct {
	Latest    string           `json:"latest"`
	Baseline  string           `json:"baseline"`
	Resources []ResourceReport `json:"resources"`
	Unmapped  []Finding        `json:"unmapped"`
	Changes   []Change         `json:"changes"`
}

// Options control which checks are reported.
type Options struct {
	IncludeReadOnly bool
}

// allowedKinds returns the provider attribute kinds compatible with an API field.
func allowedKinds(f *APIField) []string {
	switch f.Type {
	case "string", "datetime", "date", "time", "file":
		return []string{"string"}
	case "integer":
		return []string{"int32", "int64"}
	case "boolean":
		return []string{"bool"}
	case "related":
		if f.RelatedType == "to_many" {
			return []string{"set", "list", "set_nested", "list_nested"}
		}
		return []string{"string"}
	case "list":
		return []string{"list", "set"}
	case "dict":
		return []string{"map", "string", "object", "single_nested"}
	default:
		return nil
	}
}

var sensitivePattern = regexp.MustCompile(`(^|_)(password|secret|pin|private_key|credentials?)($|_)`)

// compareResource compares one provider resource against one API schema.
func compareResource(res *TFResource, api *APISchema, opts Options) []Finding {
	var out []Finding
	add := func(field, issue, detail string) {
		out = append(out, Finding{Resource: res.TypeName, Schema: api.Name, Field: field, Issue: issue, Detail: detail})
	}

	for _, name := range sortedKeys(api.Fields) {
		f := api.Fields[name]
		if ignoredAPIFields[name] || isIgnored(res.Suffix, name) {
			continue
		}
		attr, ok := res.Attrs[name]
		if !ok {
			// Fields with generated defaults (timestamps, UUIDs) are set by the server.
			if (f.ReadOnly || f.hasGeneratedDefault()) && !opts.IncludeReadOnly {
				continue
			}
			add(name, IssueMissingAttr, f.describe())
			continue
		}
		compareField(name, f, attr, add)
	}

	for _, name := range sortedKeys(res.Attrs) {
		if ignoredTFAttrs[name] || isIgnored(res.Suffix, name) {
			continue
		}
		if _, ok := api.Fields[name]; !ok {
			add(name, IssueExtraAttr, res.Attrs[name].Kind+" attribute not in API schema (removed or renamed?)")
		}
	}
	return out
}

func compareField(name string, f *APIField, attr *TFAttr, add func(field, issue, detail string)) {
	if kinds := allowedKinds(f); kinds != nil && !slices.Contains(kinds, attr.Kind) {
		add(name, IssueTypeMismatch, fmt.Sprintf("api %s, provider %s (expected %s)", f.describeType(), attr.Kind, strings.Join(kinds, "|")))
	}

	if f.ReadOnly {
		if attr.Required || attr.Optional {
			add(name, IssueReadOnly, "api field is read-only but provider attribute is configurable")
		}
	}

	// Computed-only attributes are never set by users, so input constraints
	// (required, choices, length, range) don't apply to them.
	computedOnly := attr.Computed && !attr.Required && !attr.Optional

	if !f.ReadOnly && !computedOnly && f.required() != attr.Required {
		add(name, IssueRequired, fmt.Sprintf("api %s, provider %s", requiredWord(f.required()), requiredWord(attr.Required)))
	}

	if d, ok := f.staticDefault(); ok && attr.HasDefault && !valuesEqual(d, attr.Default) {
		add(name, IssueDefault, fmt.Sprintf("api %s, provider %s", formatValue(d), formatValue(attr.Default)))
	}

	if apiChoices := f.stringChoices(); !computedOnly && (apiChoices != nil || attr.Choices != nil) {
		missing, extra := diffStrings(apiChoices, attr.Choices)
		switch {
		case attr.Choices == nil:
			add(name, IssueChoices, fmt.Sprintf(noValidatorMarker+" OneOf validator; api allows %s", summarize(apiChoices)))
		case len(missing) > 0 || len(extra) > 0:
			var parts []string
			if len(missing) > 0 {
				parts = append(parts, "add "+summarize(missing))
			}
			if len(extra) > 0 {
				parts = append(parts, "remove "+summarize(extra))
			}
			add(name, IssueChoices, strings.Join(parts, "; "))
		}
	}

	if n, ok := f.maxLength(); ok && !computedOnly && attr.Kind == "string" {
		switch {
		case attr.MaxLength == nil:
			add(name, IssueMaxLength, fmt.Sprintf("api max %d, "+noValidatorMarker+" LengthAtMost validator", n))
		case *attr.MaxLength != n:
			add(name, IssueMaxLength, fmt.Sprintf("api max %d, provider max %d", n, *attr.MaxLength))
		}
	}

	if lo, hi, ok := f.intRange(); ok && !computedOnly && (attr.Kind == "int32" || attr.Kind == "int64") {
		switch {
		case attr.Range == nil:
			add(name, IssueRange, fmt.Sprintf("api range %d-%d, "+noValidatorMarker+" Between validator", lo, hi))
		case attr.Range[0] != lo || attr.Range[1] != hi:
			add(name, IssueRange, fmt.Sprintf("api range %d-%d, provider range %d-%d", lo, hi, attr.Range[0], attr.Range[1]))
		}
	}

	if f.Deprecated {
		detail := "api field is deprecated"
		if f.DeprecatedUseField != "" {
			detail += ", use field " + f.DeprecatedUseField
		}
		if f.DeprecatedUseResource != "" {
			detail += ", use resource " + f.DeprecatedUseResource
		}
		add(name, IssueDeprecated, detail)
	}

	if sensitivePattern.MatchString(name) && !strings.HasSuffix(name, "_enabled") && !attr.Sensitive && attr.Kind == "string" {
		add(name, IssueNotSensitive, "field name suggests a secret; consider Sensitive: true")
	}
}

// summarize joins values, truncating long lists such as timezones.
func summarize(values []string) string {
	const maxShown = 10
	if len(values) <= maxShown {
		return strings.Join(values, ", ")
	}
	return fmt.Sprintf("%s, … (%d values)", strings.Join(values[:maxShown], ", "), len(values))
}

func (f *APIField) describeType() string {
	if f.RelatedType != "" {
		return f.Type + "/" + f.RelatedType
	}
	return f.Type
}

func requiredWord(b bool) string {
	if b {
		return "required"
	}
	return "optional"
}

func valuesEqual(api, tf any) bool {
	switch a := api.(type) {
	case float64:
		b, ok := tf.(float64)
		return ok && a == b
	case string:
		b, ok := tf.(string)
		return ok && a == strings.ReplaceAll(b, "\r\n", "\n")
	case bool:
		return api == tf
	default:
		return reflect.DeepEqual(api, tf)
	}
}

// diffStrings returns values in want but not have, and in have but not want.
func diffStrings(want, have []string) (missing, extra []string) {
	for _, w := range want {
		if !slices.Contains(have, w) {
			missing = append(missing, w)
		}
	}
	for _, h := range have {
		if !slices.Contains(want, h) {
			extra = append(extra, h)
		}
	}
	return missing, extra
}

// tagFindings marks each finding as new in the latest version or pre-existing,
// by checking whether the same finding is produced against the baseline schema.
func tagFindings(latest, baseline []Finding, version int) {
	byKey := make(map[string]Finding, len(baseline))
	for _, f := range baseline {
		byKey[f.key()] = f
	}
	for i := range latest {
		b, ok := byKey[latest[i].key()]
		switch {
		case !ok:
			latest[i].Tag = fmt.Sprintf(tagNewFormat, version)
		case b.Detail != latest[i].Detail:
			latest[i].Tag = fmt.Sprintf(tagNewFormat, version)
			latest[i].Baseline = b.Detail
		default:
			latest[i].Tag = tagPreExisting
		}
	}
}

// buildReport compares every provider resource against the latest schemas and
// tags findings using the baseline schemas.
func buildReport(resources []*TFResource, latest, baseline *SchemaSet, opts Options, filter string) *Report {
	rep := &Report{Latest: filepath.Base(latest.Dir), Baseline: filepath.Base(baseline.Dir)}
	mapped := map[string]bool{}

	for _, res := range resources {
		if filter != "" && res.Suffix != filter {
			continue
		}
		if _, skip := skippedResources[res.Suffix]; skip {
			continue
		}
		name := schemaNameFor(res.Suffix)
		api, ok := latest.Schemas[name]
		if !ok {
			rep.Unmapped = append(rep.Unmapped, Finding{
				Resource: res.TypeName, Issue: IssueNoAPISchema,
				Detail: fmt.Sprintf("no %s.json in %s", name, rep.Latest), Tag: tagPreExisting,
			})
			continue
		}
		mapped[name] = true

		findings := compareResource(res, api, opts)
		var baseFindings []Finding
		if base, ok := baseline.Schemas[name]; ok {
			baseFindings = compareResource(res, base, opts)
			for _, c := range schemaChanges(base, api) {
				c.Resource = res.TypeName
				rep.Changes = append(rep.Changes, c)
			}
		}
		tagFindings(findings, baseFindings, latest.Version)
		if len(findings) > 0 {
			rep.Resources = append(rep.Resources, ResourceReport{TypeName: res.TypeName, Schema: name, Findings: findings})
		}
	}

	if filter == "" {
		for _, name := range sortedKeys(latest.Schemas) {
			if mapped[name] {
				continue
			}
			tag := tagPreExisting
			if _, ok := baseline.Schemas[name]; !ok {
				tag = fmt.Sprintf(tagNewFormat, latest.Version)
			}
			rep.Unmapped = append(rep.Unmapped, Finding{
				Resource: name, Schema: name, Issue: IssueNoProviderRes,
				Detail: fmt.Sprintf("%d fields", len(latest.Schemas[name].Fields)), Tag: tag,
			})
		}
	}
	return rep
}

// schemaChanges lists fields present in both schemas whose definition changed.
func schemaChanges(base, latest *APISchema) []Change {
	var out []Change
	for _, name := range sortedKeys(latest.Fields) {
		b, ok := base.Fields[name]
		if !ok {
			continue
		}
		l := latest.Fields[name]
		keys := map[string]bool{}
		for k := range b.Raw {
			keys[k] = true
		}
		for k := range l.Raw {
			keys[k] = true
		}
		var changed, details []string
		for _, k := range sortedKeys(keys) {
			bv, bok := b.Raw[k]
			lv, lok := l.Raw[k]
			if bok == lok && (reflect.DeepEqual(bv, lv) || (k == "default" && isDynamic(bv) && isDynamic(lv))) {
				continue
			}
			changed = append(changed, k)
			if k != "help_text" {
				details = append(details, fmt.Sprintf("%s: %s → %s", k, rawValue(bv, bok), rawValue(lv, lok)))
			}
		}
		if len(changed) > 0 {
			out = append(out, Change{Schema: latest.Name, Field: name, Keys: changed, Detail: strings.Join(details, "; ")})
		}
	}
	return out
}

func isDynamic(v any) bool {
	s, ok := v.(string)
	return ok && dynamicDefaultPattern.MatchString(s)
}

func rawValue(v any, present bool) string {
	if !present {
		return "(absent)"
	}
	return formatValue(v)
}

func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
