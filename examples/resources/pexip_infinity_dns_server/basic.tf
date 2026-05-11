# DNS Server resource example
resource "pexip_infinity_dns_server" "primary_dns" {
  address     = "8.8.8.8"
  description = "Primary Google DNS server"
}