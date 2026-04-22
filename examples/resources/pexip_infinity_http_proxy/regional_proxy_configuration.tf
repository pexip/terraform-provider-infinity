variable "regional_proxies" {
  type = map(object({
    name     = string
    address  = string
    port     = number
    username = string
  }))
  default = {
    us-east = {
      name     = "US East HTTP Proxy"
      address  = "proxy-useast.company.com"
      port     = 8080
      username = "useast-pexip"
    }
    us-west = {
      name     = "US West HTTP Proxy"
      address  = "proxy-uswest.company.com"
      port     = 8080
      username = "uswest-pexip"
    }
    europe = {
      name     = "Europe HTTP Proxy"
      address  = "proxy-eu.company.com"
      port     = 8080
      username = "eu-pexip"
    }
  }
}

resource "pexip_infinity_http_proxy" "regional" {
  for_each = var.regional_proxies
  name     = each.value.name
  address  = each.value.address
  port     = each.value.port
  protocol = "http"
  username = each.value.username
  password = var.regional_proxy_passwords[each.key]
}
