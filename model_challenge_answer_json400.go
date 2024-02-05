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

type ChallengeAnswerJson400 struct {
	Id                    string `json:"id"`
	Answer                string `json:"answer"`
	ReasonCode            string `json:"reason_code,omitempty"`
	AdditionalInformation string `json:"additional_information,omitempty"`
}
