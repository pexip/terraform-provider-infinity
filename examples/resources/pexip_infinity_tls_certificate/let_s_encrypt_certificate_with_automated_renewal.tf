# Using external certificate generation (example with ACME provider)
resource "acme_certificate" "letsencrypt" {
  account_key_pem = acme_private_key.account.private_key_pem
  common_name     = "pexip.company.com"
  subject_alternative_names = [
    "pexip-web.company.com",
    "pexip-api.company.com"
  ]

  dns_challenge {
    provider = "route53"
  }
}

resource "pexip_infinity_tls_certificate" "letsencrypt" {
  certificate = "${acme_certificate.letsencrypt.certificate_pem}${acme_certificate.letsencrypt.issuer_pem}"
  private_key = acme_certificate.letsencrypt.private_key_pem
  parameters  = "Let's Encrypt certificate with automated renewal"
}
