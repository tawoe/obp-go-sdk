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

type PostBankJson500 struct {
	BankRoutings []BankRoutingJsonV121 `json:"bank_routings,omitempty"`
	Website      string                `json:"website,omitempty"`
	FullName     string                `json:"full_name,omitempty"`
	Logo         string                `json:"logo,omitempty"`
	Id           string                `json:"id,omitempty"`
	BankCode     string                `json:"bank_code"`
}