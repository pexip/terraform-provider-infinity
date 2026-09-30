# Route to AWS VPC
resource "pexip_infinity_static_route" "aws_vpc" {
  name    = "AWS VPC Network"
  address = "10.100.0.0"
  prefix  = 16
  gateway = "192.168.1.50"
}

# Route to Azure VNet
resource "pexip_infinity_static_route" "azure_vnet" {
  name    = "Azure VNet Network"
  address = "10.200.0.0"
  prefix  = 16
  gateway = "192.168.1.51"
}

# Route to Google Cloud VPC
resource "pexip_infinity_static_route" "gcp_vpc" {
  name    = "Google Cloud VPC"
  address = "10.300.0.0"
  prefix  = 16
  gateway = "192.168.1.52"
}
