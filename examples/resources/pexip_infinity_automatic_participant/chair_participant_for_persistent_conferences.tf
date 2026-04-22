resource "pexip_infinity_automatic_participant" "chair_bot" {
  alias               = "meeting-chair"
  description         = "Automatic chair to keep conference active"
  conference          = data.pexip_infinity_conference.persistent_meeting.id
  protocol            = "webrtc"
  call_type           = "audio"
  role                = "chair"
  keep_conference_alive = "keep_conference_alive"
  routing             = "auto"
  streaming           = false
  remote_display_name = "Meeting Host"
}
