variable "sip_proxy_servers" {
  type = list(object({
    name      = string
    address   = string
    port      = number
    transport = string
  }))
  default = [
    {
      name      = "SIP Proxy 1"
      address   = "sip-proxy-1.example.com"
      port      = 5060
      transport = "tcp"
    },
    {
      name      = "SIP Proxy 2"
      address   = "sip-proxy-2.example.com"
      port      = 5060
      transport = "tcp"
    }
  ]
}

resource "pexip_infinity_sip_proxy" "load_balanced_proxies" {
  count       = length(var.sip_proxy_servers)
  name        = var.sip_proxy_servers[count.index].name
  description = "Load balanced SIP proxy ${count.index + 1}"
  address     = var.sip_proxy_servers[count.index].address
  port        = var.sip_proxy_servers[count.index].port
  transport   = var.sip_proxy_servers[count.index].transport
}
