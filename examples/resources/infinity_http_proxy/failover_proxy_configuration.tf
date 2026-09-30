# Primary proxy
resource "pexip_infinity_http_proxy" "primary" {
  name     = "Primary HTTP Proxy"
  address  = "proxy1.company.com"
  port     = 8080
  protocol = "http"
  username = "pexip-primary"
  password = var.primary_proxy_password
}

# Secondary proxy for failover
resource "pexip_infinity_http_proxy" "secondary" {
  name     = "Secondary HTTP Proxy"
  address  = "proxy2.company.com"
  port     = 8080
  protocol = "http"
  username = "pexip-secondary"
  password = var.secondary_proxy_password
}
