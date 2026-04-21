resource "pexip_infinity_adfs_auth_server" "example" {
  name                               = "ADFS Server"
  description                        = "ADFS authentication server for corporate users"
  client_id                          = "12345678-1234-1234-1234-123456789012"
  federation_service_name            = "adfs.example.com"
  federation_service_identifier      = "http://adfs.example.com/adfs/services/trust"
  relying_party_trust_identifier_url = "https://pexip.example.com"
}
