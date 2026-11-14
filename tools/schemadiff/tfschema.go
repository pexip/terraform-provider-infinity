/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"context"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/defaults"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	pexipProvider "github.com/pexip/terraform-provider-infinity/internal/provider"
)

const providerTypeName = "infinity"

// TFAttr is a flattened view of a provider resource attribute.
type TFAttr struct {
	Name       string
	Kind       string
	Required   bool
	Optional   bool
	Computed   bool
	Sensitive  bool
	HasDefault bool
	Default    any
	Choices    []string
	MaxLength  *int64
	Range      *[2]int64
}

// TFResource is a provider resource and its attributes.
type TFResource struct {
	TypeName string // e.g. infinity_conference
	Suffix   string // e.g. conference
	Attrs    map[string]*TFAttr
}

// loadProviderResources instantiates every provider resource and reads its schema.
func loadProviderResources(ctx context.Context) []*TFResource {
	p := pexipProvider.New()
	factories := p.Resources(ctx)
	out := make([]*TFResource, 0, len(factories))
	for _, factory := range factories {
		r := factory()
		var meta resource.MetadataResponse
		r.Metadata(ctx, resource.MetadataRequest{ProviderTypeName: providerTypeName}, &meta)
		var sch resource.SchemaResponse
		r.Schema(ctx, resource.SchemaRequest{}, &sch)
		out = append(out, newTFResource(ctx, meta.TypeName, sch.Schema))
	}
	sort.Slice(out, func(i, j int) bool { return out[i].TypeName < out[j].TypeName })
	return out
}

func newTFResource(ctx context.Context, typeName string, s schema.Schema) *TFResource {
	res := &TFResource{
		TypeName: typeName,
		Suffix:   strings.TrimPrefix(typeName, providerTypeName+"_"),
		Attrs:    make(map[string]*TFAttr, len(s.Attributes)),
	}
	for name, attr := range s.Attributes {
		res.Attrs[name] = newTFAttr(ctx, name, attr)
	}
	return res
}

func newTFAttr(ctx context.Context, name string, attr schema.Attribute) *TFAttr {
	a := &TFAttr{
		Name:      name,
		Required:  attr.IsRequired(),
		Optional:  attr.IsOptional(),
		Computed:  attr.IsComputed(),
		Sensitive: attr.IsSensitive(),
	}
	var descs []string
	switch t := attr.(type) {
	case schema.StringAttribute:
		a.Kind = "string"
		descs = describeValidators(ctx, t.Validators)
		if t.Default != nil {
			var resp defaults.StringResponse
			t.Default.DefaultString(ctx, defaults.StringRequest{}, &resp)
			a.setDefault(resp.PlanValue.IsNull(), resp.PlanValue.ValueString())
		}
	case schema.Int32Attribute:
		a.Kind = "int32"
		descs = describeValidators(ctx, t.Validators)
		if t.Default != nil {
			var resp defaults.Int32Response
			t.Default.DefaultInt32(ctx, defaults.Int32Request{}, &resp)
			a.setDefault(resp.PlanValue.IsNull(), float64(resp.PlanValue.ValueInt32()))
		}
	case schema.Int64Attribute:
		a.Kind = "int64"
		descs = describeValidators(ctx, t.Validators)
		if t.Default != nil {
			var resp defaults.Int64Response
			t.Default.DefaultInt64(ctx, defaults.Int64Request{}, &resp)
			a.setDefault(resp.PlanValue.IsNull(), float64(resp.PlanValue.ValueInt64()))
		}
	case schema.BoolAttribute:
		a.Kind = "bool"
		if t.Default != nil {
			var resp defaults.BoolResponse
			t.Default.DefaultBool(ctx, defaults.BoolRequest{}, &resp)
			a.setDefault(resp.PlanValue.IsNull(), resp.PlanValue.ValueBool())
		}
	case schema.Float64Attribute:
		a.Kind = "float64"
	case schema.NumberAttribute:
		a.Kind = "number"
	case schema.ListAttribute:
		a.Kind = "list"
	case schema.SetAttribute:
		a.Kind = "set"
	case schema.MapAttribute:
		a.Kind = "map"
	case schema.ObjectAttribute:
		a.Kind = "object"
	case schema.ListNestedAttribute:
		a.Kind = "list_nested"
	case schema.SetNestedAttribute:
		a.Kind = "set_nested"
	case schema.MapNestedAttribute:
		a.Kind = "map_nested"
	case schema.SingleNestedAttribute:
		a.Kind = "single_nested"
	default:
		a.Kind = "unknown"
	}
	a.applyValidatorDescriptions(descs)
	return a
}

func (a *TFAttr) setDefault(isNull bool, v any) {
	a.HasDefault = true
	if !isNull {
		a.Default = v
	}
}

func describeValidators[T validator.Describer](ctx context.Context, vs []T) []string {
	out := make([]string, 0, len(vs))
	for _, v := range vs {
		out = append(out, v.Description(ctx))
	}
	return out
}

var (
	oneOfPattern      = regexp.MustCompile(`^value must be one of: \[(.*)\]$`)
	quotedPattern     = regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)
	lenAtMostPattern  = regexp.MustCompile(`^string length must be at most (\d+)$`)
	lenBetweenPattern = regexp.MustCompile(`^string length must be between \d+ and (\d+)$`)
	betweenPattern    = regexp.MustCompile(`^value must be between (-?\d+) and (-?\d+)$`)
)

// applyValidatorDescriptions extracts OneOf choices, max length, and integer
// ranges from the descriptions of the standard framework validators.
func (a *TFAttr) applyValidatorDescriptions(descs []string) {
	for _, d := range descs {
		if m := oneOfPattern.FindStringSubmatch(d); m != nil && a.Kind == "string" {
			for _, q := range quotedPattern.FindAllString(m[1], -1) {
				if s, err := strconv.Unquote(q); err == nil {
					a.Choices = append(a.Choices, s)
				}
			}
			sort.Strings(a.Choices)
			continue
		}
		if m := lenAtMostPattern.FindStringSubmatch(d); m != nil {
			a.MaxLength = parseIntPtr(m[1])
			continue
		}
		if m := lenBetweenPattern.FindStringSubmatch(d); m != nil {
			a.MaxLength = parseIntPtr(m[1])
			continue
		}
		if m := betweenPattern.FindStringSubmatch(d); m != nil {
			lo, err1 := strconv.ParseInt(m[1], 10, 64)
			hi, err2 := strconv.ParseInt(m[2], 10, 64)
			if err1 == nil && err2 == nil {
				a.Range = &[2]int64{lo, hi}
			}
		}
	}
}

func parseIntPtr(s string) *int64 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return nil
	}
	return &n
}
