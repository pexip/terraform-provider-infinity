resource "pexip_infinity_mjx_exchange_deployment" "example" {
  name                     = "Corporate Exchange Integration"
  service_account_username = "pexip-service@company.com"
  service_account_password = var.exchange_password
}
