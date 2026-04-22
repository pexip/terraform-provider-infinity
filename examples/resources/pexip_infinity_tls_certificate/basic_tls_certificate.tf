resource "pexip_infinity_tls_certificate" "example" {
  certificate = file("${path.module}/certificates/server.crt")
  private_key = file("${path.module}/certificates/server.key")
}
