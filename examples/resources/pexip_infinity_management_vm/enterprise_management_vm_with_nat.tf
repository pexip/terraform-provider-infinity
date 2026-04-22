resource "pexip_infinity_management_vm" "enterprise" {
  name        = "Enterprise Management VM"
  description = "Enterprise Pexip Infinity Manager with NAT"
  
  # Network configuration with NAT
  address           = "172.16.1.10"
  netmask           = "255.255.255.0"
  gateway           = "172.16.1.1"
  static_nat_address = "203.0.113.10"
  mtu               = 1500
  
  # Hostname configuration
  hostname         = "pexip-mgr"
  domain          = "company.com"
  alternative_fqdn = "pexip-external.company.com"
  
  # IPv6 configuration
  ipv6_address = "2001:db8:100::10"
  ipv6_gateway = "2001:db8:100::1"
  
  # SSH configuration
  enable_ssh                    = "keys_only"
  ssh_authorized_keys_use_cloud = true
  
  # HTTP proxy for external access
  http_proxy = data.pexip_infinity_http_proxy.corporate.id
  
  # Advanced SNMP configuration
  snmp_mode                      = "v3"
  snmp_username                  = "enterprise-monitor"
  snmp_authentication_password   = var.snmp_auth_password
  snmp_privacy_password         = var.snmp_privacy_password
  snmp_system_contact           = "Enterprise IT <enterprise-it@company.com>"
  snmp_system_location          = "Enterprise Data Center"
  snmp_network_management_system = data.pexip_infinity_snmp_nms.enterprise.id
  
  # Comprehensive network services
  dns_servers = [
    data.pexip_infinity_dns_server.internal_primary.id,
    data.pexip_infinity_dns_server.internal_secondary.id,
    data.pexip_infinity_dns_server.external.id
  ]
  
  ntp_servers = [
    data.pexip_infinity_ntp_server.internal.id,
    data.pexip_infinity_ntp_server.pool.id
  ]
  
  syslog_servers = [
    data.pexip_infinity_syslog_server.security.id,
    data.pexip_infinity_syslog_server.operations.id
  ]
  
  # Static routing
  static_routes = [
    data.pexip_infinity_static_route.corporate_network.id,
    data.pexip_infinity_static_route.dmz_network.id
  ]
  
  # Event collection
  event_sinks = [
    data.pexip_infinity_event_sink.siem.id,
    data.pexip_infinity_event_sink.analytics.id,
    data.pexip_infinity_event_sink.billing.id
  ]
  
  # Security certificates
  tls_certificate = data.pexip_infinity_tls_certificate.wildcard.id
  
  # Secondary configuration
  secondary_config_passphrase = var.secondary_config_passphrase
  
  initializing = false
}
