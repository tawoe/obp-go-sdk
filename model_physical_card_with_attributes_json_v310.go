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

type PhysicalCardWithAttributesJsonV310 struct {
	Allows         []string               `json:"allows"`
	ExpiresDate    string                 `json:"expires_date"`
	Networks       []string               `json:"networks"`
	CustomerId     string                 `json:"customer_id"`
	IssueNumber    string                 `json:"issue_number"`
	Replacement    *ReplacementJson       `json:"replacement"`
	Enabled        bool                   `json:"enabled"`
	Collected      string                 `json:"collected"`
	CardNumber     string                 `json:"card_number"`
	Technology     string                 `json:"technology"`
	Cancelled      bool                   `json:"cancelled"`
	BankId         string                 `json:"bank_id"`
	CardId         string                 `json:"card_id"`
	PinReset       []PinResetJson         `json:"pin_reset"`
	SerialNumber   string                 `json:"serial_number"`
	Account        *AccountBasicV310      `json:"account"`
	ValidFromDate  string                 `json:"valid_from_date"`
	CardAttributes []CardAttributeCommons `json:"card_attributes"`
	NameOnCard     string                 `json:"name_on_card"`
	Posted         string                 `json:"posted"`
	CardType       string                 `json:"card_type"`
	OnHotList      bool                   `json:"on_hot_list"`
}