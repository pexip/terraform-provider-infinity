/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

resource "tls_private_key" "tf-test-oauth-private-key" {
  algorithm = "RSA"
  rsa_bits  = 2048
}

resource "tls_self_signed_cert" "tf-test-oauth-certificate" {
  private_key_pem = tls_private_key.tf-test-oauth-private-key.private_key_pem

  subject {
    common_name  = "tf-test"
    organization = "pexip"
  }

  validity_period_hours = 12

  allowed_uses = [
    "key_encipherment",
    "digital_signature",
    "client_auth",
  ]
}

resource "infinity_mjx_graph_deployment" "test" {
  name                  = "tf-test mjx-graph-deployment full"
  description           = "Test MJX Graph deployment description"
  client_id             = "a1b2c3d4-e5f6-7890-abcd-ef1234567890"
  client_secret         = "test-client-secret"
  authentication_method = "CERTIFICATE"
  oauth_certificate     = tls_self_signed_cert.tf-test-oauth-certificate.cert_pem
  oauth_private_key     = tls_private_key.tf-test-oauth-private-key.private_key_pem
  oauth_token_url       = "https://login.microsoftonline.com/updated-tenant/oauth2/v2.0/token"
  graph_api_domain      = "graph.microsoft.com"
  request_quota         = 500000
  disable_proxy         = true
}
