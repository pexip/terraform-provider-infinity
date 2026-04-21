resource "pexip_infinity_syslog_server" "basic_syslog" {
  address   = "syslog.company.com"
  port      = 514
  transport = "udp"
}
