resource "pexip_infinity_turn_server" "public_turn" {
  name           = "Public TURN Service"
  description    = "External TURN service provider"
  address        = "global-turn.example.com"
  port           = 3478
  server_type    = "namepsw"
  transport_type = "udp"
  username       = var.public_turn_username
  password       = var.public_turn_password
}
