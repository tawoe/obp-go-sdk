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

type CoreTransactionJsonV300 struct {
	ThisAccount           *ThisAccountJsonV300               `json:"this_account"`
	Id                    string                             `json:"id"`
	Details               *CoreTransactionDetailsJson        `json:"details"`
	OtherAccount          *CoreCounterpartyJsonV300          `json:"other_account"`
	TransactionAttributes []TransactionAttributeResponseJson `json:"transaction_attributes"`
}