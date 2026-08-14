/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/require"
)

func TestInfinityManagerConfigModel(t *testing.T) {
	t.Parallel()

	base := InfinityManagerConfigModel{
		ID:              types.StringValue("test-id"),
		Hostname:        types.StringValue("manager-1"),
		Domain:          types.StringValue("example.com"),
		IP:              types.StringValue("10.5.6.7"),
		Mask:            types.StringValue("255.255.255.0"),
		GW:              types.StringValue("10.5.6.1"),
		DNS:             types.StringValue("1.1.1.1"),
		NTP:             types.StringValue("time.example.com"),
		User:            types.StringValue("admin"),
		Pass:            types.StringValue("password123"),
		AdminPassword:   types.StringValue("adminpass123"),
		EnableAnalytics: types.BoolValue(true),
	}

	tests := []struct {
		name                string
		errorReports        types.Bool
		contactEmailAddress types.String
		expectError         bool
	}{
		{
			name:                "error_reports true with valid email",
			errorReports:        types.BoolValue(true),
			contactEmailAddress: types.StringValue("test@example.com"),
			expectError:         false,
		},
		{
			name:                "error_reports true with no email",
			errorReports:        types.BoolValue(true),
			contactEmailAddress: types.StringNull(),
			expectError:         true,
		},
		{
			name:                "error_reports true with empty email",
			errorReports:        types.BoolValue(true),
			contactEmailAddress: types.StringValue(""),
			expectError:         true,
		},
		{
			name:                "error_reports true with invalid email",
			errorReports:        types.BoolValue(true),
			contactEmailAddress: types.StringValue("not-an-email"),
			expectError:         true,
		},
		{
			name:                "error_reports false with no email",
			errorReports:        types.BoolValue(false),
			contactEmailAddress: types.StringNull(),
			expectError:         false,
		},
		{
			name:                "error_reports false with no email",
			errorReports:        types.BoolValue(false),
			contactEmailAddress: types.StringNull(),
			expectError:         false,
		},
		{
			name:                "error_reports false with valid email",
			errorReports:        types.BoolValue(false),
			contactEmailAddress: types.StringValue("test@example.com"),
			expectError:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := base
			config.ErrorReports = tt.errorReports
			config.ContactEmailAddress = tt.contactEmailAddress

			diags := config.validate()
			if tt.expectError {
				require.True(t, diags.HasError(), "Expected validation error but got none")
			} else {
				require.False(t, diags.HasError(), "Expected no validation errors, got: %s", diags.Errors())
			}
		})
	}
}
