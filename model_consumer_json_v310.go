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

type ConsumerJsonV310 struct {
	AppType        string            `json:"app_type"`
	Description    string            `json:"description"`
	CreatedByUser  *ResourceUserJson `json:"created_by_user"`
	Enabled        bool              `json:"enabled"`
	RedirectUrl    string            `json:"redirect_url"`
	DeveloperEmail string            `json:"developer_email"`
	ConsumerId     string            `json:"consumer_id"`
	AppName        string            `json:"app_name"`
	Created        string            `json:"created"`
}
