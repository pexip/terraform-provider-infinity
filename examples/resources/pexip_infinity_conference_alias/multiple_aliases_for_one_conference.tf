# Primary conference
data "pexip_infinity_conference" "weekly_standup" {
  name = "Weekly Team Standup"
}

# Short alias for easy access
resource "pexip_infinity_conference_alias" "standup_short" {
  alias       = "standup"
  description = "Short alias for weekly standup"
  conference  = data.pexip_infinity_conference.weekly_standup.id
}

# Department-specific alias
resource "pexip_infinity_conference_alias" "dev_standup" {
  alias       = "dev.standup"
  description = "Development team standup alias"
  conference  = data.pexip_infinity_conference.weekly_standup.id
}

# Numeric alias for phone users
resource "pexip_infinity_conference_alias" "standup_number" {
  alias       = "12345"
  description = "Numeric alias for phone access"
  conference  = data.pexip_infinity_conference.weekly_standup.id
}
