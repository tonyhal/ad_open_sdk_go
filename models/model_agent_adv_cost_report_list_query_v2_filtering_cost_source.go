/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AgentAdvCostReportListQueryV2FilteringCostSource
type AgentAdvCostReportListQueryV2FilteringCostSource int64

// List of agent_adv_cost_report_list_query_v2_filtering_cost_source
const (
	Enum_1_AgentAdvCostReportListQueryV2FilteringCostSource   AgentAdvCostReportListQueryV2FilteringCostSource = 1
	Enum_102_AgentAdvCostReportListQueryV2FilteringCostSource AgentAdvCostReportListQueryV2FilteringCostSource = 102
	Enum_12_AgentAdvCostReportListQueryV2FilteringCostSource  AgentAdvCostReportListQueryV2FilteringCostSource = 12
	Enum_159_AgentAdvCostReportListQueryV2FilteringCostSource AgentAdvCostReportListQueryV2FilteringCostSource = 159
	Enum_2_AgentAdvCostReportListQueryV2FilteringCostSource   AgentAdvCostReportListQueryV2FilteringCostSource = 2
	Enum_3_AgentAdvCostReportListQueryV2FilteringCostSource   AgentAdvCostReportListQueryV2FilteringCostSource = 3
	Enum_31_AgentAdvCostReportListQueryV2FilteringCostSource  AgentAdvCostReportListQueryV2FilteringCostSource = 31
	Enum_4_AgentAdvCostReportListQueryV2FilteringCostSource   AgentAdvCostReportListQueryV2FilteringCostSource = 4
	Enum_5_AgentAdvCostReportListQueryV2FilteringCostSource   AgentAdvCostReportListQueryV2FilteringCostSource = 5
)

// All allowed values of AgentAdvCostReportListQueryV2FilteringCostSource enum
var AllowedAgentAdvCostReportListQueryV2FilteringCostSourceEnumValues = []AgentAdvCostReportListQueryV2FilteringCostSource{
	1,
	102,
	12,
	159,
	2,
	3,
	31,
	4,
	5,
}

// NewAgentAdvCostReportListQueryV2FilteringCostSourceFromValue returns a pointer to a valid AgentAdvCostReportListQueryV2FilteringCostSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentAdvCostReportListQueryV2FilteringCostSourceFromValue(v int64) (*AgentAdvCostReportListQueryV2FilteringCostSource, error) {
	ev := AgentAdvCostReportListQueryV2FilteringCostSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentAdvCostReportListQueryV2FilteringCostSource: valid values are %v", v, AllowedAgentAdvCostReportListQueryV2FilteringCostSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentAdvCostReportListQueryV2FilteringCostSource) IsValid() bool {
	for _, existing := range AllowedAgentAdvCostReportListQueryV2FilteringCostSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_adv_cost_report_list_query_v2_filtering_cost_source value
func (v AgentAdvCostReportListQueryV2FilteringCostSource) Ptr() *AgentAdvCostReportListQueryV2FilteringCostSource {
	return &v
}
