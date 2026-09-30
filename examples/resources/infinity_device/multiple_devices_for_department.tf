# Sales team devices
resource "pexip_infinity_device" "sales_devices" {
  count = length(var.sales_team_devices)

  alias                       = var.sales_team_devices[count.index].alias
  description                 = "Sales team device ${count.index + 1}"
  username                    = var.sales_team_devices[count.index].username
  password                    = var.sales_team_devices[count.index].password
  primary_owner_email_address = var.sales_team_devices[count.index].email
  enable_sip                  = true
  enable_h323                 = false
  tag                         = "sales-team"
  sync_tag                    = "ldap-sales"
}
