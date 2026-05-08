resource "pexip_infinity_conference" "secure_meeting" {
  name         = "board-meeting"
  description  = "Secure board meeting room"
  service_type = "conference"

  # Security configuration
  pin              = var.meeting_pin
  guest_pin        = var.guest_pin
  allow_guests     = true
  guests_muted     = false
  hosts_can_unmute = true

  # Media configuration
  max_pixels_per_second = 1920000

  # Tracking
  tag = "executive-meetings"
}
