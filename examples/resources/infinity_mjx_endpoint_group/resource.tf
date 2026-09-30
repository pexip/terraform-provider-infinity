resource "infinity_mjx_endpoint_group" "example" {
  name            = "Main Office Endpoints"
  system_location = infinity_system_location.example.id
  mjx_integration = infinity_mjx_integration.example.id
}
