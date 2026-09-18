resource "pexip_infinity_mjx_exchange_deployment" "test" {
  name                     = "tf-test mjx-exchange-deployment min"
  service_account_username = "exchange-service@example.com"
  service_account_password = "test-password"
}
