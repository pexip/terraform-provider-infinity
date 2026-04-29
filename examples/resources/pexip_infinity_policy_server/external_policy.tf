resource "pexip_infinity_policy_server" "external_policy" {
  name        = "External Policy Server"
  description = "Policy server with HTTP basic authentication"
  url         = "https://policy.example.com/pexip/api"
  username    = "pexip-service"
  password    = var.policy_server_password
}
