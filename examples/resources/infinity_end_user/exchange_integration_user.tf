resource "pexip_infinity_end_user" "exchange_user" {
  primary_email_address = "exchange.user@company.com"
  first_name            = "Exchange"
  last_name             = "User"
  display_name          = "Exchange User"
  department            = "IT"
  ms_exchange_guid      = "a1b2c3d4-e5f6-7890-abcd-123456789abc"
  sync_tag              = "exchange-sync"
}
