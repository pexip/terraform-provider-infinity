resource "pexip_infinity_turn_server" "shared_secret_turn" {
  name           = "Shared Secret TURN"
  description    = "TURN server using shared secret authentication"
  address        = "turn-shared.company.com"
  port           = 3478
  server_type    = "coturn_shared"
  transport_type = "udp"
  secret_key     = var.turn_shared_secret
}
