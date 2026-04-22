resource "pexip_infinity_global_configuration" "development" {
  enable_webrtc    = true
  enable_sip       = true
  enable_h323      = true
  enable_rtmp      = true
  crypto_mode      = "besteffort"
  
  # Relaxed conference creation for development
  conference_create_permissions = "any_authenticated"
  conference_creation_mode      = "per_node"
  
  # Extended timeouts for testing
  guests_only_timeout       = 120
  waiting_for_chair_timeout = 30
  
  # Analytics for development insights
  enable_analytics       = true
  enable_error_reporting = true
  
  administrator_email = "dev-team@company.com"
}
