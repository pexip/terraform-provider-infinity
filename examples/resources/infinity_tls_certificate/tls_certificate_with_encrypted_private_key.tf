resource "pexip_infinity_tls_certificate" "encrypted_cert" {
  certificate            = file("${path.module}/certificates/encrypted-server.crt")
  private_key            = file("${path.module}/certificates/encrypted-server.key")
  private_key_passphrase = var.private_key_passphrase
  parameters             = "SSL/TLS certificate for web interface"
}
