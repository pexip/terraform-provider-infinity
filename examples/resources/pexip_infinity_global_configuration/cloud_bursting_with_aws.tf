resource "pexip_infinity_global_configuration" "aws_bursting" {
  enable_webrtc    = true
  enable_sip       = true
  crypto_mode      = "besteffort"
  
  # Cloud bursting configuration
  bursting_enabled = true
  cloud_provider   = "aws"
  aws_access_key   = var.aws_access_key
  aws_secret_key   = var.aws_secret_key
  
  # Conference settings
  conference_create_permissions = "any_authenticated"
  conference_creation_mode      = "per_cluster"
  
  administrator_email = "pexip-admin@company.com"
}
