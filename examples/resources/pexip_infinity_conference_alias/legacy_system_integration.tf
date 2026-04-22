# Legacy phone system integration
resource "pexip_infinity_conference_alias" "legacy_bridge" {
  alias       = "9001"
  description = "Bridge number for legacy phone system"
  conference  = data.pexip_infinity_conference.bridge_conference.id
}

# H.323 system alias
resource "pexip_infinity_conference_alias" "h323_room" {
  alias       = "h323.room1"
  description = "H.323 compatible room alias"
  conference  = data.pexip_infinity_conference.video_room.id
}

# SIP endpoint alias
resource "pexip_infinity_conference_alias" "sip_endpoint" {
  alias       = "sip.meeting"
  description = "SIP endpoint for enterprise calls"
  conference  = data.pexip_infinity_conference.enterprise_meeting.id
}
