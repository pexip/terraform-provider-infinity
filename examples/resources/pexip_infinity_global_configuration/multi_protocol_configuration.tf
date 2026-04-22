resource "pexip_infinity_global_configuration" "multi_protocol" {
  # Enable all protocols
  enable_webrtc = true
  enable_sip    = true
  enable_h323   = true
  enable_rtmp   = true
  
  # Security settings
  crypto_mode = "besteffort"
  
  # Optimized port ranges for different protocols
  media_ports_start      = 32768
  media_ports_end        = 65535
  signalling_ports_start = 1024
  signalling_ports_end   = 65535
  
  # Video quality settings
  max_pixels_per_second = "3840x2160x30"  # 4K support
  
  # Conference management
  conference_create_permissions = "user_admin"
  conference_creation_mode      = "per_cluster"
  
  # Operational settings
  enable_analytics       = true
  enable_error_reporting = true
  bandwidth_restrictions = "none"
  
  administrator_email = "operations@company.com"
}
