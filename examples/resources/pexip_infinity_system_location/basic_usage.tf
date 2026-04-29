# Create DNS servers
resource "pexip_infinity_dns_server" "primary_dns" {
  address     = "8.8.8.8"
  description = "Primary DNS server"
}

resource "pexip_infinity_dns_server" "secondary_dns" {
  address     = "8.8.4.4"
  description = "Secondary DNS server"
}

# Create NTP servers
resource "pexip_infinity_ntp_server" "primary_ntp" {
  address     = "pool.ntp.org"
  description = "Primary NTP server"
}

# Create syslog server
resource "pexip_infinity_syslog_server" "central_logging" {
  address     = "syslog.example.com"
  port        = 514
  transport   = "tls"
  description = "Central syslog server"
}

# System location
resource "pexip_infinity_system_location" "main_location" {
  name            = "Main Office"
  description     = "Primary office location"
  mtu             = 1500
  dns_servers     = [
    pexip_infinity_dns_server.primary_dns.id,
    pexip_infinity_dns_server.secondary_dns.id
  ]
  ntp_servers     = [
    pexip_infinity_ntp_server.primary_ntp.id
  ]
  syslog_servers  = [
    pexip_infinity_syslog_server.central_logging.id
  ]
}
