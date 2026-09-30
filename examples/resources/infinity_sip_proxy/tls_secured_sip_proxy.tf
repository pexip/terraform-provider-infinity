resource "pexip_infinity_sip_proxy" "secure_proxy" {
  name        = "Secure SIP Proxy"
  description = "TLS-secured SIP proxy for sensitive communications"
  address     = "secure-sip.example.com"
  port        = 5061
  transport   = "tls"
}
