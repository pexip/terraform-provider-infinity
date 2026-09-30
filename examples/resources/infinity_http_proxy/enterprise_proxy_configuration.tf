variable "proxy_config" {
  type = object({
    address  = string
    port     = number
    protocol = string
    username = string
  })
  default = {
    address  = "enterprise-proxy.company.com"
    port     = 8080
    protocol = "http"
    username = "pexip-infinity"
  }
}

resource "pexip_infinity_http_proxy" "enterprise" {
  name     = "Enterprise HTTP Proxy"
  address  = var.proxy_config.address
  port     = var.proxy_config.port
  protocol = var.proxy_config.protocol
  username = var.proxy_config.username
  password = var.enterprise_proxy_password
}
