# Only create routes if gateway is available
locals {
  create_routes = var.gateway_available
}

resource "pexip_infinity_static_route" "conditional_route" {
  count   = local.create_routes ? 1 : 0
  name    = "Conditional Network Route"
  address = "10.50.0.0"
  prefix  = 16
  gateway = var.conditional_gateway
}
