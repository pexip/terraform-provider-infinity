resource "infinity_management_vm" "example" {
  name        = "Example mgr"
  dns_servers = [infinity_dns_server.example.id]
  ntp_servers = [infinity_ntp_server.example.id]
}

resource "infinity_dns_server" "example" {
  address = "8.8.8.8"
}

resource "infinity_ntp_server" "example" {
  address = "pool.ntp.org"
}