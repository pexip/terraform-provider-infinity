resource "pexip_infinity_ldap_role" "example" {
  name          = "LDAP Admin Role"
  ldap_group_dn = "cn=pexip-admins,ou=groups,dc=example,dc=com"
  role          = pexip_infinity_role.example.id
}
