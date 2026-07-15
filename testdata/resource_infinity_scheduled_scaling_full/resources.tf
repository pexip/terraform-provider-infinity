/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

resource "infinity_azure_tenant" "azure-tenant-test" {
  name        = "tf-test azure-tenant-teams-proxy scheduled scaling"
  description = "Test Azure Tenant for Scheduled Scaling"
  tenant_id   = "44444444-4444-4444-4444-444444444445"
}

resource "infinity_teams_proxy" "teams-proxy-test" {
  name         = "tf-test teams-proxy scheduled scaling"
  address      = "teams-proxy-min.pexvclab.com"
  azure_tenant = infinity_azure_tenant.azure-tenant-test.id
}

resource "infinity_scheduled_scaling" "test" {
  policy_name         = "tf-test scheduled scaling full"
  policy_type         = "TeamsConnectorScaling"
  resource_identifier = infinity_teams_proxy.teams-proxy-test.name
  enabled             = false
  local_timezone      = "America/New_York"
  start_date          = "2024-06-15"
  time_from           = "08:30:00"
  time_to             = "18:30:00"
  instances_to_add    = 5
  minutes_in_advance  = 30
  mon                 = true
  tue                 = true
  wed                 = true
  thu                 = true
  fri                 = true
  sat                 = true
  sun                 = true
}
