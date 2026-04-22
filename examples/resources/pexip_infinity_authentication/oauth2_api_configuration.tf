resource "pexip_infinity_authentication" "api_config" {
  source = "local"
  
  # OAuth2 API settings
  api_oauth2_disable_basic    = true
  api_oauth2_allow_all_perms  = false
  api_oauth2_expiration       = 7200  # 2 hours
}
