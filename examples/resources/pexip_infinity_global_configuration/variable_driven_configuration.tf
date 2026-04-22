variable "global_config" {
  type = object({
    enable_webrtc                = bool
    enable_sip                   = bool
    enable_h323                  = bool
    enable_rtmp                  = bool
    crypto_mode                  = string
    conference_create_permissions = string
    conference_creation_mode     = string
    enable_analytics             = bool
    administrator_email          = string
  })
  default = {
    enable_webrtc                = true
    enable_sip                   = true
    enable_h323                  = false
    enable_rtmp                  = true
    crypto_mode                  = "besteffort"
    conference_create_permissions = "user_admin"
    conference_creation_mode     = "per_cluster"
    enable_analytics             = true
    administrator_email          = "admin@company.com"
  }
}

resource "pexip_infinity_global_configuration" "from_variable" {
  enable_webrtc                = var.global_config.enable_webrtc
  enable_sip                   = var.global_config.enable_sip
  enable_h323                  = var.global_config.enable_h323
  enable_rtmp                  = var.global_config.enable_rtmp
  crypto_mode                  = var.global_config.crypto_mode
  conference_create_permissions = var.global_config.conference_create_permissions
  conference_creation_mode     = var.global_config.conference_creation_mode
  enable_analytics             = var.global_config.enable_analytics
  administrator_email          = var.global_config.administrator_email
}
