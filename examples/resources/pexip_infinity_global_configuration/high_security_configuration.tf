resource "pexip_infinity_global_configuration" "high_security" {
  enable_webrtc    = true
  enable_sip       = true
  enable_h323      = false
  enable_rtmp      = false
  crypto_mode      = "required"
  
  # Restricted conference creation
  conference_create_permissions = "admin_only"
  conference_creation_mode      = "per_cluster"
  
  # Short timeouts for security
  guests_only_timeout       = 15
  waiting_for_chair_timeout = 5
  
  # Network restrictions
  bandwidth_restrictions = "restricted"
  media_ports_start      = 50000
  media_ports_end        = 50999
  signalling_ports_start = 5060
  signalling_ports_end   = 5061
  
  # Security monitoring
  enable_analytics       = true
  enable_error_reporting = false
  
  administrator_email = "security-admin@company.com"
  
  # Restrict conference creation to specific groups
  global_conference_create_groups = [
    "CN=Video Admins,OU=Groups,DC=company,DC=com",
    "CN=IT Managers,OU=Groups,DC=company,DC=com"
  ]
}
