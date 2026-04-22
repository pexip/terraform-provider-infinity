resource "pexip_infinity_worker_vm" "worker" {
  name            = "worker-vm-01"
  hostname        = "worker-vm-01"
  domain          = "company.com"
  address         = "10.0.1.20"
  netmask         = "255.255.255.0"
  gateway         = "10.0.1.1"
  system_location = "Main Location"
}
