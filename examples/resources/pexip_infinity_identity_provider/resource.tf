resource "pexip_infinity_identity_provider" "saml_example" {
  name                              = "Corporate SAML IdP"
  description                       = "Corporate SAML identity provider"
  uuid                              = "12345678-1234-1234-1234-123456789012"
  idp_type                          = "saml"
  assertion_consumer_service_url    = "https://pexip.example.com/saml/acs"
  sso_url                           = "https://idp.example.com/sso"
  idp_entity_id                     = "https://idp.example.com/entity"
  service_entity_id                 = "https://pexip.example.com/entity"
  display_name_attribute_name       = "displayName"
  registration_alias_attribute_name = "mail"
}
