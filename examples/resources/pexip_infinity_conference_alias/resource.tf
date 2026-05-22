resource "pexip_infinity_conference_alias" "meeting_room_alias" {
  alias       = "boardroom"
  description = "Alias for main boardroom conference"
  conference  = pexip_infinity_conference.example.id
}
