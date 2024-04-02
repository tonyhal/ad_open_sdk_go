/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AgentAdvCostReportListQueryV2OrderType
type AgentAdvCostReportListQueryV2OrderType string

// List of agent_adv_cost_report_list_query_v2_order_type
const (
	DESC_AgentAdvCostReportListQueryV2OrderType AgentAdvCostReportListQueryV2OrderType = "DESC"
	ASC_AgentAdvCostReportListQueryV2OrderType  AgentAdvCostReportListQueryV2OrderType = "ASC"
)

// All allowed values of AgentAdvCostReportListQueryV2OrderType enum
var AllowedAgentAdvCostReportListQueryV2OrderTypeEnumValues = []AgentAdvCostReportListQueryV2OrderType{
	"DESC",
	"ASC",
}

// NewAgentAdvCostReportListQueryV2OrderTypeFromValue returns a pointer to a valid AgentAdvCostReportListQueryV2OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentAdvCostReportListQueryV2OrderTypeFromValue(v string) (*AgentAdvCostReportListQueryV2OrderType, error) {
	ev := AgentAdvCostReportListQueryV2OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentAdvCostReportListQueryV2OrderType: valid values are %v", v, AllowedAgentAdvCostReportListQueryV2OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentAdvCostReportListQueryV2OrderType) IsValid() bool {
	for _, existing := range AllowedAgentAdvCostReportListQueryV2OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_adv_cost_report_list_query_v2_order_type value
func (v AgentAdvCostReportListQueryV2OrderType) Ptr() *AgentAdvCostReportListQueryV2OrderType {
	return &v
}
