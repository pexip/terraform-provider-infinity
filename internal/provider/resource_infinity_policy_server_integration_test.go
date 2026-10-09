//go:build integration

/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package provider

import (
	"crypto/tls"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stretchr/testify/require"

	"github.com/pexip/terraform-provider-infinity/internal/test"

	"github.com/pexip/go-infinity-sdk/v42"
)

func TestInfinityPolicyServerIntegration(t *testing.T) {
	_ = os.Setenv("TF_ACC", "1")

	client, err := infinity.New(
		infinity.WithBaseURL(test.INFINITY_BASE_URL),
		infinity.WithBasicAuth(test.INFINITY_USERNAME, test.INFINITY_PASSWORD),
		infinity.WithMaxRetries(2),
		infinity.WithTransport(&http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // We need this because default certificate is not trusted
				MinVersion:         tls.VersionTLS12,
			},
			MaxIdleConns:        30,
			MaxIdleConnsPerHost: 5,
			IdleConnTimeout:     60 * time.Second,
		}),
	)
	require.NoError(t, err)

	testInfinityPolicyServerIntegration(t, client)
}

func testInfinityPolicyServerIntegration(t *testing.T, client InfinityClient) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: getTestProtoV6ProviderFactories(client),
		Steps: []resource.TestStep{
			// Step 1: Create with full configuration
			{
				Config: test.LoadTestFolder(t, "resource_infinity_policy_server_full_integration"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("infinity_policy_server.tf-test-policy-server", "id"),
					resource.TestCheckResourceAttrSet("infinity_policy_server.tf-test-policy-server", "resource_id"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "name", "tf-test-policy-server"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "description", "tf-test Policy Server Description"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "url", "http://policy.example.com"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "allow_http", "true"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "username", "tf-test-user"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_service_lookup", "true"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_participant_lookup", "true"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_registration_lookup", "true"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_directory_lookup", "true"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_avatar_lookup", "true"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_media_location_lookup", "true"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_internal_service_policy", "true"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_internal_participant_policy", "true"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_internal_media_location_policy", "true"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "prefer_local_avatar_configuration", "true"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "internal_service_policy_template", "tf-test service template"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "internal_participant_policy_template", "tf-test participant template"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "internal_media_location_policy_template", "tf-test media location template"),
				),
			},
			// Step 2: Update to min configuration
			{
				Config: test.LoadTestFolder(t, "resource_infinity_policy_server_min"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("infinity_policy_server.tf-test-policy-server", "id"),
					resource.TestCheckResourceAttrSet("infinity_policy_server.tf-test-policy-server", "resource_id"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "name", "tf-test-policy-server"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "description", ""),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "url", ""),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "allow_http", "false"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "username", ""),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_service_lookup", "false"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_participant_lookup", "false"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_registration_lookup", "false"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_directory_lookup", "false"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_avatar_lookup", "false"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_media_location_lookup", "false"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_internal_service_policy", "false"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_internal_participant_policy", "false"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_internal_media_location_policy", "false"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "prefer_local_avatar_configuration", "false"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "internal_service_policy_template", ""),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "internal_participant_policy_template", ""),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "internal_media_location_policy_template", ""),
				),
			},
			// Step 3: Destroy resources before recreate-from-scratch test
			{
				Config:       test.LoadTestFolder(t, "resource_infinity_policy_server_min"),
				ResourceName: "infinity_policy_server.tf-test-policy-server",
				Destroy:      true,
			},
			// Step 4: Create with min configuration (after destroy)
			{
				Config: test.LoadTestFolder(t, "resource_infinity_policy_server_min"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("infinity_policy_server.tf-test-policy-server", "id"),
					resource.TestCheckResourceAttrSet("infinity_policy_server.tf-test-policy-server", "resource_id"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "name", "tf-test-policy-server"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "description", ""),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "url", ""),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "allow_http", "false"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "username", ""),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_service_lookup", "false"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_participant_lookup", "false"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_registration_lookup", "false"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_directory_lookup", "false"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_avatar_lookup", "false"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_media_location_lookup", "false"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_internal_service_policy", "false"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_internal_participant_policy", "false"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_internal_media_location_policy", "false"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "prefer_local_avatar_configuration", "false"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "internal_service_policy_template", ""),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "internal_participant_policy_template", ""),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "internal_media_location_policy_template", ""),
				),
			},
			// Step 5: Update to full configuration
			{
				Config: test.LoadTestFolder(t, "resource_infinity_policy_server_full_integration"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("infinity_policy_server.tf-test-policy-server", "id"),
					resource.TestCheckResourceAttrSet("infinity_policy_server.tf-test-policy-server", "resource_id"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "name", "tf-test-policy-server"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "description", "tf-test Policy Server Description"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "url", "http://policy.example.com"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "allow_http", "true"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "username", "tf-test-user"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_service_lookup", "true"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_participant_lookup", "true"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_registration_lookup", "true"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_directory_lookup", "true"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_avatar_lookup", "true"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_media_location_lookup", "true"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_internal_service_policy", "true"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_internal_participant_policy", "true"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "enable_internal_media_location_policy", "true"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "prefer_local_avatar_configuration", "true"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "internal_service_policy_template", "tf-test service template"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "internal_participant_policy_template", "tf-test participant template"),
					resource.TestCheckResourceAttr("infinity_policy_server.tf-test-policy-server", "internal_media_location_policy_template", "tf-test media location template"),
				),
			},
		},
	})
}
