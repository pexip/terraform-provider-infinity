resource "pexip_infinity_device" "sip_phone" {
  alias                       = "sip-phone-101"
  description                 = "Reception desk SIP phone"
  username                    = "phone101"
  password                    = var.sip_phone_password
  primary_owner_email_address = "reception@company.com"
  enable_sip                  = true
  enable_h323                 = false
}
