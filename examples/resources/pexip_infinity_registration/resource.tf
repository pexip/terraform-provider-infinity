resource "pexip_infinity_registration" "example" {
  enabled = true
  registration_interval = 3600
  refresh_interval = 60
  refresh_strategy = "maximum"
}
