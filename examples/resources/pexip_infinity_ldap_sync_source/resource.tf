resource "pexip_infinity_ldap_sync_source" "example" {
  name                     = "Corporate AD"
  description              = "Corporate Active Directory synchronization"
  ldap_server              = "corp.example.com"
  ldap_base_dn             = "dc=corp,dc=example,dc=com"
  ldap_bind_username       = "pexip-service@corp.example.com"
  ldap_bind_password       = "secure-password"
  ldap_use_global_catalog  = true
  ldap_permit_no_tls       = false
}
