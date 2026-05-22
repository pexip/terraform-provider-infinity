resource "pexip_infinity_sip_proxy" "custom_port_proxy" {
  name        = "Custom Port SIP Proxy"
  description = "SIP proxy with custom port configuration"
  address     = "192.168.1.100"
  port        = 5080
  transport   = "udp"
}
