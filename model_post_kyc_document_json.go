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

type PostKycDocumentJson struct {
	Number         string `json:"number"`
	CustomerNumber string `json:"customer_number"`
	IssueDate      string `json:"issue_date"`
	Type_          string `json:"type"`
	IssuePlace     string `json:"issue_place"`
	ExpiryDate     string `json:"expiry_date"`
}