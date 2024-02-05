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

type LobbyJsonV330 struct {
	Sunday    []OpeningTimesV300 `json:"sunday"`
	Tuesday   []OpeningTimesV300 `json:"tuesday"`
	Wednesday []OpeningTimesV300 `json:"wednesday"`
	Monday    []OpeningTimesV300 `json:"monday"`
	Friday    []OpeningTimesV300 `json:"friday"`
	Thursday  []OpeningTimesV300 `json:"thursday"`
	Saturday  []OpeningTimesV300 `json:"saturday"`
}