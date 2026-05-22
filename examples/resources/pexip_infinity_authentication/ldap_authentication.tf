resource "pexip_infinity_authentication" "ldap_example" {
  source                       = "LDAP+LOCAL"
  ldap_server                  = "ldaps://ldap.example.com:636"
  ldap_base_dn                 = "dc=example,dc=com"
  ldap_bind_username           = "cn=pexip,ou=service,dc=example,dc=com"
  ldap_bind_password           = var.ldap_bind_password
  ldap_user_search_dn          = "ou=users,dc=example,dc=com"
  ldap_user_filter             = "(&(objectclass=person)(!(objectclass=computer)))"
  ldap_user_search_filter      = "(|(uid={username})(sAMAccountName={username})(mail={username}))"
  ldap_user_group_attributes   = "memberOf"
  ldap_group_search_dn         = "ou=groups,dc=example,dc=com"
  ldap_group_filter            = "(objectclass=group)"
  ldap_group_membership_filter = "(member={userdn})"
}
