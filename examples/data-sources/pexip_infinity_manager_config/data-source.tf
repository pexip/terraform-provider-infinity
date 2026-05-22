data "pexip_infinity_manager_config" "config" {
  hostname              = "manager-01"
  domain                = "example.com"
  ip                    = "192.168.1.100"
  mask                  = "255.255.255.0"
  gw                    = "192.168.1.1"
  dns                   = "8.8.8.8"
  ntp                   = "pool.ntp.org"
  user                  = "admin"
  pass                  = var.manager_password
  admin_password        = var.admin_password
  error_reports         = false
  enable_analytics      = false
  contact_email_address = "admin@example.com"
}

# Use the rendered configuration
output "manager_config" {
  value     = data.pexip_infinity_manager_config.config.rendered
  sensitive = true
}
