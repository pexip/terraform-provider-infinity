/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

resource "infinity_azure_tenant" "azure-tenant-test" {
  name        = "tf-test-azure-tenant-teams-proxy-min"
  description = "Test Azure Tenant for Teams Proxy"
  tenant_id   = "44444444-4444-4444-4444-444444444444"
}

resource "infinity_teams_proxy" "teams-proxy-test" {
  name         = "tf-test-teams-proxy-min"
  address      = "teams-proxy-min.pexvclab.com"
  azure_tenant = infinity_azure_tenant.azure-tenant-test.id
}
