resource "pexip_infinity_autobackup" "example" {
  autobackup_enabled    = true
  autobackup_passphrase = var.autobackup_passphrase
  autobackup_interval   = 24
  autobackup_start_hour = 1
}