resource "pexip_infinity_management_vm" "primary" {
  name        = "Primary Management VM"
  description = "Primary Pexip Infinity Manager"
  
  # Network configuration
  address = "192.168.1.10"
  netmask = "255.255.255.0"
  gateway = "192.168.1.1"
  mtu     = 1500
  
  # Hostname configuration
  hostname = "pexip-mgr01"
  domain   = "company.com"
  
  # SSH configuration
  enable_ssh = "yes"
  
  # SNMP configuration
  snmp_mode = "disabled"
  
  # Management VM state
  initializing = false
}
