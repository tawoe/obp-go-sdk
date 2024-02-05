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

type PhysicalCardJson struct {
	Allows         []string         `json:"allows"`
	ExpiresDate    string           `json:"expires_date"`
	Networks       []string         `json:"networks"`
	IssueNumber    string           `json:"issue_number"`
	Replacement    *ReplacementJson `json:"replacement"`
	Enabled        bool             `json:"enabled"`
	Collected      string           `json:"collected"`
	Technology     string           `json:"technology"`
	Cancelled      bool             `json:"cancelled"`
	BankId         string           `json:"bank_id"`
	PinReset       []PinResetJson   `json:"pin_reset"`
	SerialNumber   string           `json:"serial_number"`
	Account        *AccountJson     `json:"account"`
	ValidFromDate  string           `json:"valid_from_date"`
	BankCardNumber string           `json:"bank_card_number"`
	NameOnCard     string           `json:"name_on_card"`
	Posted         string           `json:"posted"`
	OnHotList      bool             `json:"on_hot_list"`
}
