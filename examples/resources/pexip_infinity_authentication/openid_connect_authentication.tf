resource "pexip_infinity_authentication" "oidc_auth" {
  source = "oidc"
  
  # OIDC provider configuration
  oidc_metadata_url     = "https://login.microsoftonline.com/tenant-id/.well-known/openid_configuration"
  oidc_client_id        = var.oidc_client_id
  oidc_client_secret    = var.oidc_client_secret
  oidc_auth_method      = "client_secret_basic"
  oidc_scope           = "openid profile email"
  
  # Claim mapping
  oidc_username_field   = "preferred_username"
  oidc_groups_field     = "groups"
  
  # Access control
  oidc_required_key     = "department"
  oidc_required_value   = "IT"
  
  # UI customization
  oidc_domain_hint      = "company.com"
  oidc_login_button     = "Sign in with Corporate SSO"
}
