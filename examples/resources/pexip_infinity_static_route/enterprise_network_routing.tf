variable "static_routes" {
  type = list(object({
    name    = string
    address = string
    prefix  = number
    gateway = string
  }))
  default = [
    {
      name    = "Branch Office 1"
      address = "10.10.0.0"
      prefix  = 16
      gateway = "192.168.1.10"
    },
    {
      name    = "Branch Office 2"
      address = "10.20.0.0"
      prefix  = 16
      gateway = "192.168.1.11"
    },
    {
      name    = "Data Center"
      address = "172.16.0.0"
      prefix  = 12
      gateway = "192.168.1.20"
    }
  ]
}

resource "pexip_infinity_static_route" "enterprise_routes" {
  count   = length(var.static_routes)
  name    = var.static_routes[count.index].name
  address = var.static_routes[count.index].address
  prefix  = var.static_routes[count.index].prefix
  gateway = var.static_routes[count.index].gateway
}
