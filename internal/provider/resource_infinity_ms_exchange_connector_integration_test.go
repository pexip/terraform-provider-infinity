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

	"github.com/pexip/go-infinity-sdk/v41"
)

func TestInfinityMsExchangeConnectorIntegration(t *testing.T) {
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

	testInfinityMsExchangeConnectorIntegration(t, client)
}

func testInfinityMsExchangeConnectorIntegration(t *testing.T, client InfinityClient) {
	fullConfig := test.LoadTestFolder(t, "resource_infinity_ms_exchange_connector_full_integration")
	minConfig := test.LoadTestFolder(t, "resource_infinity_ms_exchange_connector_min")

	fullChecks := resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttrSet("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "id"),
		resource.TestCheckResourceAttrSet("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "resource_id"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "name", "tf-test-ms-exchange-connector"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "description", "tf-test MS Exchange Connector"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "meeting_buffer_before", "60"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "meeting_buffer_after", "90"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "scheduled_alias_suffix_length", "8"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "room_mailbox_email_address", "tf-test@example.com"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "url", "https://tf-test.example.com/ews/exchange.asmx"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "authentication_method", "OAUTH"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "auth_provider", "AZURE"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "uuid", "12345678-1234-1234-1234-123456789012"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "scheduled_alias_domain", "tf-test.example.com"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "enable_dynamic_vmrs", "true"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "enable_personal_vmrs", "true"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "allow_new_users", "false"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "disable_proxy", "true"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "use_custom_add_in_sources", "true"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "enable_addin_debug_logs", "true"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "kerberos_enable_tls", "true"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "kerberos_auth_every_request", "true"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "kerberos_verify_tls_using_custom_ca", "true"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "oauth_client_id", "11111111-1111-1111-1111-111111111111"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "addin_server_domain", "tf-test.example.com"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "addin_application_id", "22222222-2222-2222-2222-222222222222"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "addin_naa_web_api_application_id", "33333333-3333-3333-3333-333333333333"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "addin_authentication_method", "EXCHANGE_USER_ID_TOKEN"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "personal_vmr_oauth_client_id", "44444444-4444-4444-4444-444444444444"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "non_idp_participants", "disallow_all"),
		resource.TestCheckResourceAttrSet("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "ivr_theme"),
	)

	minChecks := resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttrSet("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "id"),
		resource.TestCheckResourceAttrSet("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "resource_id"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "name", "tf-test-ms-exchange-connector"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "description", ""),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "meeting_buffer_before", "30"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "meeting_buffer_after", "60"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "scheduled_alias_suffix_length", "6"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "authentication_method", "BASIC"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "auth_provider", "ADFS"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "enable_dynamic_vmrs", "false"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "enable_personal_vmrs", "true"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "personal_vmr_oauth_auth_endpoint", "https://tf-test.example.com/personal/oauth/auth"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "personal_vmr_oauth_token_endpoint", "https://tf-test.example.com/personal/oauth/token"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "personal_vmr_adfs_relying_party_trust_identifier", "https://tf-test.example.com/adfs"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "allow_new_users", "true"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "disable_proxy", "false"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "use_custom_add_in_sources", "false"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "enable_addin_debug_logs", "false"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "kerberos_enable_tls", "true"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "kerberos_auth_every_request", "false"),
		resource.TestCheckResourceAttr("pexip_infinity_ms_exchange_connector.tf-test-ms-exchange-connector", "kerberos_verify_tls_using_custom_ca", "false"),
	)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: getTestProtoV6ProviderFactories(client),
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
