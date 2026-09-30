resource "pexip_infinity_static_route" "specific_server" {
  name    = "Database Server Route"
  address = "10.1.1.100"
  prefix  = 32
  gateway = "192.168.1.10"
}
