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

type TransactionType struct {
	ShortCode   string             `json:"shortCode"`
	Description string             `json:"description"`
	Id          *TransactionTypeId `json:"id"`
	Charge      *AmountOfMoney     `json:"charge"`
	BankId      *BankId            `json:"bankId"`
	Summary     string             `json:"summary"`
}
