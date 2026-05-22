# Main presentation stream
resource "pexip_infinity_automatic_participant" "webinar_stream" {
  alias                 = "webinar-main-stream"
  description           = "Main webinar stream to CDN"
  conference            = data.pexip_infinity_conference.webinar.id
  protocol              = "rtmp"
  call_type             = "video"
  role                  = "guest"
  keep_conference_alive = "end_conference_when_alone"
  routing               = "manual"
  system_location       = data.pexip_infinity_system_location.streaming_servers.id
  streaming             = true
  remote_display_name   = "Webinar Stream"
  presentation_url      = "rtmp://cdn.example.com/live/${var.webinar_stream_key}"
}

# Backup recording
resource "pexip_infinity_automatic_participant" "webinar_backup" {
  alias                 = "webinar-backup-recorder"
  description           = "Backup recording for webinar"
  conference            = data.pexip_infinity_conference.webinar.id
  protocol              = "sip"
  call_type             = "video"
  role                  = "guest"
  keep_conference_alive = "end_conference_when_alone"
  routing               = "auto"
  streaming             = true
  remote_display_name   = "Backup Recorder"
}

# Webinar host bot
resource "pexip_infinity_automatic_participant" "webinar_host" {
  alias                 = "webinar-host-bot"
  description           = "Automated webinar host"
  conference            = data.pexip_infinity_conference.webinar.id
  protocol              = "webrtc"
  call_type             = "audio"
  role                  = "chair"
  keep_conference_alive = "keep_conference_alive"
  routing               = "auto"
  streaming             = false
  remote_display_name   = "Webinar Host"
}
