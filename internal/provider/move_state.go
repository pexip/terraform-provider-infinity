/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package provider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

const (
	providerTypeName = "infinity"

	// legacyTypeNamePrefix is the resource type name prefix used before v42,
	// e.g. pexip_infinity_conference is now infinity_conference.
	legacyTypeNamePrefix = "pexip_infinity_"

	// legacyProviderAddressSuffix matches the provider source address of the
	// legacy resources, regardless of registry host.
	legacyProviderAddressSuffix = "pexip/infinity"
)

// legacyTypeNameStateMovers returns a StateMover that allows a resource to be
// migrated from its legacy type name with a moved block, e.g.
//
//	moved {
//	  from = pexip_infinity_conference.example
//	  to   = infinity_conference.example
//	}
//
// The schema is unchanged between the legacy and current type names, so the
// source state is copied as-is.
//
// This helper can be removed in v45, along with the MoveState methods and
// ResourceWithMoveState assertions on each resource. Resources added after v42
// never had a legacy type name, so they don't need a MoveState method.
func legacyTypeNameStateMovers(ctx context.Context, r resource.Resource) []resource.StateMover {
	metadataResp := &resource.MetadataResponse{}
	r.Metadata(ctx, resource.MetadataRequest{ProviderTypeName: providerTypeName}, metadataResp)
	legacyTypeName := legacyTypeNamePrefix + strings.TrimPrefix(metadataResp.TypeName, providerTypeName+"_")

	schemaResp := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResp)
	sourceSchema := schemaResp.Schema

	return []resource.StateMover{
		{
			SourceSchema: &sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != legacyTypeName || !strings.HasSuffix(req.SourceProviderAddress, legacyProviderAddressSuffix) {
					return
				}
				if req.SourceState == nil {
					resp.Diagnostics.AddError(
						"Unable to Move Resource State",
						"The source state for "+legacyTypeName+" could not be read using the current schema. "+
							"Please upgrade to the latest provider version using the legacy resource type name before moving the resource.",
					)
					return
				}
				resp.TargetState.Raw = req.SourceState.Raw
			},
		},
	}
}
