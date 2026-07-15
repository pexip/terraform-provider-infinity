resource "infinity_management_vm" "example" {
  name = "Example mgr"

  lifecycle {
    action_trigger {
      events = [after_create]
      actions = [
        action.infinity_delete_default_mgr_tls_certificate.remove_default,
      ]
    }
  }
}

action "infinity_delete_default_mgr_tls_certificate" "remove_default" {}