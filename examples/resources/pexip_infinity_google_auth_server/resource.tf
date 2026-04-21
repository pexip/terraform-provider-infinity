resource "pexip_infinity_google_auth_server" "example" {
  name             = "Google OAuth Client"
  application_type = "WEB_APP"
  client_id        = "123456789-abcdef.apps.googleusercontent.com"
}
