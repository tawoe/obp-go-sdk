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

type CustomerOverviewFlatJsonV500 struct {
	CustomerId         string                              `json:"customer_id"`
	NameSuffix         string                              `json:"name_suffix"`
	Email              string                              `json:"email"`
	BranchId           string                              `json:"branch_id"`
	MobilePhoneNumber  string                              `json:"mobile_phone_number"`
	CustomerNumber     string                              `json:"customer_number"`
	CustomerAttributes []CustomerAttributeResponseJsonV300 `json:"customer_attributes"`
	BankId             string                              `json:"bank_id"`
	DateOfBirth        string                              `json:"date_of_birth"`
	LegalName          string                              `json:"legal_name"`
	Title              string                              `json:"title"`
	Accounts           []AccountResponseJson500            `json:"accounts"`
}
