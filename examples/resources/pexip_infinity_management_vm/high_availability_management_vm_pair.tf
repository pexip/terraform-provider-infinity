# Primary Management VM
resource "pexip_infinity_management_vm" "primary" {
  name        = "Primary Management VM"
  description = "Primary Pexip Infinity Manager"
  
  address = "10.1.1.10"
  netmask = "255.255.255.0"
  gateway = "10.1.1.1"
  mtu     = 1500
  
  hostname = "pexip-mgr01"
  domain   = "enterprise.com"
  
  # IPv6 support
  ipv6_address = "2001:db8::10"
  ipv6_gateway = "2001:db8::1"
  
  # SSH with key-only access
  enable_ssh = "keys_only"
  ssh_authorized_keys = [
    data.pexip_infinity_ssh_key.admin_key.id
  ]
  
  # SNMP v3 configuration
  snmp_mode                     = "v3"
  snmp_username                 = "pexip-monitor"
  snmp_authentication_password  = var.snmp_auth_password
  snmp_privacy_password        = var.snmp_privacy_password
  snmp_system_contact          = "IT Operations <it-ops@enterprise.com>"
  snmp_system_location         = "Data Center 1, Rack A1"
  
  # Network services
  dns_servers = [
    data.pexip_infinity_dns_server.primary.id,
    data.pexip_infinity_dns_server.secondary.id
  ]
  
  ntp_servers = [
    data.pexip_infinity_ntp_server.primary.id,
    data.pexip_infinity_ntp_server.secondary.id
  ]
  
  # Logging and monitoring
  syslog_servers = [
    data.pexip_infinity_syslog_server.central.id
  ]
  
  event_sinks = [
    data.pexip_infinity_event_sink.analytics.id,
    data.pexip_infinity_event_sink.monitoring.id
  ]
  
  # Security
  tls_certificate = data.pexip_infinity_tls_certificate.management.id
  
  initializing = false
}

# Secondary Management VM
resource "pexip_infinity_management_vm" "secondary" {
  name        = "Secondary Management VM"
  description = "Secondary Pexip Infinity Manager for HA"
  
  address = "10.1.1.11"
  netmask = "255.255.255.0"
  gateway = "10.1.1.1"
  mtu     = 1500
  
  hostname = "pexip-mgr02"
  domain   = "enterprise.com"
  
  # IPv6 support
  ipv6_address = "2001:db8::11"
  ipv6_gateway = "2001:db8::1"
  
  # SSH with key-only access
  enable_ssh = "keys_only"
  ssh_authorized_keys = [
    data.pexip_infinity_ssh_key.admin_key.id
  ]
  
  # SNMP v3 configuration
  snmp_mode                     = "v3"
  snmp_username                 = "pexip-monitor"
  snmp_authentication_password  = var.snmp_auth_password
  snmp_privacy_password        = var.snmp_privacy_password
  snmp_system_contact          = "IT Operations <it-ops@enterprise.com>"
  snmp_system_location         = "Data Center 2, Rack B1"
  
  # Network services
  dns_servers = [
    data.pexip_infinity_dns_server.primary.id,
    data.pexip_infinity_dns_server.secondary.id
  ]
  
  ntp_servers = [
    data.pexip_infinity_ntp_server.primary.id,
    data.pexip_infinity_ntp_server.secondary.id
  ]
  
  # Logging and monitoring
  syslog_servers = [
    data.pexip_infinity_syslog_server.central.id
  ]
  
  event_sinks = [
    data.pexip_infinity_event_sink.analytics.id,
    data.pexip_infinity_event_sink.monitoring.id
  ]
  
  # Security
  tls_certificate = data.pexip_infinity_tls_certificate.management.id
  
  initializing = false
}
