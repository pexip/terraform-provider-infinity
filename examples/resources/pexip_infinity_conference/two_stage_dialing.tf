resource "pexip_infinity_conference" "ivr" {
  name         = "IVR Service"
  service_type = "two_stage_dialing"
  description  = "IVR example"
  match_string = "^\\d{4}$"
  aliases = [
    { alias = "exampleivr" },
    { alias = "exampleivr@example.com" },
  ]
}