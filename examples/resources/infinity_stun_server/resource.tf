resource "pexip_infinity_stun_server" "basic_stun" {
  name    = "Basic STUN Server"
  address = "stun.company.com"
  port    = 3478
}
