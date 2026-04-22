resource "pexip_infinity_policy_server" "lookup_policy" {
  name                        = "Directory Lookup Policy"
  description                 = "External directory and service lookup"
  url                         = "https://directory.company.com/policy"
  username                    = "directory-user"
  password                    = var.directory_password
  
  # Enable lookup services
  enable_service_lookup       = true
  enable_participant_lookup   = true
  enable_directory_lookup     = true
  enable_avatar_lookup        = true
}
