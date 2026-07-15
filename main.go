/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	infinityProvider "github.com/pexip/terraform-provider-infinity/internal/provider"
)

func createProvider() func() provider.Provider {
	return infinityProvider.New
}

func main() {
	err := providerserver.Serve(context.Background(), createProvider(), providerserver.ServeOpts{
		Address: "registry.terraform.io/pexip/infinity",
	})
	if err != nil {
		log.Printf("failed to serve provider: %s", err)
		os.Exit(1)
	}
}
