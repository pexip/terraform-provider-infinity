resource "pexip_infinity_gateway_routing_rule" "example" {
  name               = "SIP Gateway Rule"
  description        = "Route calls to external SIP gateway"
  priority           = 100
  enable             = true
  match_string       = "^\\+1(.*)$"
  replace_string     = "$1"
  called_device_type = "gateway"
  outgoing_protocol  = "sip"
  call_type          = "video"
}
