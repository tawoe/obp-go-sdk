/*
 * Open Bank Project API
 *
 * An Open Source API for Banks. (c) TESOBE GmbH. 2011 - 2024. Licensed under the AGPL and commercial licences.
 *
 * API version: v5.1.0
 * Contact: contact@tesobe.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package obp_golang

type ApiInfoJson400 struct {
	LocalIdentityProvider    string           `json:"local_identity_provider"`
	ResourceDocsRequiresRole bool             `json:"resource_docs_requires_role"`
	Hostname                 string           `json:"hostname"`
	VersionStatus            string           `json:"version_status"`
	Version                  string           `json:"version"`
	HostedAt                 *HostedAt400     `json:"hosted_at"`
	Connector                string           `json:"connector"`
	EnergySource             *EnergySource400 `json:"energy_source"`
	HostedBy                 *HostedBy400     `json:"hosted_by"`
	GitCommit                string           `json:"git_commit"`
}
