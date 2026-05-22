resource "pexip_infinity_sip_proxy" "primary_proxy" {
  name      = "Primary SIP Proxy"
  address   = "sip-proxy.example.com"
  transport = "tcp"
}
