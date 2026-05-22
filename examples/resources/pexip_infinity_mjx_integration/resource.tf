resource "pexip_infinity_mjx_integration" "example" {
  name                          = "Corporate MJX Integration"
  description                   = "MJX integration for corporate calendar systems"
  display_upcoming_meetings     = 5
  enable_non_video_meetings     = true
  enable_private_meetings       = false
  end_buffer                    = 5
  start_buffer                  = 10
  ep_username                   = "pexip-service@example.com"
  ep_password                   = "secure-password"
  ep_use_https                  = true
  ep_verify_certificate         = true
  exchange_deployment           = "/configuration/v1/ms_exchange_connector/1/"
  process_alias_private_meetings = false
  replace_empty_subject         = true
  replace_subject_type          = "template"
  replace_subject_template      = "Meeting in {{ "{{" }}location{{ "}}" }}"
}
