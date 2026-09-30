# Recording participant
resource "pexip_infinity_automatic_participant" "recorder" {
  alias                 = "recorder-${var.conference_name}"
  description           = "Recording for ${var.conference_name}"
  conference            = var.conference_id
  protocol              = "sip"
  call_type             = "video"
  role                  = "guest"
  keep_conference_alive = "end_conference_when_alone"
  routing               = "auto"
  streaming             = true
  remote_display_name   = "Recording Service"
}

# Streaming participant
resource "pexip_infinity_automatic_participant" "streamer" {
  alias                 = "streamer-${var.conference_name}"
  description           = "Streaming for ${var.conference_name}"
  conference            = var.conference_id
  protocol              = "rtmp"
  call_type             = "video"
  role                  = "guest"
  keep_conference_alive = "end_conference_when_alone"
  routing               = "auto"
  streaming             = true
  remote_display_name   = "Live Stream"
  presentation_url      = var.rtmp_stream_url
}

# Chair participant for persistent hosting
resource "pexip_infinity_automatic_participant" "host" {
  alias                 = "host-${var.conference_name}"
  description           = "Persistent host for ${var.conference_name}"
  conference            = var.conference_id
  protocol              = "webrtc"
  call_type             = "audio"
  role                  = "chair"
  keep_conference_alive = "keep_conference_alive"
  routing               = "auto"
  streaming             = false
  remote_display_name   = "Conference Host"
}
