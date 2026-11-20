/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package provider

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/pexip/go-infinity-sdk/v42/config"
	"github.com/pexip/go-infinity-sdk/v42/types"
	"github.com/stretchr/testify/mock"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/pexip/go-infinity-sdk/v42"

	"github.com/pexip/terraform-provider-infinity/internal/test"
)

func TestInfinityIvrTheme(t *testing.T) {
	t.Parallel()
	_ = os.Setenv("TF_ACC", "1")

	client := infinity.NewClientMock()

	mockState := &config.IVRTheme{
		ID:          123,
		ResourceURI: "/api/admin/configuration/v1/ivr_theme/123/",
		Name:        "tf-test-ivr-theme",
		Version:     2,
	}

	createResponse := &types.PostResponse{
		Body:        []byte(""),
		ResourceURI: "/api/admin/configuration/v1/ivr_theme/123/",
	}

	// POST must never include the package file — the Pexip API ignores it on creation.
	client.On("PostMultipartFormWithFieldsAndResponse", mock.Anything, "configuration/v1/ivr_theme/",
		mock.Anything, "package", "", nil, mock.Anything).Return(createResponse, nil).Run(func(args mock.Arguments) {
		fields := args.Get(2).(map[string]string)
		if name, ok := fields["name"]; ok {
			mockState.Name = name
		}
		if version, ok := fields["version"]; ok {
			mockState.Version, _ = strconv.Atoi(version)
		}
	}).Maybe()

	client.On("GetJSON", mock.Anything, "configuration/v1/ivr_theme/123/", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*config.IVRTheme)
		*result = *mockState
	}).Maybe()

	// PATCH with a file — called when package is configured (steps 1 and 5).
	client.On("PatchMultipartFormWithFieldsAndResponse", mock.Anything, "configuration/v1/ivr_theme/123/",
		mock.Anything, "package", mock.MatchedBy(func(s string) bool { return s != "" }), mock.Anything, mock.Anything).Return(nil, nil).Run(func(args mock.Arguments) {
		fields := args.Get(2).(map[string]string)
		result := args.Get(6).(*config.IVRTheme)
		if name, ok := fields["name"]; ok && name != "" {
			mockState.Name = name
		}
		if version, ok := fields["version"]; ok {
			mockState.Version, _ = strconv.Atoi(version)
		}
		*result = *mockState
	}).Maybe()

	// PATCH without a file — called when package is not configured (step 2).
	client.On("PatchMultipartFormWithFieldsAndResponse", mock.Anything, "configuration/v1/ivr_theme/123/",
		mock.Anything, "package", "", mock.Anything, mock.Anything).Return(nil, nil).Run(func(args mock.Arguments) {
		fields := args.Get(2).(map[string]string)
		result := args.Get(6).(*config.IVRTheme)
		if name, ok := fields["name"]; ok && name != "" {
			mockState.Name = name
		}
		if version, ok := fields["version"]; ok {
			mockState.Version, _ = strconv.Atoi(version)
		}
		*result = *mockState
	}).Maybe()

	client.On("DeleteJSON", mock.Anything, "configuration/v1/ivr_theme/123/", mock.Anything).Return(nil).Maybe()

	testInfinityIvrTheme(t, client)
}

// TestInfinityIvrThemeVersion checks that an explicit version is sent in the
// multipart create and update requests when no package is uploaded. The mock
// starts without a version, so a version that isn't sent fails the checks.
func TestInfinityIvrThemeVersion(t *testing.T) {
	t.Parallel()
	_ = os.Setenv("TF_ACC", "1")

	client := infinity.NewClientMock()

	mockState := &config.IVRTheme{
		ID:          456,
		ResourceURI: "/api/admin/configuration/v1/ivr_theme/456/",
	}
	// Multipart fields from the most recent create (POST) and update (PATCH).
	var postFields, patchFields map[string]string

	client.On("PostMultipartFormWithFieldsAndResponse", mock.Anything, "configuration/v1/ivr_theme/",
		mock.Anything, "package", "", nil, mock.Anything).Return(&types.PostResponse{
		Body:        []byte(""),
		ResourceURI: "/api/admin/configuration/v1/ivr_theme/456/",
	}, nil).Run(func(args mock.Arguments) {
		postFields = args.Get(2).(map[string]string)
		mockState.Name = postFields["name"]
		mockState.Version, _ = strconv.Atoi(postFields["version"])
	}).Maybe()

	client.On("GetJSON", mock.Anything, "configuration/v1/ivr_theme/456/", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		*args.Get(3).(*config.IVRTheme) = *mockState
	}).Maybe()

	client.On("PatchMultipartFormWithFieldsAndResponse", mock.Anything, "configuration/v1/ivr_theme/456/",
		mock.Anything, "package", "", mock.Anything, mock.Anything).Return(nil, nil).Run(func(args mock.Arguments) {
		patchFields = args.Get(2).(map[string]string)
		if version, ok := patchFields["version"]; ok {
			mockState.Version, _ = strconv.Atoi(version)
		}
		*args.Get(6).(*config.IVRTheme) = *mockState
	}).Maybe()

	client.On("DeleteJSON", mock.Anything, "configuration/v1/ivr_theme/456/", mock.Anything).Return(nil).Maybe()

	configWithVersion := func(version int) string {
		return fmt.Sprintf(`
provider "infinity" {
  address  = "https://dev-manager.dev.pexip.network"
  username = "admin"
  password = "admin"
  insecure = true
}

resource "infinity_ivr_theme" "ivr_theme-test" {
  name    = "tf-test-ivr-theme-version"
  version = %d
}
`, version)
	}
	checkSentVersion := func(fields *map[string]string, want string) resource.TestCheckFunc {
		return func(*terraform.State) error {
			if got, ok := (*fields)["version"]; !ok || got != want {
				return fmt.Errorf("multipart version field = %q (present: %t), want %q", got, ok, want)
			}
			return nil
		}
	}

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: getTestProtoV6ProviderFactories(client),
		Steps: []resource.TestStep{
			// Create with an explicit non-default version
			{
				Config: configWithVersion(1),
				Check: resource.ComposeTestCheckFunc(
					checkSentVersion(&postFields, "1"),
					resource.TestCheckResourceAttr("infinity_ivr_theme.ivr_theme-test", "version", "1"),
				),
			},
			// Change the version on update
			{
				Config: configWithVersion(2),
				Check: resource.ComposeTestCheckFunc(
					checkSentVersion(&patchFields, "2"),
					resource.TestCheckResourceAttr("infinity_ivr_theme.ivr_theme-test", "version", "2"),
				),
			},
			// Change it back
			{
				Config: configWithVersion(1),
				Check: resource.ComposeTestCheckFunc(
					checkSentVersion(&patchFields, "1"),
					resource.TestCheckResourceAttr("infinity_ivr_theme.ivr_theme-test", "version", "1"),
				),
			},
		},
	})
}

func testInfinityIvrTheme(t *testing.T, client InfinityClient) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: getTestProtoV6ProviderFactories(client),
		Steps: []resource.TestStep{
			// Step 1: Create with full config
			{
				Config: test.LoadTestFolder(t, "resource_infinity_ivr_theme_full"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("infinity_ivr_theme.ivr_theme-test", "id"),
					resource.TestCheckResourceAttrSet("infinity_ivr_theme.ivr_theme-test", "resource_id"),
					resource.TestCheckResourceAttr("infinity_ivr_theme.ivr_theme-test", "name", "tf-test-ivr-theme-full"),
					resource.TestCheckResourceAttr("infinity_ivr_theme.ivr_theme-test", "version", "2"),
				),
			},
			// Step 2: Update to min config
			{
				Config: test.LoadTestFolder(t, "resource_infinity_ivr_theme_min"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("infinity_ivr_theme.ivr_theme-test", "id"),
					resource.TestCheckResourceAttrSet("infinity_ivr_theme.ivr_theme-test", "resource_id"),
					resource.TestCheckResourceAttr("infinity_ivr_theme.ivr_theme-test", "name", "tf-test-ivr-theme"),
					resource.TestCheckResourceAttr("infinity_ivr_theme.ivr_theme-test", "version", "2"),
				),
			},
			// Step 3: Delete
			{
				Config:       test.LoadTestFolder(t, "resource_infinity_ivr_theme_min"),
				ResourceName: "infinity_ivr_theme.ivr_theme-test",
				Destroy:      true,
			},
			// Step 4: Create with min config
			{
				Config: test.LoadTestFolder(t, "resource_infinity_ivr_theme_min"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("infinity_ivr_theme.ivr_theme-test", "id"),
					resource.TestCheckResourceAttrSet("infinity_ivr_theme.ivr_theme-test", "resource_id"),
					resource.TestCheckResourceAttr("infinity_ivr_theme.ivr_theme-test", "name", "tf-test-ivr-theme"),
					resource.TestCheckResourceAttr("infinity_ivr_theme.ivr_theme-test", "version", "2"),
				),
			},
			// Step 5: Update to full config
			{
				Config: test.LoadTestFolder(t, "resource_infinity_ivr_theme_full"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("infinity_ivr_theme.ivr_theme-test", "id"),
					resource.TestCheckResourceAttrSet("infinity_ivr_theme.ivr_theme-test", "resource_id"),
					resource.TestCheckResourceAttr("infinity_ivr_theme.ivr_theme-test", "name", "tf-test-ivr-theme-full"),
					resource.TestCheckResourceAttr("infinity_ivr_theme.ivr_theme-test", "version", "2"),
				),
			},
		},
	})
}
