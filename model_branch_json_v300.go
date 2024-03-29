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

type BranchJsonV300 struct {
	Name               string                 `json:"name"`
	PhoneNumber        string                 `json:"phone_number"`
	Location           *LocationJsonV140      `json:"location"`
	BranchType         string                 `json:"branch_type"`
	BranchRouting      *BranchRoutingJsonV141 `json:"branch_routing"`
	DriveUp            *DriveUpJsonV330       `json:"drive_up"`
	MoreInfo           string                 `json:"more_info"`
	BankId             string                 `json:"bank_id"`
	Id                 string                 `json:"id"`
	Meta               *MetaJsonV140          `json:"meta"`
	Lobby              *LobbyJsonV330         `json:"lobby"`
	AccessibleFeatures string                 `json:"accessibleFeatures"`
	Address            *AddressJsonV300       `json:"address"`
	IsAccessible       string                 `json:"is_accessible"`
}
