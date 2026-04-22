locals {
  certificate_chain = "${file("${path.module}/certificates/server.crt")}${file("${path.module}/certificates/intermediate.crt")}${file("${path.module}/certificates/root.crt")}"
}

resource "pexip_infinity_tls_certificate" "enterprise" {
  certificate = local.certificate_chain
  private_key = file("${path.module}/certificates/server.key")
  parameters  = "Enterprise certificate with full trust chain"
}
