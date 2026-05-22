resource "pexip_infinity_mjx_endpoint_group" "example" {
  name            = "Main Office Endpoints"
  system_location = pexip_infinity_system_location.example.id
  mjx_integration = pexip_infinity_mjx_integration.example.id
}
