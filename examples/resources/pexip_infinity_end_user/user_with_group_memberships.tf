# Assume user groups are defined elsewhere
resource "pexip_infinity_end_user" "manager_user" {
  primary_email_address = "manager@company.com"
  first_name            = "John"
  last_name             = "Manager"
  display_name          = "John Manager"
  title                 = "Engineering Manager"
  department            = "Engineering"
  user_groups          = [
    "/api/admin/configuration/v1/user_group/1/",  # Managers group
    "/api/admin/configuration/v1/user_group/3/",  # Engineering group
  ]
}
