# Google STUN server
resource "pexip_infinity_stun_server" "google_stun" {
  name        = "Google STUN"
  description = "Google public STUN server"
  address     = "stun.l.google.com"
  port        = 19302
}

# Cloudflare STUN server
resource "pexip_infinity_stun_server" "cloudflare_stun" {
  name        = "Cloudflare STUN"
  description = "Cloudflare public STUN server"
  address     = "stun.cloudflare.com"
  port        = 3478
}
