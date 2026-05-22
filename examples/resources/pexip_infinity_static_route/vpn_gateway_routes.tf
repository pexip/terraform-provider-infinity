# Route through Site-to-Site VPN
resource "pexip_infinity_static_route" "vpn_site_to_site" {
  name    = "Site-to-Site VPN"
  address = "172.30.0.0"
  prefix  = 16
  gateway = "192.168.1.100"
}

# Route through Client VPN concentrator
resource "pexip_infinity_static_route" "vpn_client" {
  name    = "Client VPN Network"
  address = "172.31.0.0"
  prefix  = 24
  gateway = "192.168.1.101"
}
