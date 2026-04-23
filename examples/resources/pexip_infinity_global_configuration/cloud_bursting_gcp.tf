resource "pexip_infinity_global_configuration" "gcp_bursting" {
  bursting_enabled  = true
  cloud_provider    = "GCP"
  gcp_project_id    = var.gcp_project_id
  gcp_client_email  = var.gcp_client_email
  gcp_private_key   = var.gcp_private_key
}
