# Branch office with minimal configuration
resource "pexip_infinity_system_location" "branch_office" {
  name        = "Branch Office"
  description = "Remote branch office location"
  mtu         = 1400  # Lower MTU for WAN links
  dns_servers = [
    pexip_infinity_dns_server.primary_dns.id
  ]
  ntp_servers = [
    pexip_infinity_ntp_server.primary_ntp.id
  ]
}

# Data center location with dedicated services
resource "pexip_infinity_system_location" "datacenter" {
  name        = "Data Center"
  description = "Primary data center location"
  mtu         = 9000  # Jumbo frames for internal networks
  dns_servers = [
    pexip_infinity_dns_server.primary_dns.id,
    pexip_infinity_dns_server.secondary_dns.id
  ]
  ntp_servers = [
    pexip_infinity_ntp_server.primary_ntp.id
  ]
  syslog_servers = [
    pexip_infinity_syslog_server.central_logging.id
  ]
}
