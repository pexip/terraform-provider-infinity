resource "pexip_infinity_event_sink" "authenticated_sink" {
  name        = "Authenticated Event Sink"
  description = "Event sink with HTTP basic authentication"
  url         = "https://events.company.com/api/pexip-events"
  username    = "pexip-integration"
  password    = "secure-password-123"
  version     = 1
}
