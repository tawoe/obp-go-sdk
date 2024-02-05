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

type MethodRoutingCommons struct {
	IsBankIdExactMatch bool                 `json:"is_bank_id_exact_match"`
	MethodName         string               `json:"method_name"`
	ConnectorName      string               `json:"connector_name"`
	MethodRoutingId    string               `json:"method_routing_id,omitempty"`
	BankIdPattern      string               `json:"bank_id_pattern,omitempty"`
	Parameters         []MethodRoutingParam `json:"parameters"`
}
