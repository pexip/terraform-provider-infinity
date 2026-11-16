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

func TestInfinityMjxGraphDeploymentIntegration(t *testing.T) {
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

	testInfinityMjxGraphDeploymentIntegration(t, client)
}

func testInfinityMjxGraphDeploymentIntegration(t *testing.T, client InfinityClient) {
	fullConfig := test.LoadTestFolder(t, "resource_infinity_mjx_graph_deployment_full_integration")
	minConfig := test.LoadTestFolder(t, "resource_infinity_mjx_graph_deployment_min")

	fullChecks := resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttrSet("infinity_mjx_graph_deployment.test", "id"),
		resource.TestCheckResourceAttrSet("infinity_mjx_graph_deployment.test", "resource_id"),
		resource.TestCheckResourceAttr("infinity_mjx_graph_deployment.test", "name", "tf-test mjx-graph-deployment full"),
		resource.TestCheckResourceAttr("infinity_mjx_graph_deployment.test", "description", "Test MJX Graph deployment description"),
		resource.TestCheckResourceAttr("infinity_mjx_graph_deployment.test", "client_id", "a1b2c3d4-e5f6-7890-abcd-ef1234567890"),
		resource.TestCheckResourceAttr("infinity_mjx_graph_deployment.test", "client_secret", "test-client-secret"),
		resource.TestCheckResourceAttr("infinity_mjx_graph_deployment.test", "authentication_method", "CERTIFICATE"),
		resource.TestCheckResourceAttrPair("infinity_mjx_graph_deployment.test", "oauth_certificate", "tls_self_signed_cert.tf-test-oauth-certificate", "cert_pem"),
		resource.TestCheckResourceAttrPair("infinity_mjx_graph_deployment.test", "oauth_private_key", "tls_private_key.tf-test-oauth-private-key", "private_key_pem"),
		resource.TestCheckResourceAttr("infinity_mjx_graph_deployment.test", "oauth_token_url", "https://login.microsoftonline.com/updated-tenant/oauth2/v2.0/token"),
		resource.TestCheckResourceAttr("infinity_mjx_graph_deployment.test", "graph_api_domain", "graph.microsoft.com"),
		resource.TestCheckResourceAttr("infinity_mjx_graph_deployment.test", "request_quota", "500000"),
		resource.TestCheckResourceAttr("infinity_mjx_graph_deployment.test", "disable_proxy", "true"),
	)

	minChecks := resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttrSet("infinity_mjx_graph_deployment.test", "id"),
		resource.TestCheckResourceAttrSet("infinity_mjx_graph_deployment.test", "resource_id"),
		resource.TestCheckResourceAttr("infinity_mjx_graph_deployment.test", "name", "tf-test mjx-graph-deployment min"),
		resource.TestCheckResourceAttr("infinity_mjx_graph_deployment.test", "description", ""),
		resource.TestCheckResourceAttr("infinity_mjx_graph_deployment.test", "client_id", "12345678-1234-1234-1234-123456789012"),
		resource.TestCheckNoResourceAttr("infinity_mjx_graph_deployment.test", "client_secret"),
		resource.TestCheckResourceAttr("infinity_mjx_graph_deployment.test", "authentication_method", "CLIENT_SECRET"),
		resource.TestCheckResourceAttr("infinity_mjx_graph_deployment.test", "oauth_certificate", ""),
		resource.TestCheckNoResourceAttr("infinity_mjx_graph_deployment.test", "oauth_private_key"),
		resource.TestCheckResourceAttr("infinity_mjx_graph_deployment.test", "oauth_token_url", "https://login.microsoftonline.com/test-tenant/oauth2/v2.0/token"),
		resource.TestCheckResourceAttr("infinity_mjx_graph_deployment.test", "graph_api_domain", "graph.microsoft.com"),
		resource.TestCheckResourceAttr("infinity_mjx_graph_deployment.test", "request_quota", "1000000"),
		resource.TestCheckResourceAttr("infinity_mjx_graph_deployment.test", "disable_proxy", "false"),
	)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: getTestProtoV6ProviderFactories(client),
		ExternalProviders: map[string]resource.ExternalProvider{
			"tls": {
				Source: "hashicorp/tls",
			},
		},
		Steps: []resource.TestStep{
			// Step 1: Create with full config
			{
				Config: fullConfig,
				Check:  fullChecks,
			},
			// Step 2: Update to min config
			{
				Config: minConfig,
				Check:  minChecks,
			},
			// Step 3: Destroy
			{
				Config:  minConfig,
				Destroy: true,
			},
			// Step 4: Create with min config
			{
				Config: minConfig,
				Check:  minChecks,
			},
			// Step 5: Update to full config
			{
				Config: fullConfig,
				Check:  fullChecks,
			},
		},
	})
}
