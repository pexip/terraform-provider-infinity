resource "infinity_ldap_role" "example" {
  name          = "LDAP Admin Role"
  ldap_group_dn = "cn=pexip-admins,ou=groups,dc=example,dc=com"
  role          = infinity_role.example.id
}
