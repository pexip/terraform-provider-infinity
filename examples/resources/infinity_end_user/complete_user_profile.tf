resource "pexip_infinity_end_user" "jane_smith" {
  primary_email_address = "jane.smith@company.com"
  first_name            = "Jane"
  last_name             = "Smith"
  display_name          = "Jane Smith"
  telephone_number      = "+1-555-0123"
  mobile_number         = "+1-555-0124"
  title                 = "Senior Engineer"
  department            = "Engineering"
  avatar_url            = "https://avatars.company.com/jane.smith.jpg"
}
