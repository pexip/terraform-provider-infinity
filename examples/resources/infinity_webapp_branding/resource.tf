resource "pexip_infinity_webapp_branding" "example" {
  name          = "Corporate Branding"
  description   = "Corporate branding for Pexip web applications"
  webapp_type   = "webapp1"
  is_default    = true
  branding_file = "/path/to/branding/package.zip"
}
