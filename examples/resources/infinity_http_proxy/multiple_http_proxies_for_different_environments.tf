# Production proxy
resource "pexip_infinity_http_proxy" "production" {
  name     = "Production HTTP Proxy"
  address  = "proxy-prod.company.com"
  port     = 3128
  protocol = "http"
  username = "prod-pexip"
  password = var.prod_proxy_password
}

# Development proxy
resource "pexip_infinity_http_proxy" "development" {
  name     = "Development HTTP Proxy"
  address  = "proxy-dev.company.com"
  port     = 3128
  protocol = "http"
  username = "dev-pexip"
  password = var.dev_proxy_password
}

# Test proxy without authentication
resource "pexip_infinity_http_proxy" "test" {
  name     = "Test HTTP Proxy"
  address  = "proxy-test.company.com"
  port     = 8080
  protocol = "http"
}
