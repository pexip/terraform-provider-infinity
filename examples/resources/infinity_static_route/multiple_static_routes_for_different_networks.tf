# Route to internal corporate network
resource "pexip_infinity_static_route" "corporate_network" {
  name    = "Corporate Network"
  address = "10.0.0.0"
  prefix  = 8
  gateway = "192.168.100.1"
}

# Route to DMZ network
resource "pexip_infinity_static_route" "dmz_network" {
  name    = "DMZ Network"
  address = "172.20.0.0"
  prefix  = 16
  gateway = "192.168.100.2"
}

# Route to guest network
resource "pexip_infinity_static_route" "guest_network" {
  name    = "Guest Network"
  address = "192.168.50.0"
  prefix  = 24
  gateway = "192.168.100.3"
}
