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

type PostHistoricalTransactionJson struct {
	Description  string                                `json:"description"`
	To           *HistoricalTransactionAccountJsonV310 `json:"to"`
	Completed    string                                `json:"completed"`
	ChargePolicy string                                `json:"charge_policy"`
	From         *HistoricalTransactionAccountJsonV310 `json:"from"`
	Type_        string                                `json:"type"`
	Value        *AmountOfMoneyJsonV121                `json:"value"`
	Posted       string                                `json:"posted"`
}
