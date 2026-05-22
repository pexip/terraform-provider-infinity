---
page_title: "Pexip Infinity Provider"
subcategory: ""
description: |-
  The Pexip Infinity provider enables you to manage Infinity infrastructure using Infrastructure as Code.
---

~> **Beta Notice:** This Terraform provider is currently in Beta and is not recommended for production deployments. Please use with caution.

# Pexip Infinity Provider

The Pexip Infinity Terraform provider enables you to manage [Pexip Infinity](https://www.pexip.com/products/infinity/) infrastructure using Infrastructure as Code. Automate the provisioning and management of Pexip Infinity components including the management/conferencing nodes and all other related resources.

## Features

- **Infinity node deployment**: Fully automate the deployment of the management node and conferencing nodes
- **Infinity configuration resources**: Create and manage any Infinity configuration resource
- **Infrastructure as Code**: Version control your Pexip infrastructure
- **Terraform Integration**: Native Terraform resource lifecycle management

## Requirements

- [Terraform](https://developer.hashicorp.com/terraform/downloads) >= 1.14
- Pexip Infinity Manager with API access

### Supported cloud and virtualization platforms
* GCP
* Azure
* Openstack

## Authentication

The provider supports basic authentication using username and password credentials for the Pexip Infinity Manager API.

### Provider Configuration

```terraform
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

### Basic Configuration

```terraform
terraform {
  required_providers {
    pexip = {
      source  = "pexip/infinity"
      version = "~> 0.9.11"
    }
  }
}

# Configure the Pexip Provider
provider "pexip" {
  address  = "https://manager.example.com"
  username = var.pexip_username
  password = var.pexip_password
  insecure = true
}
```
For more details on how to use the provider, see the [getting started guide](https://registry.terraform.io/providers/pexip/infinity/latest/docs/guides/getting_started) and [Common issues/FAQs](https://registry.terraform.io/providers/pexip/infinity/latest/docs/guides/common_issues_faqs).

## Support

- **Documentation**: [Pexip Documentation](https://docs.pexip.com/)
- **Questions**: [Pexip Developer Community (General)](https://community.pexip.com/c/general/16)
- **Issues**: [GitHub Issues](https://github.com/pexip/terraform-provider-pexip/issues)
- **Security**: For security concerns, please email security@pexip.com
