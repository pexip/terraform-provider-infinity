/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

resource "infinity_event_sink" "tf-test-event-sink" {
  name                   = "tf-test-event-sink"
  description            = "tf-test Event Sink Description"
  url                    = "https://tf-test-webhook.example.com/events"
  username               = "tf-test-user"
  password               = "tf-test-password"
  bulk_support           = true
  verify_tls_certificate = true
  version                = 2
}
