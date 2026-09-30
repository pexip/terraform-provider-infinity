resource "pexip_infinity_ntp_server" "internal_ntp_primary" {
  address     = "ntp1.company.com"
  description = "Primary internal NTP server"
}

resource "pexip_infinity_ntp_server" "internal_ntp_secondary" {
  address     = "ntp2.company.com"
  description = "Secondary internal NTP server"
}

# IP address based NTP server
resource "pexip_infinity_ntp_server" "ip_based_ntp" {
  address     = "10.0.1.100"
  description = "Internal NTP server by IP"
}
