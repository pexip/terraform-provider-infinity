# Primary Teams proxy for production
resource "pexip_infinity_teams_proxy" "production_teams" {
  name                    = "Production Teams Proxy"
  description             = "Production Teams proxy for enterprise"
  address                 = "teams-prod.enterprise.com"
  port                    = 443
  azure_tenant            = "enterprise.onmicrosoft.com"
  eventhub_id             = "production-teams-events"
  min_number_of_instances = 10
  notifications_enabled   = true
  notifications_queue     = "prod-teams-notifications"
}

# Development Teams proxy
resource "pexip_infinity_teams_proxy" "development_teams" {
  name                    = "Development Teams Proxy"
  description             = "Development Teams proxy for testing"
  address                 = "teams-dev.enterprise.com"
  port                    = 443
  azure_tenant            = "enterprise-dev.onmicrosoft.com"
  min_number_of_instances = 1
  notifications_enabled   = false
}
