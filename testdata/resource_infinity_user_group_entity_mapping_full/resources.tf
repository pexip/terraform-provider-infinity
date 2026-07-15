/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

resource "infinity_user_group" "tf-test-group" {
  name = "tf-test-user-group-full"
}

resource "infinity_conference" "tf-test-conference" {
  name         = "tf-test-conference-full"
  service_type = "conference"
}

resource "infinity_user_group_entity_mapping" "tf-test-mapping" {
  description         = "tf-test user group entity mapping description"
  entity_resource_uri = infinity_conference.tf-test-conference.id
  user_group          = infinity_user_group.tf-test-group.id

  depends_on = [
    infinity_user_group.tf-test-group,
    infinity_conference.tf-test-conference
  ]
}
