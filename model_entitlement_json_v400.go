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

type EntitlementJsonV400 struct {
	EntitlementId string `json:"entitlement_id"`
	RoleName      string `json:"role_name"`
	BankId        string `json:"bank_id"`
	UserId        string `json:"user_id"`
}
