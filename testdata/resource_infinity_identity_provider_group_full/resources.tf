/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

resource "infinity_identity_provider_group" "tf-test-identity-provider-group" {
  name        = "tf-test-identity-provider-group"
  description = "Test Identity Provider Group Description"
}
