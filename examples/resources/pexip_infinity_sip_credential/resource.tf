resource "pexip_infinity_sip_credential" "example" {
  realm    = "sip.example.com"
  username = "sip-user"
  password = var.sip_password
}
