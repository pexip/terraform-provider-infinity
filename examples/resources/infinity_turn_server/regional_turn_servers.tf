# TURN servers for different regions
resource "pexip_infinity_turn_server" "regional_turn" {
  for_each = var.regional_turn_servers

  name           = "TURN Server - ${each.key}"
  description    = "Regional TURN server for ${each.key}"
  address        = each.value.address
  port           = each.value.port
  server_type    = "namepsw"
  transport_type = each.value.transport
  username       = each.value.username
  password       = each.value.password
}
