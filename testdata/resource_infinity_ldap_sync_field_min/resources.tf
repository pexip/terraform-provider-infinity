/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

resource "infinity_ldap_sync_field" "test" {
  name                   = "tf-test-min"
  template_variable_name = "testmin"
}