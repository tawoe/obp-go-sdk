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

type CreateAccountRequestJsonV310 struct {
	BranchId        string                   `json:"branch_id"`
	AccountRoutings []AccountRoutingJsonV121 `json:"account_routings"`
	Label           string                   `json:"label"`
	Balance         *AmountOfMoneyJsonV121   `json:"balance"`
	UserId          string                   `json:"user_id"`
	ProductCode     string                   `json:"product_code"`
}
