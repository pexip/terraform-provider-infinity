resource "pexip_infinity_role_mapping" "example" {
  name  = "Admin Role Mapping"
  value = "pexip-admins"
  role  = pexip_infinity_role.example.id
}
