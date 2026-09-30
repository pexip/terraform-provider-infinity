variable "event_sinks" {
  type = list(object({
    name                   = string
    description            = string
    url                    = string
    username               = string
    bulk_support           = bool
    verify_tls_certificate = bool
    version                = number
  }))
  default = [
    {
      name                   = "SIEM Integration"
      description            = "Security information and event management"
      url                    = "https://siem.company.com/api/events"
      username               = "siem-integration"
      bulk_support           = true
      verify_tls_certificate = true
      version                = 2
    },
    {
      name                   = "Data Lake"
      description            = "Long-term storage for analytics"
      url                    = "https://datalake.company.com/ingest/pexip"
      username               = "datalake-service"
      bulk_support           = true
      verify_tls_certificate = true
      version                = 2
    }
  ]
}

resource "pexip_infinity_event_sink" "enterprise" {
  count                  = length(var.event_sinks)
  name                   = var.event_sinks[count.index].name
  description            = var.event_sinks[count.index].description
  url                    = var.event_sinks[count.index].url
  username               = var.event_sinks[count.index].username
  password               = var.event_sink_passwords[count.index]
  bulk_support           = var.event_sinks[count.index].bulk_support
  verify_tls_certificate = var.event_sinks[count.index].verify_tls_certificate
  version                = var.event_sinks[count.index].version
}
