/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// APIField is a single field definition from an Infinity API schema file.
type APIField struct {
	Type                  string         `json:"type"`
	RelatedType           string         `json:"related_type,omitempty"`
	Nullable              bool           `json:"nullable"`
	Blank                 bool           `json:"blank"`
	ReadOnly              bool           `json:"readonly"`
	Unique                bool           `json:"unique"`
	Default               any            `json:"default,omitempty"`
	HasDefault            bool           `json:"-"`
	ValidChoices          []any          `json:"valid_choices,omitempty"`
	HelpText              string         `json:"help_text"`
	Deprecated            bool           `json:"deprecated,omitempty"`
	DeprecatedUseField    string         `json:"deprecated_use_field,omitempty"`
	DeprecatedUseResource string         `json:"deprecated_use_resource,omitempty"`
	Raw                   map[string]any `json:"-"`
}

// APISchema is the parsed schema for one API resource.
type APISchema struct {
	Name   string
	Fields map[string]*APIField
}

// SchemaSet is all API schemas for a single version, keyed by resource name.
type SchemaSet struct {
	Dir     string
	Version int
	Schemas map[string]*APISchema
}

var schemaDirPattern = regexp.MustCompile(`^(config|configuration)-v(\d+)$`)

// required reports whether the API rejects a create request without this field.
func (f *APIField) required() bool {
	return !f.Blank && !f.Nullable && !f.ReadOnly && (!f.HasDefault || f.Default == nil || f.Default == "")
}

var dynamicDefaultPattern = regexp.MustCompile(`^(\d{4}-\d{2}-\d{2}T[\d:.]+|[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})$`)

func (f *APIField) hasGeneratedDefault() bool {
	_, static := f.staticDefault()
	return f.HasDefault && f.Default != nil && !static
}

// staticDefault returns the field default and whether it is a fixed value.
// Timestamps and UUIDs are generated when the schema is downloaded, so they
// are treated as having no fixed default.
func (f *APIField) staticDefault() (any, bool) {
	if !f.HasDefault || f.Default == nil {
		return nil, false
	}
	if s, ok := f.Default.(string); ok {
		if dynamicDefaultPattern.MatchString(s) {
			return nil, false
		}
		return strings.ReplaceAll(s, "\r\n", "\n"), true
	}
	return f.Default, true
}

// stringChoices returns valid_choices for string fields, or nil for other types.
func (f *APIField) stringChoices() []string {
	if f.Type != "string" || len(f.ValidChoices) == 0 {
		return nil
	}
	out := make([]string, 0, len(f.ValidChoices))
	for _, c := range f.ValidChoices {
		out = append(out, fmt.Sprint(c))
	}
	sort.Strings(out)
	return out
}

var (
	maxLengthPattern = regexp.MustCompile(`Maximum length: (\d+) characters`)
	rangePattern     = regexp.MustCompile(`Range: (-?\d+) to (-?\d+)`)
)

// maxLength returns the maximum length documented in the help text, if any.
func (f *APIField) maxLength() (int64, bool) {
	m := maxLengthPattern.FindStringSubmatch(f.HelpText)
	if m == nil {
		return 0, false
	}
	n, err := strconv.ParseInt(m[1], 10, 64)
	return n, err == nil
}

// intRange returns the integer range documented in the help text, if any.
func (f *APIField) intRange() (lo, hi int64, ok bool) {
	m := rangePattern.FindStringSubmatch(f.HelpText)
	if m == nil {
		return 0, 0, false
	}
	lo, err1 := strconv.ParseInt(m[1], 10, 64)
	hi, err2 := strconv.ParseInt(m[2], 10, 64)
	return lo, hi, err1 == nil && err2 == nil
}

// describe returns a short human-readable summary of the field.
func (f *APIField) describe() string {
	parts := []string{f.Type}
	if f.RelatedType != "" {
		parts[0] = f.Type + "/" + f.RelatedType
	}
	switch {
	case f.ReadOnly:
		parts = append(parts, "read-only")
	case f.required():
		parts = append(parts, "required")
	default:
		parts = append(parts, "optional")
	}
	if d, ok := f.staticDefault(); ok {
		parts = append(parts, "default "+formatValue(d))
	} else if f.HasDefault && f.Default != nil {
		parts = append(parts, "generated default")
	}
	if c := f.stringChoices(); c != nil {
		parts = append(parts, "choices "+strings.Join(c, "|"))
	}
	if f.Deprecated {
		parts = append(parts, "deprecated")
	}
	return strings.Join(parts, ", ")
}

func parseSchemaFile(path string) (*APISchema, error) {
	data, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, err
	}
	var doc struct {
		Fields map[string]json.RawMessage `json:"fields"`
	}
	if err := json.Unmarshal(data, &doc); err != nil {
		return nil, fmt.Errorf("parsing %s: %w", path, err)
	}
	s := &APISchema{
		Name:   strings.TrimSuffix(filepath.Base(path), ".json"),
		Fields: make(map[string]*APIField, len(doc.Fields)),
	}
	for name, raw := range doc.Fields {
		f := &APIField{}
		if err := json.Unmarshal(raw, f); err != nil {
			return nil, fmt.Errorf("parsing %s field %s: %w", path, name, err)
		}
		if err := json.Unmarshal(raw, &f.Raw); err != nil {
			return nil, fmt.Errorf("parsing %s field %s: %w", path, name, err)
		}
		_, f.HasDefault = f.Raw["default"]
		s.Fields[name] = f
	}
	return s, nil
}

func loadSchemaSet(dir string) (*SchemaSet, error) {
	m := schemaDirPattern.FindStringSubmatch(filepath.Base(dir))
	if m == nil {
		return nil, fmt.Errorf("%s is not a config-vNN schema directory", dir)
	}
	version, _ := strconv.Atoi(m[2])
	files, err := filepath.Glob(filepath.Join(dir, "*.json"))
	if err != nil {
		return nil, err
	}
	if len(files) == 0 {
		return nil, fmt.Errorf("no schema files found in %s", dir)
	}
	set := &SchemaSet{Dir: dir, Version: version, Schemas: make(map[string]*APISchema, len(files))}
	for _, file := range files {
		s, err := parseSchemaFile(file)
		if err != nil {
			return nil, err
		}
		set.Schemas[s.Name] = s
	}
	return set, nil
}

// findSchemaDirs returns the schema directories under root keyed by version.
// config-vNN is preferred over configuration-vNN when both exist.
func findSchemaDirs(root string) (map[int]string, error) {
	entries, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}
	dirs := map[int]string{}
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		m := schemaDirPattern.FindStringSubmatch(e.Name())
		if m == nil {
			continue
		}
		v, _ := strconv.Atoi(m[2])
		if existing, ok := dirs[v]; ok && strings.HasPrefix(filepath.Base(existing), "config-") {
			continue
		}
		dirs[v] = filepath.Join(root, e.Name())
	}
	return dirs, nil
}

func formatValue(v any) string {
	switch t := v.(type) {
	case nil:
		return "null"
	case string:
		return strconv.Quote(t)
	case float64:
		return strconv.FormatFloat(t, 'f', -1, 64)
	default:
		b, err := json.Marshal(t)
		if err != nil {
			return fmt.Sprint(t)
		}
		return string(b)
	}
}
