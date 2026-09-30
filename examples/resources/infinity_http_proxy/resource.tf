resource "pexip_infinity_http_proxy" "corporate_proxy" {
  name     = "Corporate HTTP Proxy"
  address  = "proxy.company.com"
  port     = 8080
  protocol = "http"
}
