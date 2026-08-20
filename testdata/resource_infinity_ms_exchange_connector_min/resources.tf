/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

resource "pexip_infinity_ms_exchange_connector" "tf-test-ms-exchange-connector" {
  name                                             = "tf-test-ms-exchange-connector"
  addin_server_domain                              = "test-server.example.com"
  enable_dynamic_vmrs                              = false
  enable_personal_vmrs                             = true
  personal_vmr_oauth_auth_endpoint                 = "https://tf-test.example.com/personal/oauth/auth"
  personal_vmr_oauth_token_endpoint                = "https://tf-test.example.com/personal/oauth/token"
  personal_vmr_adfs_relying_party_trust_identifier = "https://tf-test.example.com/adfs"
}
