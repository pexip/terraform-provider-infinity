variable "tls_certificate" {
  description = "TLS certificate in PEM format"
  type        = string
  sensitive   = true
}

variable "tls_private_key" {
  description = "TLS private key in PEM format"
  type        = string
  sensitive   = true
}

variable "private_key_passphrase" {
  description = "Passphrase for encrypted private key"
  type        = string
  sensitive   = true
  default     = ""
}

resource "pexip_infinity_tls_certificate" "from_variables" {
  certificate            = var.tls_certificate
  private_key            = var.tls_private_key
  private_key_passphrase = var.private_key_passphrase != "" ? var.private_key_passphrase : null
  parameters             = "Certificate deployed from Terraform variables"
}
