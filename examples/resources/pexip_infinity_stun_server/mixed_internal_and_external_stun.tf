# Internal STUN for office networks
resource "pexip_infinity_stun_server" "internal_office_stun" {
  name        = "Office Internal STUN"
  description = "STUN server for internal office network"
  address     = "10.0.1.100"
  port        = 3478
}

# External STUN for remote users
resource "pexip_infinity_stun_server" "external_remote_stun" {
  name        = "Remote User STUN"
  description = "External STUN server for remote users"
  address     = "stun.external.company.com"
  port        = 3478
}
