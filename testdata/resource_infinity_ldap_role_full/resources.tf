/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

resource "infinity_role" "test1" {
  name = "tf-test role 1 for ldap role"
}

resource "infinity_role" "test2" {
  name = "tf-test role 2 for ldap role"
}

resource "infinity_ldap_role" "ldap_role-test" {
  name          = "tf-test full"
  ldap_group_dn = "testfull"
  roles = [
    infinity_role.test1.id,
    infinity_role.test2.id,
  ]
}