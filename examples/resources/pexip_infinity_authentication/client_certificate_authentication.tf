resource "pexip_infinity_authentication" "cert_auth" {
  source             = "local"
  client_certificate = "required"
  
  # Combined with OAuth2 API settings
  api_oauth2_disable_basic = true
  api_oauth2_expiration    = 3600
}
