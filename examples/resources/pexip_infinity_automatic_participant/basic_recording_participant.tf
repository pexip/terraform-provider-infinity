resource "pexip_infinity_automatic_participant" "recorder" {
  alias               = "conference-recorder"
  description         = "Automatic recording participant"
  conference          = data.pexip_infinity_conference.meeting_room.id
  protocol            = "sip"
  call_type           = "video"
  role                = "guest"
  keep_conference_alive = "end_conference_when_alone"
  routing             = "auto"
  streaming           = true
  remote_display_name = "Conference Recorder"
}
