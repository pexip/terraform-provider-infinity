resource "pexip_infinity_automatic_participant" "rtmp_streamer" {
  alias                 = "live-stream"
  description           = "RTMP streaming to social media"
  conference            = data.pexip_infinity_conference.webinar.id
  protocol              = "rtmp"
  call_type             = "video"
  role                  = "guest"
  keep_conference_alive = "end_conference_when_alone"
  routing               = "manual"
  system_location       = data.pexip_infinity_system_location.streaming.id
  streaming             = true
  remote_display_name   = "Live Stream"
  presentation_url      = "rtmp://streaming.example.com/live/stream-key"
}
