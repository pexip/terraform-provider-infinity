resource "infinity_ntp_server" "primary_ntp" {
  address     = "pool.ntp.org"
  description = "Primary NTP server"
}
