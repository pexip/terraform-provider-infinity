resource "pexip_infinity_ldap_sync_field" "example" {
  name                   = "Department Field"
  template_variable_name = "department"
  ldap_attribute_name    = "department"
}
