# Different syslog servers per region
resource "pexip_infinity_syslog_server" "regional_syslog" {
  for_each = var.regional_syslog_config
  
  address     = each.value.address
  description = "Regional syslog server for ${each.key}"
  port        = each.value.port
  transport   = each.value.transport
  audit_log   = true
  support_log = true
  web_log     = false
}
