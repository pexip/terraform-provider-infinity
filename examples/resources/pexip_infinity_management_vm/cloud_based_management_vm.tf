resource "pexip_infinity_management_vm" "cloud" {
  name        = "Cloud Management VM"
  description = "Pexip Infinity Manager in cloud environment"
  
  # Cloud network configuration
  address = "10.0.1.10"
  netmask = "255.255.255.0"
  gateway = "10.0.1.1"
  mtu     = 9000  # Jumbo frames for cloud
  
  hostname = "pexip-cloud-mgr"
  domain   = "cloud.company.com"
  
  # Cloud-specific SSH configuration
  enable_ssh                    = "keys_only"
  ssh_authorized_keys_use_cloud = true
  
  # Cloud proxy for external services
  http_proxy = data.pexip_infinity_http_proxy.cloud.id
  
  # Basic SNMP for cloud monitoring
  snmp_mode                     = "v1v2c"
  snmp_community               = var.cloud_snmp_community
  snmp_system_contact          = "Cloud Operations"
  snmp_system_location         = "AWS us-east-1"
  
  # Cloud DNS and NTP
  dns_servers = [
    data.pexip_infinity_dns_server.cloud_primary.id,
    data.pexip_infinity_dns_server.cloud_secondary.id
  ]
  
  ntp_servers = [
    data.pexip_infinity_ntp_server.cloud.id
  ]
  
  # Cloud logging
  syslog_servers = [
    data.pexip_infinity_syslog_server.cloudwatch.id
  ]
  
  # Cloud event collection
  event_sinks = [
    data.pexip_infinity_event_sink.cloud_analytics.id
  ]
  
  # Cloud certificate
  tls_certificate = data.pexip_infinity_tls_certificate.cloud.id
  
  initializing = false
}
