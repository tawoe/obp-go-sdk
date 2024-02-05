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

type FastFirehoseAccountJsonV400 struct {
	Number            string                   `json:"number"`
	AccountAttributes []FastFirehoseAttributes `json:"account_attributes"`
	AccountRoutings   []FastFirehoseRoutings   `json:"account_routings"`
	Label             string                   `json:"label"`
	Owners            []FastFirehoseOwners     `json:"owners"`
	Balance           *AmountOfMoneyJsonV121   `json:"balance"`
	ProductCode       string                   `json:"product_code"`
	BankId            string                   `json:"bank_id"`
	Id                string                   `json:"id"`
}
