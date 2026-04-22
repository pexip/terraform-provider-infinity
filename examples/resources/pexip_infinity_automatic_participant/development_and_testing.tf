resource "pexip_infinity_automatic_participant" "test_participant" {
  alias               = "test-${random_id.test.hex}"
  description         = "Test automatic participant"
  conference          = data.pexip_infinity_conference.test_conference.id
  protocol            = "webrtc"
  call_type           = "video"
  role                = "guest"
  keep_conference_alive = "end_conference_when_alone"
  routing             = "auto"
  streaming           = false
  remote_display_name = "Test Participant"
}

resource "random_id" "test" {
  byte_length = 4
}
