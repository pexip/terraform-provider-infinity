resource "pexip_infinity_stun_server" "corporate_stun" {
  name        = "Corporate STUN Server"
  description = "Internal STUN server for NAT traversal"
  address     = "stun.company.com"
  port        = 3478
}
