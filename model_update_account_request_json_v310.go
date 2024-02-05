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

type UpdateAccountRequestJsonV310 struct {
	Label           string                   `json:"label"`
	Type_           string                   `json:"type"`
	BranchId        string                   `json:"branch_id"`
	AccountRoutings []AccountRoutingJsonV121 `json:"account_routings"`
}
