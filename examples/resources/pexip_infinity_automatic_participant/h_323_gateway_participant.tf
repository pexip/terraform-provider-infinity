resource "pexip_infinity_automatic_participant" "h323_gateway" {
  alias                 = "legacy-system"
  description           = "H.323 gateway for legacy video systems"
  conference            = data.pexip_infinity_conference.hybrid_meeting.id
  protocol              = "h323"
  call_type             = "video"
  role                  = "guest"
  keep_conference_alive = "end_conference_when_alone"
  routing               = "manual"
  system_location       = data.pexip_infinity_system_location.gateway.id
  dtmf_sequence         = "12345#"
  remote_display_name   = "Legacy Video System"
}
