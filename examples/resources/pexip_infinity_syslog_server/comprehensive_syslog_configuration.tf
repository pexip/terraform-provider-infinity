resource "pexip_infinity_syslog_server" "central_logging" {
  address     = "logs.company.com"
  description = "Central logging server for all Pexip logs"
  port        = 514
  transport   = "udp"
  audit_log   = true
  support_log = true
  web_log     = true
}
