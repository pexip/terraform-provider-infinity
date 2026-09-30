# Audit logs to compliance system
resource "pexip_infinity_syslog_server" "audit_logs" {
  address     = "audit.company.com"
  description = "Compliance audit logging system"
  port        = 514
  transport   = "tcp"
  audit_log   = true
  support_log = false
  web_log     = false
}

# Support logs to monitoring system
resource "pexip_infinity_syslog_server" "support_logs" {
  address     = "monitoring.company.com"
  description = "Technical support and monitoring logs"
  port        = 514
  transport   = "udp"
  audit_log   = false
  support_log = true
  web_log     = true
}

# High-priority logs to SIEM
resource "pexip_infinity_syslog_server" "siem_logs" {
  address     = "siem.company.com"
  description = "Security Information and Event Management"
  port        = 1514
  transport   = "tls"
  audit_log   = true
  support_log = true
  web_log     = true
}
