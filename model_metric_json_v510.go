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

type MetricJsonV510 struct {
	Duration                     int64                       `json:"duration"`
	ResponseBody                 *MetricJsonV510ResponseBody `json:"response_body"`
	ImplementedInVersion         string                      `json:"implemented_in_version"`
	TargetIp                     string                      `json:"target_ip"`
	Url                          string                      `json:"url"`
	CorrelationId                string                      `json:"correlation_id"`
	ImplementedByPartialFunction string                      `json:"implemented_by_partial_function"`
	UserId                       string                      `json:"user_id"`
	DeveloperEmail               string                      `json:"developer_email"`
	Date                         string                      `json:"date"`
	ConsumerId                   string                      `json:"consumer_id"`
	Verb                         string                      `json:"verb"`
	AppName                      string                      `json:"app_name"`
	SourceIp                     string                      `json:"source_ip"`
	UserName                     string                      `json:"user_name"`
}
