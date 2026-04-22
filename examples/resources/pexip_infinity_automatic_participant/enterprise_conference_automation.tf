variable "automatic_participants" {
  type = list(object({
    alias               = string
    description         = string
    protocol            = string
    call_type           = string
    role                = string
    keep_conference_alive = string
    routing             = string
    streaming           = bool
    remote_display_name = string
    presentation_url    = string
  }))
  default = [
    {
      alias               = "enterprise-recorder"
      description         = "Enterprise recording service"
      protocol            = "sip"
      call_type           = "video"
      role                = "guest"
      keep_conference_alive = "end_conference_when_alone"
      routing             = "auto"
      streaming           = true
      remote_display_name = "Enterprise Recorder"
      presentation_url    = ""
    },
    {
      alias               = "compliance-monitor"
      description         = "Compliance monitoring participant"
      protocol            = "webrtc"
      call_type           = "audio"
      role                = "guest"
      keep_conference_alive = "end_conference_when_alone"
      routing             = "auto"
      streaming           = false
      remote_display_name = "Compliance Monitor"
      presentation_url    = ""
    }
  ]
}

resource "pexip_infinity_automatic_participant" "enterprise" {
  count               = length(var.automatic_participants)
  alias               = var.automatic_participants[count.index].alias
  description         = var.automatic_participants[count.index].description
  conference          = data.pexip_infinity_conference.enterprise_meeting.id
  protocol            = var.automatic_participants[count.index].protocol
  call_type           = var.automatic_participants[count.index].call_type
  role                = var.automatic_participants[count.index].role
  keep_conference_alive = var.automatic_participants[count.index].keep_conference_alive
  routing             = var.automatic_participants[count.index].routing
  streaming           = var.automatic_participants[count.index].streaming
  remote_display_name = var.automatic_participants[count.index].remote_display_name
  presentation_url    = var.automatic_participants[count.index].presentation_url != "" ? var.automatic_participants[count.index].presentation_url : null
}
