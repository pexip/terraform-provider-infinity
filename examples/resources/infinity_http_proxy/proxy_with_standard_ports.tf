# HTTP proxy using standard port (implied port 80)
resource "pexip_infinity_http_proxy" "standard_http" {
  name     = "Standard HTTP Proxy"
  address  = "http-proxy.company.com"
  protocol = "http"
}

# HTTPS proxy using standard port (implied port 443)
resource "pexip_infinity_http_proxy" "standard_https" {
  name     = "Standard HTTPS Proxy"
  address  = "https-proxy.company.com"
  protocol = "https"
  username = "secure-user"
  password = var.secure_proxy_password
}
