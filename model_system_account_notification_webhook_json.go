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

type SystemAccountNotificationWebhookJson struct {
	CreatedByUserId string `json:"created_by_user_id"`
	Url             string `json:"url"`
	TriggerName     string `json:"trigger_name"`
	HttpProtocol    string `json:"http_protocol"`
	HttpMethod      string `json:"http_method"`
	WebhookId       string `json:"webhook_id"`
}
