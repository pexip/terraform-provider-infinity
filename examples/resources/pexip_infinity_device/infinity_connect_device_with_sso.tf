resource "pexip_infinity_device" "infinity_connect_device" {
  alias                       = "connect-device-01"
  description                 = "Infinity Connect device with SSO"
  primary_owner_email_address = "user@company.com"
  enable_infinity_connect_sso = true
  enable_standard_sso         = true
  sso_identity_provider_group = "corporate-users"
  tag                         = "infinity-connect"
}
