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

type PostConsentRequestJsonV500 struct {
	PhoneNumber   string                           `json:"phone_number,omitempty"`
	TimeToLive    int64                            `json:"time_to_live,omitempty"`
	Email         string                           `json:"email,omitempty"`
	BankId        string                           `json:"bank_id,omitempty"`
	AccountAccess []AccountAccessV500              `json:"account_access"`
	Everything    bool                             `json:"everything"`
	ConsumerId    string                           `json:"consumer_id,omitempty"`
	ValidFrom     string                           `json:"valid_from,omitempty"`
	Entitlements  []PostConsentEntitlementJsonV310 `json:"entitlements,omitempty"`
}