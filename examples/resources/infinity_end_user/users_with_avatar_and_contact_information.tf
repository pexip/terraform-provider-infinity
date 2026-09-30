resource "pexip_infinity_end_user" "executive_user" {
  primary_email_address = "ceo@company.com"
  first_name            = "Chief"
  last_name             = "Executive"
  display_name          = "CEO"
  title                 = "Chief Executive Officer"
  department            = "Executive"
  telephone_number      = "+1-555-0100"
  mobile_number         = "+1-555-0101"
  avatar_url            = "https://cdn.company.com/avatars/ceo.jpg"
  user_groups = [
    "/api/admin/configuration/v1/user_group/1/", # Executives
    "/api/admin/configuration/v1/user_group/2/", # Board Members
  ]
}
