resource "pexip_infinity_conference_alias" "phone_number" {
  alias       = "8005551234"
  description = "Toll-free number for customer support meetings"
  conference  = data.pexip_infinity_conference.customer_support.id
}
