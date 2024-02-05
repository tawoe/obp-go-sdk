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

type CustomerWithAttributesJsonV310 struct {
	CustomerId               string                              `json:"customer_id"`
	NameSuffix               string                              `json:"name_suffix"`
	Email                    string                              `json:"email"`
	BranchId                 string                              `json:"branch_id"`
	MobilePhoneNumber        string                              `json:"mobile_phone_number"`
	CustomerNumber           string                              `json:"customer_number"`
	CustomerAttributes       []CustomerAttributeResponseJsonV300 `json:"customer_attributes"`
	HighestEducationAttained string                              `json:"highest_education_attained"`
	DobOfDependants          []string                            `json:"dob_of_dependants"`
	BankId                   string                              `json:"bank_id"`
	DateOfBirth              string                              `json:"date_of_birth"`
	CreditRating             *CustomerCreditRatingJson           `json:"credit_rating,omitempty"`
	LastOkDate               string                              `json:"last_ok_date"`
	EmploymentStatus         string                              `json:"employment_status"`
	LegalName                string                              `json:"legal_name"`
	CreditLimit              *AmountOfMoneyJsonV121              `json:"credit_limit,omitempty"`
	Title                    string                              `json:"title"`
	FaceImage                *CustomerFaceImageJson              `json:"face_image"`
	Dependants               int32                               `json:"dependants"`
	RelationshipStatus       string                              `json:"relationship_status"`
	KycStatus                bool                                `json:"kyc_status"`
}
