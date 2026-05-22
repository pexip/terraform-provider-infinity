resource "pexip_infinity_management_vm" "example" {
  name        = "Example mgr"
  dns_servers = [pexip_infinity_dns_server.example.id]
  ntp_servers = [pexip_infinity_ntp_server.example.id]
}

resource "pexip_infinity_dns_server" "example" {
  address = "8.8.8.8"
}

resource "pexip_infinity_ntp_server" "example" {
  address = "pool.ntp.org"
}