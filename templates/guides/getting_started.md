---
page_title: "Getting Started with the Pexip Infinity Provider"
---

~> **Beta Notice:** This Terraform provider is currently in Beta and is not recommended for production deployments. Please use with caution.

# Getting Started with the Pexip Infinity Provider

[Pexip Infinity](https://docs.pexip.com/admin/admin_intro.htm) is a self-hosted, virtualized and distributed multipoint video conferencing platform. The Pexip provider enables you to manage Pexip Infinity infrastructure using Infrastructure as Code.

## Initial deployment and authentication

When performing the initial deployment of the mangement node, the Infinity provider must be configured to authenticate with the local admin account and `insecure` must be set to `true`. Once the manager has been deployed and configured to use LDAP and assigned a TLS certificate, then these settings can be updated if desired.

### Provider Configuration

```terraform
terraform {
  required_providers {
    pexip = {
      source  = "pexip/infinity"
      version = "~> 0.9.11"
    }
  }
}

provider "pexip" {
  address  = "https://manager.example.com"
  username = "admin"
  password = "secure_password"
  insecure = true
}
```

### Environment Variables

You can also configure the provider using environment variables:

```bash
export PEXIP_ADDRESS="https://manager.example.com"
export PEXIP_USERNAME="admin"
export PEXIP_PASSWORD="secure_password"
export PEXIP_INSECURE="true"
```

## Example Usage

### Management node

Here is a configuration example that deploys a managment node into GCP. There are several resources that must be configured.

#### SSH and web passwords
The deployment of the management node requires that both the `admin_password` and the `pass` need to be hashed a particular way. The `admin_passowrd` is used when accessing the cli of the management node either via the console or SSH. The `pass` is the web admin account password. Both passwords have a resource in the provider that hashes the passwords in the correct way. `pexip_infinity_ssh_password_hash` is for `admin_password` and `pexip_infinity_web_password_hash` is for the `pass`.

#### Manager config
To deploy the management node, the data resource `pexip_infinity_manager_config` must be created. This data holds the settings that would normally be entered in the installation wizard that runs when the management node is booted up for the first time. The manager config is then used by the VM that is deployed into the cloud provider.

#### GCP VM
Next is the actual VM. This is created as a typical `google_compute_instance` resource. Some items to note: 
* `metadata` refers to the `pexip_infinity_manager_config`.
* `image` refers to the GCP custom disk image. [For details on creating this image see here](https://docs.pexip.com/admin/gcp_disk_images.htm).
* The `lifecycle` block ignores changes to `metadata["management_node_config"]` and `bootdisk[0]`. `metadata["management_node_config"]` is only used when the VM is created and should be removed after deployment (more on this below). `boot_disk[0].initialize_params[0].image` is ignored so that the Infinity image can be updated in the configuration without affecting the manager.
* Network tags have been applied to allow HTTPS access to the manager and allow IPsec between other nodes.

#### Null resources
There are two null resources. `wait_for_infinity_manager_http` is used to determine when the management node is up and ready to accept management requests. Infinity resources must have an explicit dependency on this resource, otherwise Terraform will attempt to create them before the management node is fully up and running. `remove_metadata_key` automates the removal of the `management_node_config` from the VM metadata.

#### Infinity license resource
Finally there is a `pexip_infinity_licence` resource. Note the `depends_on` statement that references the `wait_for_infinity_manager_http` null resource. All Infinity resources must have the same explicit dependency, otherwise an error may occur during deployment.

```terraform
resource "pexip_infinity_ssh_password_hash" "default" {
  password = var.admin_password
}

resource "pexip_infinity_web_password_hash" "default" {
  password = var.password
}

data "pexip_infinity_manager_config" "config" {
  hostname              = var.mgr_hostname
  domain                = var.domain
  ip                    = var.mgr_prv_ip_address
  mask                  = "255.255.255.255" # GCP requrires /32
  gw                    = var.default_gateway
  dns                   = var.dns_server
  ntp                   = var.ntp_server
  user                  = var.username
  pass                  = pexip_infinity_web_password_hash.default.hash
  admin_password        = pexip_infinity_ssh_password_hash.default.hash
  error_reports         = var.report_errors
  enable_analytics      = var.enable_analytics
  contact_email_address = var.contact_email_address
}

resource "google_compute_instance" "infinity_manager" {
  name                      = var.mgr_hostname
  zone                      = var.zone
  machine_type              = var.manager_machine_type
  min_cpu_platform          = var.cpu_platform
  allow_stopping_for_update = var.allow_stopping_for_update
  deletion_protection       = var.deletion_protection
  labels                    = var.labels

  metadata = {
    management_node_config = data.pexip_infinity_manager_config.config.management_node_config
  }

  boot_disk {
    initialize_params {
      image = var.mgr_vm_image
      type  = var.disk_type
    }
  }

  tags = var.tags

  network_interface {
    network    = var.network_id
    subnetwork = var.subnet_id
    network_ip = var.mgr_prv_ip_address

    access_config {
      nat_ip = var.mgr_pub_ip_address
    }
  }

  lifecycle {
    ignore_changes = [
      metadata["management_node_config"],
      boot_disk[0].initialize_params[0].image
    ]
  }
}

resource "null_resource" "wait_for_infinity_manager_http" {
  depends_on = [google_compute_instance.infinity_manager]

  # Re‑run this null_resource whenever the instance is replaced
  triggers = {
    instance_id = google_compute_instance.infinity_manager.id
  }

  provisioner "local-exec" {
    command     = <<EOT
      echo "Waiting for Infinity Manager (HTTP 200 expected) ..."
      for i in $(seq 1 60); do
        status=$(curl --silent --show-error --insecure --location --output /dev/null --write-out "%%{http_code}" https://${var.mgr_pub_ip_address}/admin/login/)

        if [ "$status" -eq 200 ]; then
          sleep 10 # Wait for the service to become available
          echo "Infinity Manager is ready (HTTP 200)."
          exit 0
        fi

        sleep 10
      done

      echo "Timed out: Infinity Manager did not return HTTP 200" >&2
      exit 1
    EOT
    interpreter = ["/bin/bash", "-c"]
  }
}

resource "null_resource" "remove_metadata_key" {
  depends_on = [null_resource.wait_for_infinity_manager_http]

  provisioner "local-exec" {
    command = <<EOT
      gcloud compute instances remove-metadata ${google_compute_instance.infinity_manager.name} \
        --project ${var.project_id} \
        --zone ${google_compute_instance.infinity_manager.zone} \
        --keys management_node_config
    EOT
  }
}

resource "pexip_infinity_licence" "license" {
  entitlement_id = var.license_key

  depends_on = [
    null_resource.wait_for_infinity_manager_http
  ]
}
```

### Conferencing node

Here is a configuration example that deploys a conferencing node into GCP. As with the manager, there are several resources that must be configured.

#### Infinity DNS, NTP, and System location

A conferencing node must be assigned a system location. In this example a system location called `dc1` has been created. Additionally, a DNS server called `gcp_dns` and a NTP server `gcp_ntp` have also been created and assigned to the system location.

#### Worker VM
The `worker_vm` resource is the conferencing node in Infinity. Notice that the argument `node_type` is used to define whether the worker is a `CONFERENCING` or `PROXYING` node. To configure the worker to be a cloud bursting node, set `enable_distributed_database` to `false`.

#### GCP VM

The VM is also deployed using as GCP VM. Some items to note:
* When the `worker_vm` resource is created in Infinity, an XML configuration file is returned by the API. This file is used in the VM metadata. A separate resource is not required as with the manager.
* `image` refers to the GCP custom disk image. [For details on creating this image see here](https://docs.pexip.com/admin/gcp_disk_images.htm).
* The `lifecycle` block ignores changes to `metadata["conferencing_node_config"]` and `bootdisk[0]`. `metadata["conferencing_node_config"]` is only used when the VM is created and should be removed after deployment (more on this below). `boot_disk[0].initialize_params[0].image` is ignored so that the Infinity image can be updated in the configuration without affecting the manager.
* Network tags have been applied to allow HTTPS access to the conferencing node and allow IPsec between other nodes.

#### Null resources
There are two null resources that are used to remove the `metadata["conferencing_node_config"]`. Before removing the metadata, the conferencing node must sync with the manager. `wait_for_infinity_node_http` is used to determine when the conferencing node has synced with the management node. Once the node has synced, `remove_metadata_key_cnf` automates the removal of the `conferencing_node_config` from the VM metadata.

```
resource "pexip_infinity_dns_server" "gcp_dns" {
  address    = "169.254.169.254"
  depends_on = [null_resource.wait_for_infinity_manager_http]
}

resource "pexip_infinity_ntp_server" "gcp_ntp" {
  address    = "169.254.169.254"
  depends_on = [null_resource.wait_for_infinity_manager_http]
}

resource "pexip_infinity_system_location" "dc1" {
  name        = "DC1"
  dns_servers = [pexip_infinity_dns_server.gcp.id]
  ntp_servers = [pexip_infinity_ntp_server.gcp.id]
  mtu         = 1460 # GCP default
}

resource "pexip_infinity_worker_vm" "worker" {
  name                        = var.cnf_hostname
  hostname                    = var.cnf_hostname
  address                     = var.cnf_prv_ip_address
  netmask                     = "255.255.255.255" # GCP requires /32
  domain                      = var.domain
  gateway                     = var.default_gateway
  system_location             = pexip_infinity_system_location.dc1.id
  description                 = var.description
  password                    = var.admin_password
  node_type                   = "CONFERENCING" # Set to "PROXYING" for edge nodes
  enable_distributed_database = true           # Set to false for cloud bursting nodes
  depends_on = [null_resource.wait_for_infinity_manager_http]
}

resource "google_compute_instance" "infinity_worker" {
  name         = var.cnf_hostname
  zone         = var.zone
  machine_type = var.cnf_machine_type

  metadata = {
    conferencing_node_config = pexip_infinity_worker_vm.worker.config
  }

  boot_disk {
    initialize_params {
      image = var.cnf_vm_image
      type  = var.disk_type
    }
  }

  tags = var.tags

  network_interface {
    network    = var.network_id
    subnetwork = var.subnet_id
    network_ip = var.cnf_prv_ip_address

    access_config {
      nat_ip = var.cnf_pub_ip_address
    }
  }

  lifecycle {
    ignore_changes = [
      metadata["conferencing_node_config"],
      boot_disk.0.initialize_params.0.image
    ]
  }
}

resource "null_resource" "wait_for_infinity_node_http" {
  depends_on = [google_compute_instance.infinity_worker]

  # Re‑run this null_resource whenever the instance is replaced
  triggers = {
    instance_id = google_compute_instance.infinity_worker.id
  }

  provisioner "local-exec" {

    environment = {
      PASSWORD = var.password
    }

    command     = <<EOT
      echo "Waiting for Infinity Node to sync ..."
      for i in $(seq 1 60); do
        status=""
        attempts=0

        # Retry curl up to 3 times if it fails
        while [ $attempts -lt 3 ]; do
          status=$(curl --silent --show-error --insecure --location \
            -u ${var.username}:$PASSWORD \
            https://${var.mgr_pub_ip_address}/api/admin/status/v1/worker_vm/${pexip_infinity_worker_vm.worker.resource_id}/ | jq -r '.sync_status')

          if [ -n "$status" ]; then
            break  # success, exit retry loop
          fi

          attempts=$((attempts+1))
          echo "Curl attempt $attempts failed, retrying..."
          sleep 2
        done

        if [ -z "$status" ]; then
          echo "Failed to retrieve sync status after 3 attempts." >&2
          exit 1
        fi

        if [ "$status" = "SYNCED" ]; then
          echo "Infinity Node is synced."
          exit 0
        else
          echo "Infinity Node not synced, status: $status"
        fi

        sleep 10
      done

      echo "Timed out: unable to connect to Management Node" >&2
      exit 1
    EOT
    interpreter = ["/bin/bash", "-c"]
  }
}

resource "null_resource" "remove_metadata_key_cnf" {
  depends_on = [null_resource.wait_for_infinity_node_http]

  provisioner "local-exec" {
    command = <<EOT
      gcloud compute instances remove-metadata ${google_compute_instance.infinity_worker.name} \
        --project ${var.project_id} \
        --zone ${google_compute_instance.infinity_worker.zone} \
        --keys conferencing_node_config
    EOT
  }
}
```

## Complete Example

For a comprehensive example that includes Google Cloud Platform infrastructure deployment along with Pexip Infinity resources, see the `examples/cloud-providers/gcp/` directory in this repository. This example demonstrates:

- GCP VM instances for Pexip Infinity Manager and worker nodes
- Network configuration and firewall rules
- DNS setup and TLS certificates
- Service account and IAM permissions
- Integration between cloud infrastructure and Pexip resources