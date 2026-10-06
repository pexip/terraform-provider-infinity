//go:build unit

/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/require"
)

func TestLegacyTypeNameMoveState(t *testing.T) {
	t.Parallel()

	sourceState := []byte(`{"id":"/api/admin/configuration/v1/dns_server/123/","resource_id":123,"address":"192.168.1.1","description":"tf-test"}`)

	tests := []struct {
		name            string
		providerAddress string
		sourceTypeName  string
		wantMoved       bool
	}{
		{"legacy type name", "registry.terraform.io/pexip/infinity", "pexip_infinity_dns_server", true},
		{"other resource type", "registry.terraform.io/pexip/infinity", "pexip_infinity_ntp_server", false},
		{"other provider", "registry.terraform.io/hashicorp/other", "pexip_infinity_dns_server", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			server, err := providerserver.NewProtocol6WithError(newTestProvider(nil))()
			require.NoError(t, err)

			resp, err := server.MoveResourceState(context.Background(), &tfprotov6.MoveResourceStateRequest{
				SourceProviderAddress: tt.providerAddress,
				SourceTypeName:        tt.sourceTypeName,
				SourceState:           &tfprotov6.RawState{JSON: sourceState},
				TargetTypeName:        "infinity_dns_server",
			})
			require.NoError(t, err)

			if !tt.wantMoved {
				require.NotEmpty(t, resp.Diagnostics)
				require.Nil(t, resp.TargetState)
				return
			}

			require.Empty(t, resp.Diagnostics)
			require.NotNil(t, resp.TargetState)

			objType := tftypes.Object{AttributeTypes: map[string]tftypes.Type{
				"id":          tftypes.String,
				"resource_id": tftypes.Number,
				"address":     tftypes.String,
				"description": tftypes.String,
			}}
			value, err := resp.TargetState.Unmarshal(objType)
			require.NoError(t, err)

			var attrs map[string]tftypes.Value
			require.NoError(t, value.As(&attrs))
			var address string
			require.NoError(t, attrs["address"].As(&address))
			require.Equal(t, "192.168.1.1", address)
		})
	}
}
