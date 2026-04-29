---
page_title: "Common Issues and FAQs"
---

## Common Issues

### Management node deployment

#### TLS Certificate Verification
When deploying the management node for the first time, `insecure` must be set to `true` for the provider. This is because the manager is initially configured with a self-signed certificate. Once a cert signed by a trusted CA has been assigned to the manager, then `insecure` can be set to `false`.

#### DNS & NTP
The DNS and NTP servers defined in the `pexip_infinity_manager_config` are added to the DNS and NTP servers in Infinity, but they are not added to the Terraform state. This means if the same DNS and NTP servers are also configured as resources in the Terraform configuration, and error will occur when the configuration is applied because the resources already exist in Infinity. 

There are two ways to address this. First, use temporary values in the `pexip_infinity_manager_config`. Then add the DNS and NTP servers that should be used in the Terraform config and assign them to the management node in the `management_vm` resource. The second option would be to import the DNS and NTP servers set in the `pexip_infinity_manager_config` into the Terraform config.

### Circular references

Some resources have circular references. Due to the nature of how Terraform constructs dependency graphs, the reference can only be one-way. In most cases the reference is configured on the resource where the attribute is a single resource and not a list. Some examples are:
* `tls_certificate` and `worker_vm`/`management_vm`
* `mjx_endpoint` and `mjx_endpoint_group`

### Singleton resources

Some Infinity resources are singletons, meaning there is only a single instance of the resource that always exists and can only be updated, not created or destroyed. When Terraform creates the resource, it is actually just updating it according to the Terraform configuration. When it is destroyed, all arguments are reset to the default value.

### Resource IDs

For all resources the argument `id` refers to the API field `resource_uri` and the argument `resource_id` refers to API field `id`. 

### Virtual Machine Recreation

If a config update forces the virtual machine of a conferencing node node to be recreated (e.g. changing the zone of a VM in GCP), the corresponding `worker_vm` in Infinity must also be recreated. This can be done multiple ways. One option is to destroying the Infinity `worker_vm` resource, updating the virtual machine resource config, then recreating the `worker_vm`. Another method is to update the virtual machine config then run `terraform apply -replace=<infinity-worker-vm-resource>`.
