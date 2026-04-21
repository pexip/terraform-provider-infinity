resource "pexip_infinity_break_in_allow_list_address" "example" {
  name    = "Corporate Office"
  address = "10.0.0.0"
  prefix  = 8
}
