resource "pexip_infinity_static_route" "remote_office" {
  name    = "Remote Office Network"
  address = "172.16.0.0"
  prefix  = 12
  gateway = "10.0.1.254"
}
