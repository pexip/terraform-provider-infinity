# Internal SIP proxy for corporate network
resource "pexip_infinity_sip_proxy" "internal_sip_proxy" {
  name        = "Internal SIP Proxy"
  description = "Internal corporate SIP proxy server"
  address     = "10.1.1.50"
  port        = 5060
  transport   = "tcp"
}

# External SIP proxy for internet traffic
resource "pexip_infinity_sip_proxy" "external_sip_proxy" {
  name        = "External SIP Proxy"
  description = "External SIP proxy for internet-facing traffic"
  address     = "sip.company.com"
  port        = 5061
  transport   = "tls"
}

# SIP proxy for legacy systems
resource "pexip_infinity_sip_proxy" "legacy_sip_proxy" {
  name        = "Legacy SIP Proxy"
  description = "SIP proxy for legacy phone systems"
  address     = "legacy-pbx.internal.com"
  port        = 5060
  transport   = "udp"
}
