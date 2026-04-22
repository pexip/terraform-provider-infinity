resource "pexip_infinity_management_vm" "example" {
  name        = "Primary Management VM"
  description = "Primary Pexip Infinity Manager"
  dns_servers = [pexip_infinity_dns_server.example.id]
  ntp_servers = [pexip_infinity_ntp_server.example.id]
}
