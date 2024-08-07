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

// AgentAdvCostReportListQueryV2DataListCostSource
type AgentAdvCostReportListQueryV2DataListCostSource int64

// List of agent_adv_cost_report_list_query_v2_data_list_cost_source
const (
	Enum_1_AgentAdvCostReportListQueryV2DataListCostSource   AgentAdvCostReportListQueryV2DataListCostSource = 1
	Enum_102_AgentAdvCostReportListQueryV2DataListCostSource AgentAdvCostReportListQueryV2DataListCostSource = 102
	Enum_12_AgentAdvCostReportListQueryV2DataListCostSource  AgentAdvCostReportListQueryV2DataListCostSource = 12
	Enum_159_AgentAdvCostReportListQueryV2DataListCostSource AgentAdvCostReportListQueryV2DataListCostSource = 159
	Enum_2_AgentAdvCostReportListQueryV2DataListCostSource   AgentAdvCostReportListQueryV2DataListCostSource = 2
	Enum_3_AgentAdvCostReportListQueryV2DataListCostSource   AgentAdvCostReportListQueryV2DataListCostSource = 3
	Enum_31_AgentAdvCostReportListQueryV2DataListCostSource  AgentAdvCostReportListQueryV2DataListCostSource = 31
	Enum_4_AgentAdvCostReportListQueryV2DataListCostSource   AgentAdvCostReportListQueryV2DataListCostSource = 4
	Enum_5_AgentAdvCostReportListQueryV2DataListCostSource   AgentAdvCostReportListQueryV2DataListCostSource = 5
)

// All allowed values of AgentAdvCostReportListQueryV2DataListCostSource enum
var AllowedAgentAdvCostReportListQueryV2DataListCostSourceEnumValues = []AgentAdvCostReportListQueryV2DataListCostSource{
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

// NewAgentAdvCostReportListQueryV2DataListCostSourceFromValue returns a pointer to a valid AgentAdvCostReportListQueryV2DataListCostSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentAdvCostReportListQueryV2DataListCostSourceFromValue(v int64) (*AgentAdvCostReportListQueryV2DataListCostSource, error) {
	ev := AgentAdvCostReportListQueryV2DataListCostSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentAdvCostReportListQueryV2DataListCostSource: valid values are %v", v, AllowedAgentAdvCostReportListQueryV2DataListCostSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentAdvCostReportListQueryV2DataListCostSource) IsValid() bool {
	for _, existing := range AllowedAgentAdvCostReportListQueryV2DataListCostSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_adv_cost_report_list_query_v2_data_list_cost_source value
func (v AgentAdvCostReportListQueryV2DataListCostSource) Ptr() *AgentAdvCostReportListQueryV2DataListCostSource {
	return &v
}
