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

	"github.com/pexip/go-infinity-sdk/v40"
)

func TestInfinityGatewayRoutingRuleIntegration(t *testing.T) {
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

	testInfinityGatewayRoutingRuleFullIntegration(t, client)
}

func testInfinityGatewayRoutingRuleFullIntegration(t *testing.T, client InfinityClient) {
	fullConfig := test.LoadTestFolder(t, "resource_infinity_gateway_routing_rule_full_integration")
	minConfig := test.LoadTestFolder(t, "resource_infinity_gateway_routing_rule_min")

	fullChecks := resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttrSet("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "id"),
		resource.TestCheckResourceAttrSet("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "resource_id"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "name", "tf-test-gateway-routing-rule"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "description", "tf-test Gateway Routing Rule Description"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "match_string", ".*@example.com"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "priority", "66"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "enable", "false"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "match_string_full", "true"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "replace_string", "replaced@example.com"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "tag", "tf-test-tag"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "called_device_type", "registration"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "outgoing_protocol", "h323"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "call_type", "audio"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "crypto_mode", "on"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "denoise_audio", "false"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "max_pixels_per_second", "fullhd"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "max_callrate_in", "2048"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "max_callrate_out", "4096"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "match_incoming_calls", "false"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "match_outgoing_calls", "true"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "match_incoming_sip", "false"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "match_incoming_h323", "false"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "match_incoming_mssip", "false"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "match_incoming_webrtc", "false"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "match_incoming_teams", "true"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "match_incoming_only_if_registered", "false"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "enable_participant_avatar_lookup", "yes"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "live_captions_enabled", "yes"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "treat_as_trusted", "true"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "disabled_codecs.#", "2"),
		resource.TestCheckTypeSetElemAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "disabled_codecs.*", "H261"),
		resource.TestCheckTypeSetElemAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "disabled_codecs.*", "H263"),
		resource.TestCheckResourceAttrPair("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "ivr_theme", "pexip_infinity_ivr_theme.tf-test-theme", "id"),
	)

	minChecks := resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttrSet("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "id"),
		resource.TestCheckResourceAttrSet("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "resource_id"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "name", "tf-test-gateway-routing-rule"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "match_string", ".*@example.com"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "priority", "66"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "description", ""),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "enable", "true"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "tag", ""),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "replace_string", ""),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "called_device_type", "external"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "outgoing_protocol", "sip"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "call_type", "video"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "denoise_audio", "true"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "match_incoming_calls", "true"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "match_outgoing_calls", "false"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "match_incoming_sip", "true"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "match_incoming_h323", "true"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "match_incoming_mssip", "true"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "match_incoming_webrtc", "true"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "match_incoming_teams", "false"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "match_incoming_only_if_registered", "false"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "enable_participant_avatar_lookup", "default"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "live_captions_enabled", "default"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "treat_as_trusted", "false"),
		resource.TestCheckResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "disabled_codecs.#", "0"),
		resource.TestCheckNoResourceAttr("pexip_infinity_gateway_routing_rule.tf-test-gateway-routing-rule", "ivr_theme"),
	)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: getTestProtoV6ProviderFactories(client),
		Steps: []resource.TestStep{
			// Step 1: Create with min config
			{
				Config: minConfig,
				Check:  minChecks,
			},
			// Step 2: Update to full config
			{
				Config: fullConfig,
				Check:  fullChecks,
			},
			// Step 3: Destroy
			{
				Config:  fullConfig,
				Destroy: true,
			},
			// Step 4: Create with full config
			{
				Config: fullConfig,
				Check:  fullChecks,
			},
			// Step 5: Update to min config
			{
				Config: minConfig,
				Check:  minChecks,
			},
		},
	})
}
