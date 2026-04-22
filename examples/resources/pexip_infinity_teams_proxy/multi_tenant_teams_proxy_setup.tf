variable "teams_tenants" {
  type = list(object({
    name         = string
    azure_tenant = string
    instances    = number
  }))
  default = [
    {
      name         = "Tenant A Teams Proxy"
      azure_tenant = "tenanta.onmicrosoft.com"
      instances    = 2
    },
    {
      name         = "Tenant B Teams Proxy"
      azure_tenant = "tenantb.onmicrosoft.com"
      instances    = 3
    }
  ]
}

resource "pexip_infinity_teams_proxy" "multi_tenant" {
  count                    = length(var.teams_tenants)
  name                     = var.teams_tenants[count.index].name
  description              = "Teams proxy for ${var.teams_tenants[count.index].azure_tenant}"
  address                  = "teams-proxy-${count.index + 1}.example.com"
  port                     = 443
  azure_tenant             = var.teams_tenants[count.index].azure_tenant
  min_number_of_instances  = var.teams_tenants[count.index].instances
  notifications_enabled    = true
}
