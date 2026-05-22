# delete_default_mgr_tls_certificate (Action)

Deletes the default management node TLS certificate (ID 1) that is automatically created when a management node is deployed.

## Example Usage

```terraform
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
```
