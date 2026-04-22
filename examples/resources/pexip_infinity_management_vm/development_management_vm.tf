resource "pexip_infinity_management_vm" "development" {
  name        = "Development Management VM"
  description = "Development environment Pexip Infinity Manager"
  
  address = "192.168.100.10"
  netmask = "255.255.255.0"
  gateway = "192.168.100.1"
  mtu     = 1500
  
  hostname = "pexip-dev"
  domain   = "dev.company.local"
  
  # Relaxed SSH for development
  enable_ssh = "yes"
  
  # Simple SNMP for development
  snmp_mode      = "v1v2c"
  snmp_community = "public"
  
  # Minimal services for development
  dns_servers = [
    data.pexip_infinity_dns_server.dev.id
  ]
  
  ntp_servers = [
    data.pexip_infinity_ntp_server.dev.id
  ]
  
  initializing = false
}
