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

type SepaCreditTransfers struct {
	DebtorAccount    *PaymentAccount        `json:"debtorAccount"`
	InstructedAmount *AmountOfMoneyJsonV121 `json:"instructedAmount"`
	CreditorAccount  *PaymentAccount        `json:"creditorAccount"`
	CreditorName     string                 `json:"creditorName"`
}
