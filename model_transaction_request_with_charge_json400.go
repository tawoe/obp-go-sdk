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

type TransactionRequestWithChargeJson400 struct {
	StartDate      string                             `json:"start_date"`
	Id             string                             `json:"id"`
	EndDate        string                             `json:"end_date"`
	Status         string                             `json:"status"`
	From           *TransactionRequestAccountJsonV140 `json:"from"`
	Details        *TransactionRequestBodyAllTypes    `json:"details"`
	Charge         *TransactionRequestChargeJsonV200  `json:"charge"`
	Type_          string                             `json:"type"`
	TransactionIds []string                           `json:"transaction_ids"`
	Challenges     []ChallengeJsonV400                `json:"challenges"`
}