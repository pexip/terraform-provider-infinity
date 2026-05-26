resource "pexip_infinity_authentication" "pkce_example" {
  source                  = "OIDC+LOCAL"
  oidc_auth_method        = "pkce"
  oidc_client_id          = var.oidc_client_id
  oidc_metadata_url       = "https://login.microsoftonline.com/${var.azure_tenant_id}/v2.0/.well-known/openid-configuration"
  oidc_authorize_url      = "https://login.microsoftonline.com/${var.azure_tenant_id}/oauth2/v2.0/authorize"
  oidc_token_endpoint_url = "https://login.microsoftonline.com/${var.azure_tenant_id}/oauth2/v2.0/token"
  oidc_login_button       = "Login with Entra (PKCE)"
}
