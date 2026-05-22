resource "pexip_infinity_ntp_server" "primary_ntp" {
  address     = "pool.ntp.org"
  description = "Primary NTP pool server"
}

resource "pexip_infinity_ntp_server" "secondary_ntp" {
  address     = "time.cloudflare.com"
  description = "Cloudflare NTP server"
}
