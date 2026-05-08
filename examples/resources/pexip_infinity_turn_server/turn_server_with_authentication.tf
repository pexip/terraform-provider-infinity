resource "pexip_infinity_turn_server" "authenticated_turn" {
  name           = "Corporate TURN Server"
  description    = "TURN server with username/password authentication"
  address        = "turn.company.com"
  port           = 3478
  server_type    = "namepsw"
  transport_type = "udp"
  username       = "turnuser"
  password       = var.turn_password
}
