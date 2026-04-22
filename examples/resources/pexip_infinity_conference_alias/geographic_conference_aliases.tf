# US East Coast conference room
resource "pexip_infinity_conference_alias" "us_east_room" {
  alias       = "us-east.conference"
  description = "US East Coast conference room"
  conference  = data.pexip_infinity_conference.us_east.id
}

# US West Coast conference room
resource "pexip_infinity_conference_alias" "us_west_room" {
  alias       = "us-west.conference"
  description = "US West Coast conference room"
  conference  = data.pexip_infinity_conference.us_west.id
}

# European conference room
resource "pexip_infinity_conference_alias" "europe_room" {
  alias       = "europe.conference"
  description = "European conference room"
  conference  = data.pexip_infinity_conference.europe.id
}

# Asia Pacific conference room
resource "pexip_infinity_conference_alias" "apac_room" {
  alias       = "apac.conference"
  description = "Asia Pacific conference room"
  conference  = data.pexip_infinity_conference.apac.id
}
