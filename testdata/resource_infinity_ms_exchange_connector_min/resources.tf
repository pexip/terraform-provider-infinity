/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

resource "pexip_infinity_ms_exchange_connector" "tf-test-ms-exchange-connector" {
  name               = "tf-test-ms-exchange-connector"
  addin_server_domain = "test-server.example.com"
}
