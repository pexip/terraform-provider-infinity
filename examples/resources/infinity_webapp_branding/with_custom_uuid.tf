resource "pexip_infinity_webapp_branding" "example_custom_uuid" {
  name          = "Corporate Branding"
  description   = "Corporate branding for Pexip web applications"
  uuid          = "12345678-1234-1234-1234-123456789012"
  webapp_type   = "webapp2"
  is_default    = true
  branding_file = "/path/to/branding/package.zip"
}
