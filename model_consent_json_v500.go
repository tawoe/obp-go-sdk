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

type ConsentJsonV500 struct {
	ConsentId        string `json:"consent_id"`
	Jwt              string `json:"jwt"`
	Status           string `json:"status"`
	ConsentRequestId string `json:"consent_request_id,omitempty"`
}
