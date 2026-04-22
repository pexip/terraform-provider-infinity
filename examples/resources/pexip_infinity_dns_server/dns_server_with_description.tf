resource "pexip_infinity_dns_server" "primary_dns" {
  address     = "8.8.8.8"
  description = "Primary Google DNS server"
}

resource "pexip_infinity_dns_server" "secondary_dns" {
  address     = "8.8.4.4"
  description = "Secondary Google DNS server"
}
