resource "pexip_infinity_ca_certificate" "example" {
  certificate = file("${path.module}/certificates/ca.crt")
}
