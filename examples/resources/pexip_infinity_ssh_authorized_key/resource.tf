resource "pexip_infinity_ssh_authorized_key" "example" {
  keytype = "ssh-rsa"
  key     = "AAAAB3NzaC1yc2EAAAADAQABAAABgQC7vbqajDhA..."
  comment = "admin@example.com"
}
