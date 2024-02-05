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

type TransactionRequestBodyCardJsonV400 struct {
	Card        *CardJsonV400          `json:"card"`
	To          *CounterpartyIdJson    `json:"to"`
	Value       *AmountOfMoneyJsonV121 `json:"value"`
	Description string                 `json:"description"`
}