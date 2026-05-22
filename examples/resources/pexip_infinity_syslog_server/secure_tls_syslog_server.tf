resource "pexip_infinity_syslog_server" "secure_syslog" {
  address     = "secure-logs.company.com"
  description = "Secure syslog server with TLS encryption"
  port        = 6514
  transport   = "tls"
  audit_log   = true
  support_log = false
  web_log     = false
}
