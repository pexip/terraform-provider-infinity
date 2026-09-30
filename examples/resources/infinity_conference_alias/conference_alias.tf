resource "pexip_infinity_conference_alias" "example_sip_alias" {
  alias       = "meetingroom@example.com"
  description = "SIP alias"
  conference  = pexip_infinity_conference.example.id
}

resource "pexip_infinity_conference_alias" "example_alias" {
  alias       = "meetingroom"
  description = "Name only alias"
  conference  = pexip_infinity_conference.example.id
}

resource "pexip_infinity_conference" "example" {
  name = "Example Conference"
}