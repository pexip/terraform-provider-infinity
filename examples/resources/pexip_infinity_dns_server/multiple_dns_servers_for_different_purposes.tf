# External DNS for internet resolution
resource "pexip_infinity_dns_server" "external_dns" {
  count       = length(var.external_dns_servers)
  address     = var.external_dns_servers[count.index]
  description = "External DNS server ${count.index + 1}"
}

# Internal DNS for corporate domains
resource "pexip_infinity_dns_server" "internal_dns" {
  count       = length(var.internal_dns_servers)
  address     = var.internal_dns_servers[count.index]
  description = "Internal DNS server ${count.index + 1}"
}
