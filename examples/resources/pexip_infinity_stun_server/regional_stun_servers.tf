# STUN servers for different regions
locals {
  regional_stun_servers = {
    "us-east" = {
      address = "stun-us-east.company.com"
      port    = 3478
    }
    "us-west" = {
      address = "stun-us-west.company.com"
      port    = 3478
    }
    "europe" = {
      address = "stun-eu.company.com"
      port    = 3478
    }
    "asia" = {
      address = "stun-asia.company.com"
      port    = 3478
    }
  }
}

resource "pexip_infinity_stun_server" "regional_stun" {
  for_each = local.regional_stun_servers

  name        = "STUN Server - ${each.key}"
  description = "Regional STUN server for ${each.key}"
  address     = each.value.address
  port        = each.value.port
}
