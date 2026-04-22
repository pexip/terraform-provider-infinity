resource "pexip_infinity_policy_server" "comprehensive_policy" {
  name        = "Comprehensive Policy Server"
  description = "Full-featured policy server with all capabilities"
  url         = "https://policy-engine.company.com/api/v2"
  username    = "pexip-integration"
  password    = var.comprehensive_policy_password
  
  # Enable all lookup types
  enable_service_lookup             = true
  enable_participant_lookup         = true
  enable_registration_lookup        = true
  enable_directory_lookup           = true
  enable_avatar_lookup              = true
  enable_media_location_lookup      = true
  
  # Enable internal policies
  enable_internal_service_policy       = true
  enable_internal_participant_policy   = true
  enable_internal_media_location_policy = true
  
  # Avatar preferences
  prefer_local_avatar_configuration = false
  
  # Configuration templates
  service_configuration_template = file("${path.module}/templates/service-policy.json")
  participant_configuration_template = file("${path.module}/templates/participant-policy.json")
  media_location_configuration_template = file("${path.module}/templates/media-location-policy.json")
}
