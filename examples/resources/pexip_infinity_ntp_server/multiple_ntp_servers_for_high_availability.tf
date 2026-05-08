# Public NTP servers
resource "pexip_infinity_ntp_server" "public_ntp" {
  count       = length(var.ntp_servers)
  address     = var.ntp_servers[count.index]
  description = "Public NTP server ${count.index + 1}"
}

# Regional NTP servers
locals {
  regional_ntp_servers = {
    "us"   = ["time1.google.com", "time2.google.com"]
    "eu"   = ["ntp1.ntp.se", "ntp2.ntp.se"]
    "asia" = ["ntp.nict.jp", "ntp1.jst.mfeed.ad.jp"]
  }
}

resource "pexip_infinity_ntp_server" "regional_ntp" {
  for_each    = toset(local.regional_ntp_servers[var.region])
  address     = each.value
  description = "Regional NTP server for ${var.region} - ${each.value}"
}
