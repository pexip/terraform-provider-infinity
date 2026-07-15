/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

resource "tls_private_key" "ssh" {
  algorithm = "RSA"
  rsa_bits  = 4096
}

resource "infinity_dns_server" "tf-test-dns" {
  address = "1.0.0.1"
}

resource "infinity_ntp_server" "tf-test-ntp" {
  address = "europe.pool.ntp.org"
}

resource "infinity_syslog_server" "tf-test-syslog" {
  address = "syslog.example.com"
}

resource "infinity_ssh_authorized_key" "tf-test-ssh-key" {
  keytype = "ssh-rsa"
  key     = element(split(" ", trimspace(tls_private_key.ssh.public_key_openssh)), 1)
  comment = "tf-test-ssh-key"
}

resource "infinity_event_sink" "tf-test-event-sink" {
  name = "tf-test-event-sink"
  url  = "https://events.example.com/webhook"
}

resource "infinity_static_route" "tf-test-static-route" {
  name    = "tf-test-static-route"
  address = "10.10.0.0"
  prefix  = 16
  gateway = "192.168.1.254"
}

resource "infinity_http_proxy" "tf-test-http-proxy" {
  name    = "tf-test-http-proxy"
  address = "proxy.example.com"
}

resource "infinity_snmp_network_management_system" "tf-test-snmp-nms" {
  name    = "tf-test-snmp-nms"
  address = "192.168.1.200"
}

resource "infinity_management_vm" "management_vm-test" {
  name                           = "tf-test-management-vm"
  description                    = "Integration test management VM"
  mtu                            = 1400
  enable_ssh                     = "ON"
  ssh_authorized_keys_use_cloud  = false
  snmp_mode                      = "AUTHPRIV"
  snmp_community                 = "public"
  snmp_username                  = "tf-test-snmp-user"
  snmp_authentication_password   = "test-auth-pass123!"
  snmp_privacy_password          = "test-priv-pass123!"
  snmp_system_contact            = "admin@example.com"
  snmp_system_location           = "tf-test-datacenter"
  dns_servers                    = [infinity_dns_server.tf-test-dns.id]
  ntp_servers                    = [infinity_ntp_server.tf-test-ntp.id]
  syslog_servers                 = [infinity_syslog_server.tf-test-syslog.id]
  ssh_authorized_keys            = [infinity_ssh_authorized_key.tf-test-ssh-key.id]
  event_sinks                    = [infinity_event_sink.tf-test-event-sink.id]
  static_routes                  = [infinity_static_route.tf-test-static-route.id]
  http_proxy                     = infinity_http_proxy.tf-test-http-proxy.id
  snmp_network_management_system = infinity_snmp_network_management_system.tf-test-snmp-nms.id
}
