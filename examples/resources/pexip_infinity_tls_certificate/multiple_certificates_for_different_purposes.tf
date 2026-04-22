# Web interface certificate
resource "pexip_infinity_tls_certificate" "web_interface" {
  certificate = file("${path.module}/certificates/web.crt")
  private_key = file("${path.module}/certificates/web.key")
  parameters  = "Certificate for web interface (HTTPS)"
}

# API certificate
resource "pexip_infinity_tls_certificate" "api" {
  certificate = file("${path.module}/certificates/api.crt")
  private_key = file("${path.module}/certificates/api.key")
  parameters  = "Certificate for API access"
}

# SIP TLS certificate
resource "pexip_infinity_tls_certificate" "sip_tls" {
  certificate = file("${path.module}/certificates/sip.crt")
  private_key = file("${path.module}/certificates/sip.key")
  parameters  = "Certificate for SIP over TLS"
}
