resource "pexip_infinity_conference_alias" "meeting_room_alias" {
  alias       = "boardroom"
  description = "Alias for main boardroom conference"
  conference  = data.pexip_infinity_conference.main_boardroom.id
}
