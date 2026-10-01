resource "infinity_mjx_exchange_autodiscover_url" "example" {
  name                = "Corporate Autodiscover"
  url                 = "https://autodiscover.company.com/autodiscover/autodiscover.svc"
  exchange_deployment = infinity_mjx_exchange_deployment.example.id
}
