resource "pexip_infinity_event_sink" "bulk_sink" {
  name                   = "Bulk Event Sink"
  description            = "High-volume event sink with bulk processing"
  url                    = "https://analytics.company.com/events/bulk"
  username               = "analytics-user"
  password               = var.analytics_password
  bulk_support           = true
  verify_tls_certificate = true
  version                = 2
}
