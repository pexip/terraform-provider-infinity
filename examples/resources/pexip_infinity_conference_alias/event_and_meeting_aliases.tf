# All-hands meeting
resource "pexip_infinity_conference_alias" "all_hands" {
  alias       = "all-hands"
  description = "Monthly all-hands meeting"
  conference  = data.pexip_infinity_conference.all_hands_meeting.id
}

# Board meeting
resource "pexip_infinity_conference_alias" "board_meeting" {
  alias       = "board"
  description = "Executive board meeting"
  conference  = data.pexip_infinity_conference.executive_board.id
}

# Customer event
resource "pexip_infinity_conference_alias" "customer_event" {
  alias       = "customer-event-2024"
  description = "Annual customer event 2024"
  conference  = data.pexip_infinity_conference.annual_event.id
}
