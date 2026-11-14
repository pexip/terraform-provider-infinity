/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

// schemaNameOverrides maps provider resource suffixes to SDK schema file names
// where they differ.
var schemaNameOverrides = map[string]string{
	"diagnostic_graph":     "diagnostic_graphs",
	"global_configuration": "global",
}

// skippedResources are provider resources that have no API schema.
var skippedResources = map[string]string{
	"ssh_password_hash": "computed locally, no API",
	"web_password_hash": "computed locally, no API",
}

// ignoredAPIFields are API fields never expected as provider attributes.
var ignoredAPIFields = map[string]bool{
	"id":           true, // exposed as resource_id
	"resource_uri": true, // exposed as id
}

// ignoredTFAttrs are provider attributes never expected in the API schema.
var ignoredTFAttrs = map[string]bool{
	"id":          true, // resource URI
	"resource_id": true, // API integer id
}

// ignoredFields lists per-resource fields (API or provider side) that differ
// intentionally, keyed by provider resource suffix.
var ignoredFields = map[string]map[string]string{
	"conference": {
		"scheduled_conferences": "managed by scheduling, intentionally not exposed",
	},
}

func isIgnored(suffix, field string) bool {
	_, ok := ignoredFields[suffix][field]
	return ok
}

func schemaNameFor(suffix string) string {
	if name, ok := schemaNameOverrides[suffix]; ok {
		return name
	}
	return suffix
}
