/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */
resource "infinity_turn_server" "tf-test-turn-server" {
  name    = "tf-test-turn-server"
  address = "turn.example.com"
}
