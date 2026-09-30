resource "pexip_infinity_mjx_google_deployment" "example" {
  name         = "Corporate Google Workspace Integration"
  client_email = "pexip-service@project.iam.gserviceaccount.com"
  private_key  = var.google_private_key
}
