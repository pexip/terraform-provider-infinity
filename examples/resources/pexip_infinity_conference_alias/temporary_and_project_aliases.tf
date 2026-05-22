# Project-specific conference
resource "pexip_infinity_conference_alias" "project_alpha" {
  alias       = "project-alpha"
  description = "Project Alpha team meetings"
  conference  = data.pexip_infinity_conference.project_room.id
}

# Temporary event alias
resource "pexip_infinity_conference_alias" "temp_event" {
  alias       = "temp-${formatdate("YYYY-MM-DD", timestamp())}"
  description = "Temporary conference for today's event"
  conference  = data.pexip_infinity_conference.temp_conference.id
}

# Contractor access
resource "pexip_infinity_conference_alias" "contractor" {
  alias       = "contractor-access"
  description = "External contractor conference access"
  conference  = data.pexip_infinity_conference.contractor_room.id
}
