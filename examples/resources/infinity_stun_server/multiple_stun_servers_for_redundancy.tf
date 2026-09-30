# Primary internal STUN server
resource "pexip_infinity_stun_server" "internal_stun_primary" {
  name        = "Internal STUN Primary"
  description = "Primary internal STUN server"
  address     = "stun1.company.com"
  port        = 3478
}

# Secondary internal STUN server
resource "pexip_infinity_stun_server" "internal_stun_secondary" {
  name        = "Internal STUN Secondary"
  description = "Secondary internal STUN server for redundancy"
  address     = "stun2.company.com"
  port        = 3478
}

# External backup STUN server
resource "pexip_infinity_stun_server" "external_stun_backup" {
  name        = "External STUN Backup"
  description = "External STUN server for backup"
  address     = "stun.stunprotocol.org"
  port        = 3478
}
