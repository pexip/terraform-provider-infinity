resource "pexip_infinity_conference" "test_call" {
  name         = "Test Call Service"
  service_type = "test_call"
  description  = "Service for testing audio and video"
  aliases = [
    { alias = "calltest" },
    { alias = "calltest@example.com" },
  ]
}
