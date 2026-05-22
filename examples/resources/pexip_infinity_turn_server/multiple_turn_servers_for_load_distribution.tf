# Primary TURN server
resource "pexip_infinity_turn_server" "turn_primary" {
  name           = "Primary TURN Server"
  description    = "Primary TURN server for main office"
  address        = "turn1.company.com"
  port           = 3478
  server_type    = "namepsw"
  transport_type = "udp"
  username       = "primary-turn"
  password       = var.primary_turn_password
}

# Secondary TURN server
resource "pexip_infinity_turn_server" "turn_secondary" {
  name           = "Secondary TURN Server"
  description    = "Secondary TURN server for redundancy"
  address        = "turn2.company.com"
  port           = 3478
  server_type    = "namepsw"
  transport_type = "tcp"
  username       = "secondary-turn"
  password       = var.secondary_turn_password
}
