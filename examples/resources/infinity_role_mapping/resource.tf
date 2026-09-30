resource "infinity_role_mapping" "example" {
  name  = "Admin Role Mapping"
  value = "pexip-admins"
  role  = infinity_role.example.id
}
