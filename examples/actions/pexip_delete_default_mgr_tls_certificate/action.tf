resource "pexip_infinity_management_vm" "manager" {
  name = var.hostname

  lifecycle {
    action_trigger {
      events = [ after_create ]
      actions = [ 
        action.pexip_delete_default_mgr_tls_certificate.remove_default,
      ]
    }
  }
}

action "pexip_delete_default_mgr_tls_certificate" "remove_default" {}
