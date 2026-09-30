---
page_title: "Upgrading to v42 of the Pexip Infinity Provider"
---

# Upgrading to v42 of the Pexip Infinity Provider

Version 42 renames every resource, data source and action so that the prefix matches the provider name, following the standard Terraform naming convention. The redundant `pexip_` prefix has been removed:

| Before v42 | v42 and later |
|---|---|
| `pexip_infinity_conference` | `infinity_conference` |
| `pexip_infinity_worker_vm` | `infinity_worker_vm` |
| `data.pexip_infinity_manager_config` | `data.infinity_manager_config` |
| `action.pexip_delete_default_mgr_tls_certificate` | `action.infinity_delete_default_mgr_tls_certificate` |

The provider source address (`pexip/infinity`) and the resource schemas are unchanged.

## Step 1: Rename the provider local name

The provider local name in `required_providers` and in the `provider` block should now be `infinity`:

```terraform
terraform {
  required_providers {
    infinity = {
      source  = "pexip/infinity"
      version = "~> 42.0"
    }
  }
}

provider "infinity" {
  address  = "https://manager.example.com"
  username = var.infinity_username
  password = var.infinity_password
}
```

## Step 2: Rename resources in your configuration

On macOS or Linux, run the following command in the root of your configuration to update the type names:

```shell
grep -rl --include='*.tf' 'pexip_' . | xargs sed -i.bak \
  -e 's/pexip_infinity_/infinity_/g' \
  -e 's/pexip_delete_default_mgr_tls_certificate/infinity_delete_default_mgr_tls_certificate/g'
```

Check the result and delete the `.bak` files. Make sure the command didn't rename any of your own variables or locals that start with `pexip_`.

## Step 3: Migrate existing state

Existing resources must be moved to their new type names in state. If you skip this step, Terraform will plan to create the resources under their new type names and destroy the old ones. Choose one of the following options.

### Option 1: `moved` blocks (recommended)

-> `moved` blocks between resource types require Terraform 1.8 or later.

A `moved` block tells Terraform to move the existing state to the new type rather than destroying and recreating the resource:

```terraform
moved {
  from = pexip_infinity_conference.example
  to   = infinity_conference.example
}
```

You don't need to write these by hand. Run the following command in the root of your configuration to generate a `moved.tf` file containing a block for every resource in state, including resources in modules and `count`/`for_each` instances:

```shell
terraform state list | grep 'pexip_infinity_' | while read -r addr; do
  new_addr=$(printf '%s' "$addr" | sed 's/pexip_infinity_/infinity_/')
  printf 'moved {\n  from = %s\n  to   = %s\n}\n\n' "$addr" "$new_addr"
done > moved.tf
```

Data sources and actions don't hold state, so they don't need `moved` blocks.

Run `terraform plan` and confirm that each resource shows as moved, with no changes. Then run `terraform apply`.

Because `moved` blocks are part of your configuration, every workspace using the configuration is migrated on its next apply. You can delete `moved.tf` once every workspace has been applied.

### Option 2: Edit the state file

This option works with any Terraform version and doesn't need `moved` blocks, but it has to be done separately for each workspace. The moves won't appear in `terraform plan` for review.

!> Editing state by hand can corrupt it. Keep a backup, and make sure nobody else runs Terraform against the workspace while you migrate it.

Only the resource type names change. The provider address stays `provider["registry.terraform.io/pexip/infinity"]`.

1. Download the current state and keep a backup:

   ```shell
   terraform state pull > state.json
   cp state.json state.backup.json
   ```

2. Update the state, either with a script or by hand.

   **With a script:**

   ```shell
   sed 's/pexip_infinity_/infinity_/g' state.json | jq '.serial += 1' > state.new.json
   ```

   **By hand:** copy `state.json` to `state.new.json`, open it in a text editor and make these changes:

   - Increment the top-level `serial` value by 1, e.g. `"serial": 42` becomes `"serial": 43`. Terraform refuses to push the state if the serial isn't higher than the current one.
   - In every entry of the `resources` list, change the `type` from `pexip_infinity_<name>` to `infinity_<name>`:

     ```json
     {
       "mode": "managed",
       "type": "infinity_conference",
       "name": "example",
       "provider": "provider[\"registry.terraform.io/pexip/infinity\"]",
       ...
     }
     ```

   - Update any `dependencies` entries inside `instances` that refer to the old type names, e.g. `"pexip_infinity_system_location.example"` becomes `"infinity_system_location.example"`.

   Don't change anything else, including `lineage`, the `provider` address and the resource `attributes`. Your editor's find-and-replace for `pexip_infinity_` → `infinity_` covers both the `type` and `dependencies` changes.

3. Push the updated state:

   ```shell
   terraform state push state.new.json
   ```

Run `terraform plan` using the updated configuration from Steps 1 and 2 and confirm that it reports no changes.

~> Push the updated state only when you're switching to the updated configuration. If the old configuration runs against the updated state, Terraform will plan to destroy the renamed resources.

If something goes wrong, restore the backup with `terraform state push -force state.backup.json`.
