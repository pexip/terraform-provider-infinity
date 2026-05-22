resource "pexip_infinity_global_configuration" "example" {
  logon_banner  = "Welcome to Pexip Infinity"
  enable_webrtc = true
  enable_sip    = true
  enable_h323   = false
  enable_rtmp   = true
  crypto_mode   = "besteffort"
}
