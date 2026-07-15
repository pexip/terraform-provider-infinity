/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

resource "infinity_end_user" "tf-test-end-user" {
  primary_email_address = "tf-test-user@example.com"
}
