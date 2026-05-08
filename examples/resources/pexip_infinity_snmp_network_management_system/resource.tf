resource "pexip_infinity_snmp_network_management_system" "example" {
  name                = "Primary NMS"
  description         = "Primary network management system for monitoring"
  address             = "nms.example.com"
  port                = 162
  snmp_trap_community = "public"
}
