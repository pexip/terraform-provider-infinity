resource "pexip_infinity_static_route" "default_route" {
  name    = "Custom Default Route"
  address = "0.0.0.0"
  prefix  = 0
  gateway = "192.168.1.254"
}
