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

type ProductFeeValueJsonV400 struct {
	Currency  string `json:"currency"`
	Amount    string `json:"amount"`
	Frequency string `json:"frequency"`
	Type_     string `json:"type"`
}