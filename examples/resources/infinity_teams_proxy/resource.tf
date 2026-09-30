resource "pexip_infinity_teams_proxy" "primary_teams_proxy" {
  name                    = "Primary Teams Proxy"
  address                 = "teams-proxy.example.com"
  port                    = 443
  azure_tenant            = "contoso.onmicrosoft.com"
  min_number_of_instances = 2
}
