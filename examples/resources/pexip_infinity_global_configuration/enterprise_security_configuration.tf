resource "pexip_infinity_global_configuration" "enterprise" {
  enable_webrtc    = true
  enable_sip       = true
  enable_h323      = true
  enable_rtmp      = true
  crypto_mode      = "required"
  
  # Network configuration
  media_ports_start       = 40000
  media_ports_end         = 49999
  signalling_ports_start  = 5060
  signalling_ports_end    = 5080
  
  # Conference settings
  conference_create_permissions = "user_admin"
  conference_creation_mode      = "per_cluster"
  guests_only_timeout          = 30
  waiting_for_chair_timeout    = 10
  
  # Security and compliance
  enable_analytics      = true
  enable_error_reporting = false
  bandwidth_restrictions = "restricted"
  
  administrator_email = "admin@company.com"
}
