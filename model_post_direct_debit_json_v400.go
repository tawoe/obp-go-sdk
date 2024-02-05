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

type PostDirectDebitJsonV400 struct {
	DateStarts     string `json:"date_starts"`
	CustomerId     string `json:"customer_id"`
	DateSigned     string `json:"date_signed,omitempty"`
	UserId         string `json:"user_id"`
	DateExpires    string `json:"date_expires,omitempty"`
	CounterpartyId string `json:"counterparty_id"`
}