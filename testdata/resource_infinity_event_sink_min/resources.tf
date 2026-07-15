/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

resource "infinity_event_sink" "tf-test-event-sink" {
  name = "tf-test-event-sink"
  url  = "https://tf-test-webhook.example.com/events"
}
