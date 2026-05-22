> **Beta Notice:** This Terraform provider is currently in Beta and is not recommended for production deployments. Please use with caution.

# Terraform Provider for Pexip Infinity

[![Build Status](https://github.com/pexip/terraform-provider-infinity/actions/workflows/test.yml/badge.svg)](https://github.com/pexip/terraform-provider-infinity/actions/workflows/test.yml)
[![Integration Tests](https://github.com/pexip/terraform-provider-infinity/actions/workflows/integration-test.yml/badge.svg)](https://github.com/pexip/terraform-provider-infinity/actions/workflows/integration-test.yml)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

The official Terraform provider for [Pexip Infinity](https://docs.pexip.com/) enables comprehensive Infrastructure as Code management of your Pexip video conferencing platform. Manage everything from basic node configurations to advanced integrations with Microsoft 365, Google Workspace, and external authentication systems.

## Features

- **🏗️ Complete Infrastructure Management**: 80+ resources covering 90% of Pexip Infinity API capabilities
- **🔐 Security & Authentication**: LDAP, Active Directory, OAuth 2.0, SAML, and certificate management
- **🎯 Conference Management**: Virtual meeting rooms, aliases, scheduled conferences, and participant controls
- **🌐 Network Configuration**: System locations, routing rules, proxies, and DNS management
- **📊 System Administration**: Licensing, logging, diagnostics, and scaling policies
- **🔗 Enterprise Integrations**: Microsoft 365, Google Workspace, Exchange, and telehealth platforms
- **📱 Media & Content**: IVR themes, media libraries, streaming credentials
- **⚡ Infrastructure as Code**: Full lifecycle management with Terraform's plan/apply workflow

## Requirements

- [Terraform](https://developer.hashicorp.com/terraform/downloads) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.21 (for development)
- Pexip Infinity Manager >= v38 with API access
- Valid authentication credentials for Pexip Infinity Manager

## Installation

### Terraform Registry (Recommended)

```hcl
terraform {
  required_version = ">= 1.14"
  required_providers {
    pexip = {
      source  = "pexip/infinity"
      version = "~> 0.9.11"
    }
  }
}
```

### Manual Installation

1. Download the latest release from [GitHub Releases](https://github.com/pexip/terraform-provider-infinity/releases)
2. Extract the binary to your Terraform plugins directory
3. Configure Terraform to use the local provider

## Documentation

For a quickstart guide, details on various resources, and other documentation, see https://registry.terraform.io/providers/pexip/infinity/latest/docs.


## Development

### Prerequisites

- [Go](https://golang.org/doc/install) >= 1.21
- [Terraform](https://developer.hashicorp.com/terraform/downloads) >= 1.0
- [Make](https://www.gnu.org/software/make/)
- Access to a Pexip Infinity Manager for testing

### Building from Source

```bash
git clone https://github.com/pexip/terraform-provider-infinity.git
cd terraform-provider-infinity
make build
```

### Local Development Setup

1. Create a `.terraformrc` file in your home directory:

```hcl
provider_installation {
  dev_overrides {
    "pexip/infinity" = "<home dir>/.terraform.d/plugins"
  }
  direct {}
}
```

2. Build and test locally:

```bash
make build
make test
```

### Available Make Targets

```bash
make build      # Build the provider binary
make install    # Build and install to local Terraform plugins directory  
make test       # Run unit tests
make testacc    # Run acceptance tests (requires Pexip environment)
make lint       # Run linting checks
make fmt        # Format Go code
make clean      # Clean build artifacts
make check      # Run all checks (lint + test)
```

### Testing

```bash
# Unit tests
make test

# Acceptance tests (requires real Pexip environment)
export TF_ACC=1
export PEXIP_ADDRESS="https://your-manager.example.com"
export PEXIP_USERNAME="admin" 
export PEXIP_PASSWORD="your-password"
make testacc

# Test specific resource
go test -v ./internal/provider -run TestAccInfinityConference
```

## Troubleshooting

### Common Issues

**Authentication Failures**
```bash
# Verify connectivity
curl -k -u username:password https://manager.example.com/api/admin/status/v1/system_summary/

# Check credentials
export TF_LOG=DEBUG
terraform plan
```

**TLS Certificate Issues**
```hcl
# For development/testing with self-signed certificates
provider "pexip" {
  address  = "https://manager.example.com"
  username = var.username
  password = var.password
  insecure = true  # Only for development!
}
```

**Network Connectivity**
```bash
# Test network access
telnet manager.example.com 443
nslookup manager.example.com
```

**Resource Import Issues**
```bash
# Get resource ID from Pexip Manager
curl -k -u admin:password "https://manager.example.com/api/admin/configuration/v1/conference/" | jq '.objects[] | {id, name}'

# Import with correct ID
terraform import pexip_infinity_conference.room 123
```

### Debug Logging

Enable comprehensive logging for troubleshooting:

```bash
export TF_LOG=DEBUG
export TF_LOG_PATH=./terraform-debug.log
terraform plan
```

### Performance Optimization

For large environments with many resources:

```hcl
# Use for_each instead of count for better performance
resource "pexip_infinity_conference" "rooms" {
  for_each = var.conference_rooms
  
  name        = each.key
  description = each.value.description
  # ... other configuration
}

# Parallelize independent resources
resource "pexip_infinity_system_location" "locations" {
  for_each = var.locations
  # ... configuration
}

resource "pexip_infinity_worker_vm" "workers" {
  for_each = var.worker_nodes
  
  name            = each.key
  hostname        = each.value.hostname
  domain          = each.value.domain
  address         = each.value.address
  netmask         = each.value.netmask
  gateway         = each.value.gateway
  system_location = each.value.location
  # ... other configuration
}
```

## Best Practices

### Security
- **Never hardcode credentials** - Use variables and environment variables
- **Use HTTPS** - Always connect to Pexip Manager over HTTPS
- **Limit permissions** - Create dedicated service accounts with minimal required permissions
- **Rotate credentials** - Regularly rotate API credentials

### Resource Organization
- **Use modules** - Organize related resources into reusable modules
- **Consistent naming** - Follow a consistent naming convention across resources
- **Tag resources** - Use descriptions to document resource purpose
- **State management** - Use remote state storage for team environments

### Performance
- **Batch operations** - Use `for_each` for creating multiple similar resources
- **Minimize dependencies** - Avoid unnecessary resource dependencies
- **Parallel execution** - Structure resources to allow parallel creation/updates

## Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

### Development Guidelines

- Follow [Go best practices](https://golang.org/doc/effective_go.html)
- Use the [Terraform Plugin Framework](https://developer.hashicorp.com/terraform/plugin/framework)
- Write comprehensive tests with >80% coverage
- Document all public APIs and configuration options
- Ensure backward compatibility when possible
- Follow semantic versioning for releases

## Support

### Community Support
- [Discussions](https://github.com/pexip/terraform-provider-infinity/discussions) - Community Q&A
- [GitHub Issues](https://github.com/pexip/terraform-provider-infinity/issues) - Bug reports and feature requests

### Documentation
- [Pexip Infinity Documentation](https://docs.pexip.com/)
- [Terraform Provider Development](https://developer.hashicorp.com/terraform/plugin)
- [Terraform Best Practices](https://developer.hashicorp.com/terraform/docs/configuration)

### Professional Support
For enterprise support, contact [Pexip Support](https://www.pexip.com/support/).

## Security

For security vulnerabilities, please email [security@pexip.com](mailto:security@pexip.com) instead of using the public issue tracker.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Pexip Engineering Team](https://www.pexip.com/) for developing and maintaining Pexip Infinity
- [HashiCorp](https://www.hashicorp.com/) for the Terraform Plugin Framework
- [Go Community](https://golang.org/) for the excellent programming language and ecosystem

---

**Made with ❤️ by the Pexip team**