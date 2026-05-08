resource "pexip_infinity_device" "boardroom_system" {
  alias                       = "boardroom.company.com"
  description                 = "Main boardroom video system"
  username                    = "boardroom"
  password                    = var.room_system_password
  primary_owner_email_address = "facilities@company.com"
  enable_sip                  = true
  enable_h323                 = true
  tag                         = "video-systems"
}
