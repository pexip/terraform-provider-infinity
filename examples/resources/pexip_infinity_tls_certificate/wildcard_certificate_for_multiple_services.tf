resource "pexip_infinity_tls_certificate" "wildcard" {
  certificate = file("${path.module}/certificates/wildcard.company.com.crt")
  private_key = file("${path.module}/certificates/wildcard.company.com.key")
  parameters  = "Wildcard certificate for *.company.com"
}
