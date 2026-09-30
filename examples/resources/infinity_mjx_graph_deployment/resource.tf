resource "pexip_infinity_mjx_graph_deployment" "example" {
  name            = "Corporate O365 Graph Integration"
  client_id       = "12345678-1234-1234-1234-123456789012"
  oauth_token_url = "https://login.microsoftonline.com/tenant-id/oauth2/v2.0/token"
}
