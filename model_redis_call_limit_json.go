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

type RedisCallLimitJson struct {
	PerWeek   *RateLimit `json:"per_week,omitempty"`
	PerSecond *RateLimit `json:"per_second,omitempty"`
	PerMonth  *RateLimit `json:"per_month,omitempty"`
	PerDay    *RateLimit `json:"per_day,omitempty"`
	PerMinute *RateLimit `json:"per_minute,omitempty"`
	PerHour   *RateLimit `json:"per_hour,omitempty"`
}