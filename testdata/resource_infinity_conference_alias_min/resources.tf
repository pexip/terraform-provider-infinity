/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

resource "infinity_conference" "tf-test-conference" {
  name         = "tf-test-conference"
  description  = "Test Conference"
  service_type = "conference"
}

resource "infinity_conference_alias" "tf-test-conference-alias" {
  alias      = "tf-test-alias"
  conference = infinity_conference.tf-test-conference.id
}
