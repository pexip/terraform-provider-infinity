/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

resource "infinity_azure_tenant" "azure_tenant-test" {
  name        = "tf-test-azure-tenant-full"
  description = "Test AzureTenant tf-test full"
  tenant_id   = "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"
}