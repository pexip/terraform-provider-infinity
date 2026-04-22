resource "pexip_infinity_dns_server" "internal_dns_primary" {
  address     = "10.0.1.10"
  description = "Primary internal DNS server"
}

resource "pexip_infinity_dns_server" "internal_dns_secondary" {
  address     = "10.0.1.11"
  description = "Secondary internal DNS server"
}

# IPv6 DNS server
resource "pexip_infinity_dns_server" "ipv6_dns" {
  address     = "2001:4860:4860::8888"
  description = "Google IPv6 DNS server"
}
