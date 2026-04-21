resource "pexip_infinity_scheduled_scaling" "example" {
  policy_name         = "Peak Hours Scaling"
  policy_type         = "TeamsConnectorScaling"
  resource_identifier = pexip_infinity_teams_proxy.example.id
  start_date          = "2024-01-01"
  time_from           = "08:00:00"
  time_to             = "18:00:00"
}
