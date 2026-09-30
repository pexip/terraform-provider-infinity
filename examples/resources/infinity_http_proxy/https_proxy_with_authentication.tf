resource "pexip_infinity_http_proxy" "secure_proxy" {
  name     = "Secure HTTP Proxy"
  address  = "secure-proxy.company.com"
  port     = 8443
  protocol = "https"
  username = "pexip-service"
  password = var.proxy_password
}
