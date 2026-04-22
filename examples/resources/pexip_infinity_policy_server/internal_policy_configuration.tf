resource "pexip_infinity_policy_server" "internal_policy" {
  name                               = "Internal Policy Engine"
  description                        = "Internal policy with custom templates"
  url                                = "https://internal-policy.company.com"
  
  # Enable internal policy features
  enable_internal_service_policy     = true
  enable_internal_participant_policy = true
  enable_internal_media_location_policy = true
  
  # Custom policy templates
  service_configuration_template     = jsonencode({
    default_bandwidth = "2048"
    recording_enabled = true
    encryption_required = true
  })
  
  participant_configuration_template = jsonencode({
    default_role = "guest"
    max_participants = 50
    allow_dial_out = false
  })
  
  media_location_configuration_template = jsonencode({
    preferred_location = "primary"
    failover_enabled = true
    quality_threshold = 0.8
  })
}
