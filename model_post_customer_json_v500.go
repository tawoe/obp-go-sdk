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

type PostCustomerJsonV500 struct {
	NameSuffix               string                    `json:"name_suffix,omitempty"`
	Email                    string                    `json:"email,omitempty"`
	BranchId                 string                    `json:"branch_id,omitempty"`
	MobilePhoneNumber        string                    `json:"mobile_phone_number"`
	CustomerNumber           string                    `json:"customer_number,omitempty"`
	HighestEducationAttained string                    `json:"highest_education_attained,omitempty"`
	DobOfDependants          []string                  `json:"dob_of_dependants,omitempty"`
	DateOfBirth              string                    `json:"date_of_birth,omitempty"`
	CreditRating             *CustomerCreditRatingJson `json:"credit_rating,omitempty"`
	LastOkDate               string                    `json:"last_ok_date,omitempty"`
	EmploymentStatus         string                    `json:"employment_status,omitempty"`
	LegalName                string                    `json:"legal_name"`
	CreditLimit              *AmountOfMoneyJsonV121    `json:"credit_limit,omitempty"`
	Title                    string                    `json:"title,omitempty"`
	FaceImage                *CustomerFaceImageJson    `json:"face_image,omitempty"`
	Dependants               int32                     `json:"dependants,omitempty"`
	RelationshipStatus       string                    `json:"relationship_status,omitempty"`
	KycStatus                bool                      `json:"kyc_status,omitempty"`
}
