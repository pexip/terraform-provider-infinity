resource "pexip_infinity_gms_gateway_token" "example" {
  certificate = file("${path.module}/certificates/gms.crt")
  private_key = var.gms_private_key
}
