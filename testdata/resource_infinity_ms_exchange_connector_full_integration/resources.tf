/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

resource "pexip_infinity_ivr_theme" "tf-test-theme" {
  name = "tf-test-ms-exchange-connector-theme"
}

resource "pexip_infinity_ms_exchange_connector" "tf-test-ms-exchange-connector" {
  name                                             = "tf-test-ms-exchange-connector"
  description                                      = "tf-test MS Exchange Connector"
  meeting_buffer_before                            = 60
  meeting_buffer_after                             = 90
  scheduled_alias_suffix_length                    = 8
  room_mailbox_email_address                       = "tf-test@example.com"
  room_mailbox_name                                = "tf-test-ms-exchange-connector"
  url                                              = "https://tf-test.example.com/ews/exchange.asmx"
  username                                         = "tf-test-user"
  password                                         = "tf-test-password"
  authentication_method                            = "OAUTH"
  auth_provider                                    = "AZURE"
  uuid                                             = "12345678-1234-1234-1234-123456789012"
  scheduled_alias_prefix                           = "tf-test"
  scheduled_alias_domain                           = "tf-test.example.com"
  enable_dynamic_vmrs                              = true
  enable_personal_vmrs                             = true
  allow_new_users                                  = false
  disable_proxy                                    = true
  use_custom_add_in_sources                        = true
  enable_addin_debug_logs                          = true
  oauth_client_id                                  = "11111111-1111-1111-1111-111111111111"
  oauth_client_secret                              = "tf-test-oauth-client-secret"
  oauth_auth_endpoint                              = "https://tf-test.example.com/oauth/auth"
  oauth_token_endpoint                             = "https://tf-test.example.com/oauth/token"
  oauth_redirect_uri                               = "https://tf-test.example.com/admin/platform/msexchangeconnector/oauth_redirect/"
  kerberos_realm                                   = "TF-TEST.EXAMPLE.COM"
  kerberos_kdc                                     = "tf-test-kdc.example.com"
  kerberos_kdc_https_proxy                         = "https://tf-test-kdc-proxy.example.com"
  kerberos_exchange_spn                            = "exchangeMDB/tf-test.example.com"
  kerberos_enable_tls                              = true
  kerberos_auth_every_request                      = true
  kerberos_verify_tls_using_custom_ca              = true
  addin_server_domain                              = "tf-test.example.com"
  addin_display_name                               = "tf-test-ms-exchange-connector"
  addin_description                                = "tf-test MS Exchange Connector"
  addin_provider_name                              = "tf-test-ms-exchange-connector"
  addin_button_label                               = "tf-test Create a Pexip meeting"
  addin_group_label                                = "tf-test Pexip meeting"
  addin_supertip_title                             = "tf-test Makes this a Pexip meeting"
  addin_supertip_description                       = "tf-test MS Exchange Connector add-in"
  addin_application_id                             = "22222222-2222-2222-2222-222222222222"
  addin_authority_url                              = "https://tf-test.example.com/authority"
  addin_oidc_metadata_url                          = "https://tf-test.example.com/.well-known/openid-configuration"
  addin_authentication_method                      = "EXCHANGE_USER_ID_TOKEN"
  addin_naa_web_api_application_id                 = "33333333-3333-3333-3333-333333333333"
  personal_vmr_oauth_client_id                     = "44444444-4444-4444-4444-444444444444"
  personal_vmr_oauth_client_secret                 = "tf-test-personal-vmr-oauth-client-secret"
  personal_vmr_oauth_auth_endpoint                 = "https://tf-test.example.com/personal/oauth/auth"
  personal_vmr_oauth_token_endpoint                = "https://tf-test.example.com/personal/oauth/token"
  personal_vmr_adfs_relying_party_trust_identifier = "https://tf-test.example.com/adfs"
  office_js_url                                    = "https://appsforoffice.microsoft.com/lib/1/hosted/office.js"
  microsoft_fabric_url                             = "https://appsforoffice.microsoft.com/fabric/1.0/fabric.min.css"
  microsoft_fabric_components_url                  = "https://appsforoffice.microsoft.com/fabric/1.0/fabric.components.min.css"
  additional_add_in_script_sources                 = "https://tf-test.example.com/custom.js"
  ivr_theme                                        = pexip_infinity_ivr_theme.tf-test-theme.id
  non_idp_participants                             = "disallow_all"
}
