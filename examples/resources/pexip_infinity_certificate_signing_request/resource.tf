resource "pexip_infinity_certificate_signing_request" "example" {
  subject_name     = "pexip.example.com"
  private_key_type = "RSA2048"
  ad_compatible    = false
}
