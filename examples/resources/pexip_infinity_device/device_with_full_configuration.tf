resource "pexip_infinity_device" "executive_room" {
  alias                       = "executive-suite.company.com"
  description                 = "Executive conference room with full capabilities"
  username                    = "executive_room"
  password                    = var.executive_room_password
  primary_owner_email_address = "executive-assistant@company.com"

  # Protocol support
  enable_sip                      = true
  enable_h323                     = true
  enable_infinity_connect_non_sso = true
  enable_infinity_connect_sso     = true
  enable_standard_sso             = true

  # SSO configuration
  sso_identity_provider_group = "executives"

  # Organizational tags
  tag      = "executive-rooms"
  sync_tag = "ldap-executives"
}
