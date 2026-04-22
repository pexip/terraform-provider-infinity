# Stratum 1 NTP servers (high accuracy)
resource "pexip_infinity_ntp_server" "stratum1_ntp" {
  for_each = toset([
    "time.nist.gov",
    "time-a.nist.gov",
    "time-b.nist.gov"
  ])
  address     = each.value
  description = "NIST Stratum 1 NTP server - ${each.value}"
}
