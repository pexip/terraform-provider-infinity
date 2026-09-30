resource "pexip_infinity_oauth2_client" "example" {
  client_name = "API Automation Client"
  role        = "/configuration/v1/system_location/1/role/1/"
}
