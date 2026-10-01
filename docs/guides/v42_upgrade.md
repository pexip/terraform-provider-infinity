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

The provider local name should now be `infinity`. Terraform works out which provider a resource belongs to from the prefix of its type name, so `infinity_*` resources need a provider with the local name `infinity`.

~> Every module declares its own `required_providers`. Make this change in **every** module that uses Infinity resources, including child modules and shared modules, not just the root module. If a module still uses the local name `pexip`, `terraform init` fails with an error saying `registry.terraform.io/hashicorp/infinity` is required.

In each module, rename the entry in `required_providers`:

```terraform
terraform {
  required_providers {
    infinity = {
      source  = "pexip/infinity"
      version = "~> 42.0"
    }
  }
}
```

In the root module, also rename the `provider` block:

```terraform
provider "infinity" {
  address  = "https://manager.example.com"
  username = var.infinity_username
  password = var.infinity_password
}
```

If you pass providers to modules explicitly, or set `provider` on individual resources, update those references too:

```terraform
module "manager" {
  source = "./modules/infinity-manager"

  providers = {
    infinity = infinity
  }
}
```

## Step 2: Rename resources in your configuration

On macOS or Linux, run the following command in the root of your configuration to update the type names. The command only updates files in the current directory and its subdirectories, so also run it in any directory containing modules that live outside your configuration, e.g. `../modules`:

```shell
grep -rl --include='*.tf' 'pexip_' . | xargs sed -i.bak \
  -e 's/pexip_infinity_/infinity_/g' \
  -e 's/pexip_delete_default_mgr_tls_certificate/infinity_delete_default_mgr_tls_certificate/g'
```

Check the result and delete the `.bak` files. Make sure the command didn't rename any of your own variables or locals that start with `pexip_`.

~> The command replaces `pexip_infinity_` wherever it appears, including in the middle of your own resource names. For example, `null_resource.wait_for_pexip_infinity_manager` would become `null_resource.wait_for_infinity_manager`. Terraform treats a renamed resource as a new resource, so it would destroy the existing one and create a replacement. Compare the files against the `.bak` copies before deleting them and undo any changes to your own resource names, so that only the resource types have changed.

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
terraform state list | grep -E '(^|\.)pexip_infinity_' | grep -vE '(^|\.)data\.' | while read -r addr; do
  new_addr=$(printf '%s' "$addr" | sed -E 's/(^|\.)pexip_infinity_/\1infinity_/')
  printf 'moved {\n  from = %s\n  to   = %s\n}\n\n' "$addr" "$new_addr"
done > moved.tf
```

Data sources and actions don't hold state that can be moved, so the command skips them. Don't add `moved` blocks for data sources, or the plan fails with a `Resource Type Not Found` error.

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
   sed -E 's/([".])pexip_infinity_/\1infinity_/g' state.json | jq '.serial += 1' > state.new.json
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

   Don't change anything else, including `lineage`, the `provider` address and the resource `attributes`. If you use your editor's find-and-replace for `pexip_infinity_` → `infinity_`, check each match: only replace it where it starts a type name, not inside your own resource names such as `null_resource.wait_for_pexip_infinity_manager`.

3. Push the updated state:

   ```shell
   terraform state push state.new.json
   ```

Run `terraform plan` using the updated configuration from Steps 1 and 2 and confirm that it reports no changes.

~> Push the updated state only when you're switching to the updated configuration. If the old configuration runs against the updated state, Terraform will plan to destroy the renamed resources.

If something goes wrong, restore the backup with `terraform state push -force state.backup.json`.
