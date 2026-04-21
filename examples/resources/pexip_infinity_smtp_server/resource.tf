resource "pexip_infinity_smtp_server" "example" {
  name                = "Corporate SMTP Server"
  description         = "SMTP server for email notifications"
  address             = "smtp.example.com"
  port                = 587
  username            = "pexip@example.com"
  password            = var.smtp_password
  from_email_address  = "noreply@example.com"
  connection_security = "starttls"
}
