resource "pexip_infinity_mjx_meeting_processing_rule" "example" {
  name            = "Teams Meeting Rule"
  priority        = 10
  meeting_type    = "teams"
  mjx_integration = pexip_infinity_mjx_integration.example.id
}
