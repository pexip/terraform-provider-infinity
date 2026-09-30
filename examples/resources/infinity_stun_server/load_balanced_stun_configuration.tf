# STUN servers behind load balancer
resource "pexip_infinity_stun_server" "lb_stun" {
  count = var.stun_server_count

  name        = "Load Balanced STUN ${count.index + 1}"
  description = "STUN server ${count.index + 1} behind load balancer"
  address     = "stun-lb.company.com"
  port        = 3478 + count.index
}
