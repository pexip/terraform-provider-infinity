resource "pexip_infinity_identity_provider" "saml_example" {
  name                              = "Corporate SAML IdP"
  description                       = "Corporate SAML identity provider"
  idp_type                          = "saml"
  signature_algorithm               = "rsa-sha256"
  digest_algorithm                  = "sha256"
  assertion_consumer_service_url    = "https://pexip.example.com/saml/acs"
  sso_url                           = "https://idp.example.com/sso"
  idp_entity_id                     = "https://idp.example.com/entity"
  service_entity_id                 = "https://pexip.example.com/entity"
  display_name_attribute_name       = "displayName"
  registration_alias_attribute_name = "mail"
  worker_fqdn_acs_urls              = false
}
