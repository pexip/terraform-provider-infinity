resource "pexip_infinity_teams_proxy" "teams_proxy_with_events" {
  name                     = "Teams Proxy with Events"
  description              = "Teams proxy with event hub integration"
  address                  = "teams-proxy.company.com"
  port                     = 443
  azure_tenant             = "company.onmicrosoft.com"
  eventhub_id              = "teams-events-hub"
  min_number_of_instances  = 3
  notifications_enabled    = true
  notifications_queue      = "teams-notifications-queue"
}
