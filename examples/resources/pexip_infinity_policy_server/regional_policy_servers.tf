# Different policy servers for different regions
resource "pexip_infinity_policy_server" "regional_policy" {
  for_each = var.regional_policy_servers
  
  name                          = "Policy Server - ${each.key}"
  description                   = "Regional policy server for ${each.key}"
  url                           = each.value.url
  username                      = each.value.username
  password                      = each.value.password
  
  enable_service_lookup         = true
  enable_participant_lookup     = true
  enable_directory_lookup       = each.value.enable_directory
  enable_avatar_lookup          = each.value.enable_avatars
}
