# Analytics event sink
resource "pexip_infinity_event_sink" "analytics" {
  name                   = "Analytics Event Sink"
  description            = "Event sink for call analytics and reporting"
  url                    = "https://analytics.company.com/pexip/events"
  username               = "analytics-service"
  password               = var.analytics_password
  bulk_support           = true
  verify_tls_certificate = true
  version                = 2
}

# Monitoring event sink
resource "pexip_infinity_event_sink" "monitoring" {
  name        = "Monitoring Event Sink"
  description = "Event sink for real-time monitoring"
  url         = "https://monitoring.company.com/webhooks/pexip"
  username    = "monitoring-user"
  password    = var.monitoring_password
  version     = 1
}

# Billing event sink
resource "pexip_infinity_event_sink" "billing" {
  name                   = "Billing Event Sink"
  description            = "Event sink for usage tracking and billing"
  url                    = "https://billing.company.com/api/usage-events"
  username               = "billing-service"
  password               = var.billing_password
  bulk_support           = true
  verify_tls_certificate = true
  version                = 2
}
