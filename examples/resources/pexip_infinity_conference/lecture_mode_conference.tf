resource "pexip_infinity_conference" "webinar" {
  name             = "company-webinar"
  service_type     = "lecture"
  description      = "Monthly company-wide webinar"
  pin              = "987654321"
  allow_guests     = true
  guests_muted     = true
  hosts_can_unmute = true
}
