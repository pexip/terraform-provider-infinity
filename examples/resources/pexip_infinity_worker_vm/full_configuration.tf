resource "pexip_infinity_worker_vm" "worker" {
  name                  = "worker-vm-01"
  hostname              = "worker-vm-01"
  domain                = "company.com" 
  address               = "10.0.1.20"
  netmask               = "255.255.255.0"
  gateway               = "10.0.1.1"
  system_location       = "Main Location"
  
  # VM specifications
  vm_cpu_count          = 8
  vm_system_memory      = 8192
  
  # Node configuration
  node_type             = "conferencing"
  transcoding           = true
  maintenance_mode      = false
  maintenance_mode_reason = ""
  
  # Network configuration
  ipv6_address          = "2001:db8::20"
  ipv6_gateway          = "2001:db8::1"
  secondary_address     = "10.0.2.20"
  secondary_netmask     = "255.255.255.0"
  static_nat_address    = "203.0.113.20"
  
  # Security
  password              = var.vm_password
  enable_ssh            = "on"
  
  # SNMP configuration
  snmp_mode             = "standard"
  snmp_community        = "private"
  snmp_system_contact   = "admin@company.com"
  snmp_system_location  = "Data Center 1"
}
