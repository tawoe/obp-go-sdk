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

type CertificateInfoJsonV510 struct {
	SubjectDomainName string   `json:"subject_domain_name"`
	NotBefore         string   `json:"not_before"`
	RolesInfo         string   `json:"roles_info,omitempty"`
	Roles             []string `json:"roles,omitempty"`
	IssuerDomainName  string   `json:"issuer_domain_name"`
	NotAfter          string   `json:"not_after"`
}
