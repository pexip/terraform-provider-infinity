resource "pexip_infinity_global_configuration" "azure_integration" {
  enable_webrtc    = true
  enable_sip       = true
  enable_h323      = false
  enable_rtmp      = true
  crypto_mode      = "required"
  
  # Azure cloud bursting
  bursting_enabled = true
  cloud_provider   = "azure"
  azure_client_id  = var.azure_client_id
  azure_secret     = var.azure_secret
  
  # Media configuration
  max_pixels_per_second = "1920x1080x30"
  media_ports_start     = 36000
  media_ports_end       = 59999
  
  # Administrative settings
  enable_analytics       = true
  enable_error_reporting = true
  administrator_email    = "video-admin@company.com"
}
