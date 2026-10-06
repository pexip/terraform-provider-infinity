/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

resource "infinity_role" "test1" {
  name = "tf-test role 1 for role mapping"
}

resource "infinity_role" "test2" {
  name = "tf-test role 2 for role mapping"
}

resource "infinity_role_mapping" "test" {
  name   = "tf-test role mapping full"
  value  = "testfull"
  source = "OIDC"
  roles = [
    infinity_role.test1.id,
    infinity_role.test2.id
  ]
}