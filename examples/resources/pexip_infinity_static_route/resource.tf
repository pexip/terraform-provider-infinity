resource "pexip_infinity_static_route" "basic_route" {
  name    = "Internal Network Route"
  address = "10.1.0.0"
  prefix  = 16
  gateway = "192.168.1.1"
}
