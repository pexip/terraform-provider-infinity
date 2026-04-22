resource "pexip_infinity_authentication" "oidc_jwt_auth" {
  source = "oidc"
  
  # OIDC configuration with JWT
  oidc_client_id        = var.oidc_client_id
  oidc_private_key      = var.oidc_private_key
  oidc_auth_method      = "private_key_jwt"
  oidc_authorize_url    = "https://auth.company.com/oauth2/authorize"
  oidc_token_endpoint_url = "https://auth.company.com/oauth2/token"
  oidc_scope           = "openid profile email groups"
  
  # Claim configuration
  oidc_username_field   = "sub"
  oidc_groups_field     = "department_groups"
}
