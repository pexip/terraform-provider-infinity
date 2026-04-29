resource "pexip_infinity_policy_server" "local_policy" {
  name        = "Local Policy Server"
  description = "Local policy example using template files"
  enable_internal_service_policy       = true
  enable_internal_media_location_policy = true
  service_configuration_template = file("${path.module}/templates/service-policy.j2")
  media_location_configuration_template = file("${path.module}/templates/media-location-policy.j2")
}
