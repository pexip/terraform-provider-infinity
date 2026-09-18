resource "pexip_infinity_mjx_exchange_deployment" "test" {
  name                            = "tf-test mjx-exchange-deployment full"
  description                     = "Test MJX Exchange deployment description"
  service_account_username        = "exchange-service-full@example.com"
  service_account_password        = "test-password-full"
  authentication_method           = "NTLM"
  ews_url                         = "https://exchange.example.com/EWS/Exchange.asmx"
  disable_proxy                   = true
  find_items_request_quota        = 500000
  kerberos_realm                  = "EXAMPLE.COM"
  kerberos_kdc                    = "kdc.example.com"
  kerberos_exchange_spn           = "exchangeMDB/exchange.example.com"
  kerberos_auth_every_request     = true
  kerberos_enable_tls             = false
  kerberos_kdc_https_proxy        = "https://kdc-proxy.example.com"
  kerberos_verify_tls_using_custom_ca = true
}
