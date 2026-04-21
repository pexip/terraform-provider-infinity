resource "pexip_infinity_oauth2_client" "example" {
  client_name = "API Automation Client"
  role        = pexip_infinity_role.example.id
}
