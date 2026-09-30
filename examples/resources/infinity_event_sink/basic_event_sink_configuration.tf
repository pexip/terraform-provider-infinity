resource "pexip_infinity_event_sink" "basic_sink" {
  name = "Basic Event Sink"
  url  = "https://events.example.com/webhook"
}
