resource "pexip_infinity_scheduled_alias" "example" {
  alias              = "meeting-alias"
  alias_number       = 1001
  numeric_alias      = "8001"
  uuid               = "12345678-1234-1234-1234-123456789012"
  exchange_connector = pexip_infinity_ms_exchange_connector.example.id
  is_used            = false
}
