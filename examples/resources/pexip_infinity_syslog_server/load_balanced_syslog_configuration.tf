resource "pexip_infinity_syslog_server" "lb_syslog" {
  count = length(var.syslog_servers)
  
  address     = var.syslog_servers[count.index].address
  description = "Load-balanced syslog server ${count.index + 1}"
  port        = var.syslog_servers[count.index].port
  transport   = "udp"
  audit_log   = true
  support_log = true
  web_log     = true
}
