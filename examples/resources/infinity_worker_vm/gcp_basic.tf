resource "pexip_infinity_worker_vm" "worker" {
  name               = var.hostname
  hostname           = var.hostname
  address            = var.prv_ip_address
  netmask            = "255.255.255.255" # GCP requires a /32 netmask for the vm
  domain             = var.domain
  gateway            = var.default_gateway
  system_location    = var.system_location
  password           = var.admin_password
  static_nat_address = var.pub_ip_address
}

resource "google_compute_instance" "infinity_worker" {
  name         = var.hostname
  zone         = var.zone
  machine_type = var.machine_type

  metadata = {
    conferencing_node_config = pexip_infinity_worker_vm.worker.config
  }

  boot_disk {
    initialize_params {
      image = data.google_compute_image.pexip-infinity-node-image.self_link
      type  = "pd-ssd"
    }
  }

  tags = var.tags

  network_interface {
    network    = var.network_id
    subnetwork = var.subnet_id
    network_ip = var.prv_ip_address

    access_config {
      nat_ip = var.pub_ip_address
    }
  }
}