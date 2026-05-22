# Create multiple users from variable
resource "pexip_infinity_end_user" "department_users" {
  for_each = var.department_users

  primary_email_address = each.value.email
  first_name            = each.value.first_name
  last_name             = each.value.last_name
  display_name          = "${each.value.first_name} ${each.value.last_name}"
  telephone_number      = each.value.phone
  department            = each.value.department
  title                 = each.value.title
  sync_tag              = "ldap-import"
}
