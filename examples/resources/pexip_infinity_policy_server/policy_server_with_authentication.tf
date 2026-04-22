resource "pexip_infinity_policy_server" "authenticated_policy" {
  name        = "Authenticated Policy Server"
  description = "Policy server with HTTP basic authentication"
  url         = "https://policy.company.com/pexip/api"
  username    = "pexip-service"
  password    = var.policy_server_password
}
