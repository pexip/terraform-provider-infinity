resource "pexip_infinity_worker_vm" "workers" {
  count           = 3
  name            = "worker-vm-${count.index + 1}"
  hostname        = "worker-vm-${count.index + 1}"
  domain          = "company.com"
  address         = "10.0.1.${20 + count.index}"
  netmask         = "255.255.255.0"
  gateway         = "10.0.1.1"
  system_location = "Main Location"
}
