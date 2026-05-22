resource "pexip_infinity_teams_proxy" "ha_teams_proxy" {
  name                    = "HA Teams Proxy"
  description             = "High availability Teams proxy configuration"
  address                 = "teams-ha.example.com"
  port                    = 443
  azure_tenant            = "example.onmicrosoft.com"
  min_number_of_instances = 5
  notifications_enabled   = true
}
