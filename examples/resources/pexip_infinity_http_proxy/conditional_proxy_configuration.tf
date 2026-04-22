locals {
  use_proxy = var.environment == "production" || var.environment == "staging"
}

resource "pexip_infinity_http_proxy" "conditional" {
  count    = local.use_proxy ? 1 : 0
  name     = "${title(var.environment)} HTTP Proxy"
  address  = var.proxy_address
  port     = var.proxy_port
  protocol = var.proxy_protocol
  username = var.proxy_username
  password = var.proxy_password
}
