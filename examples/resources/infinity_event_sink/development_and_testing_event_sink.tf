resource "pexip_infinity_event_sink" "development" {
  name                   = "Development Event Sink"
  description            = "Event sink for development and testing"
  url                    = "https://webhook.site/unique-id"
  verify_tls_certificate = false
  version                = 1
}
