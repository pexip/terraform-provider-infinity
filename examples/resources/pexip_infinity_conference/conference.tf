resource "pexip_infinity_conference" "vmr" {
  name         = "Example VMR"
  service_type = "conference"
  aliases = [
    { alias = "examplevmr" },
    { alias = "examplevmr@example.com" },
    { alias = "1776" }
  ]
}
