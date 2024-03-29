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

type ResourceUserJson struct {
	Provider   string `json:"provider"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	ProviderId string `json:"provider_id"`
	UserId     string `json:"user_id"`
}
