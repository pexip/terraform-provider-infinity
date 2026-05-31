resource "pexip_infinity_conference" "virtual_auditorium" {
  name             = "company-webinar"
  service_type     = "lecture"
  description      = "Monthly company-wide webinar"
  pin              = "987654"
  allow_guests     = true
  mute_all_guests  = true
  aliases = [
    { alias = "exampleva" },
    { alias = "exampleva@example.com" },
  ]
}
