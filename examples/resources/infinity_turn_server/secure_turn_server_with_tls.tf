resource "pexip_infinity_turn_server" "secure_turn" {
  name           = "Secure TURN Server"
  description    = "TLS-encrypted TURN server"
  address        = "turns.company.com"
  port           = 5349
  server_type    = "namepsw"
  transport_type = "tls"
  username       = "secure-turn-user"
  password       = var.secure_turn_password
}
