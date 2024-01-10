/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AgentAdvCostReportListQueryV2DataListAppName
type AgentAdvCostReportListQueryV2DataListAppName int64

// List of agent_adv_cost_report_list_query_v2_data_list_app_name
const (
	Enum_1_AgentAdvCostReportListQueryV2DataListAppName  AgentAdvCostReportListQueryV2DataListAppName = 1
	Enum_10_AgentAdvCostReportListQueryV2DataListAppName AgentAdvCostReportListQueryV2DataListAppName = 10
	Enum_11_AgentAdvCostReportListQueryV2DataListAppName AgentAdvCostReportListQueryV2DataListAppName = 11
	Enum_12_AgentAdvCostReportListQueryV2DataListAppName AgentAdvCostReportListQueryV2DataListAppName = 12
	Enum_13_AgentAdvCostReportListQueryV2DataListAppName AgentAdvCostReportListQueryV2DataListAppName = 13
	Enum_14_AgentAdvCostReportListQueryV2DataListAppName AgentAdvCostReportListQueryV2DataListAppName = 14
	Enum_15_AgentAdvCostReportListQueryV2DataListAppName AgentAdvCostReportListQueryV2DataListAppName = 15
	Enum_16_AgentAdvCostReportListQueryV2DataListAppName AgentAdvCostReportListQueryV2DataListAppName = 16
	Enum_17_AgentAdvCostReportListQueryV2DataListAppName AgentAdvCostReportListQueryV2DataListAppName = 17
	Enum_2_AgentAdvCostReportListQueryV2DataListAppName  AgentAdvCostReportListQueryV2DataListAppName = 2
	Enum_3_AgentAdvCostReportListQueryV2DataListAppName  AgentAdvCostReportListQueryV2DataListAppName = 3
	Enum_4_AgentAdvCostReportListQueryV2DataListAppName  AgentAdvCostReportListQueryV2DataListAppName = 4
	Enum_5_AgentAdvCostReportListQueryV2DataListAppName  AgentAdvCostReportListQueryV2DataListAppName = 5
	Enum_6_AgentAdvCostReportListQueryV2DataListAppName  AgentAdvCostReportListQueryV2DataListAppName = 6
	Enum_7_AgentAdvCostReportListQueryV2DataListAppName  AgentAdvCostReportListQueryV2DataListAppName = 7
	Enum_8_AgentAdvCostReportListQueryV2DataListAppName  AgentAdvCostReportListQueryV2DataListAppName = 8
	Enum_9_AgentAdvCostReportListQueryV2DataListAppName  AgentAdvCostReportListQueryV2DataListAppName = 9
)

// All allowed values of AgentAdvCostReportListQueryV2DataListAppName enum
var AllowedAgentAdvCostReportListQueryV2DataListAppNameEnumValues = []AgentAdvCostReportListQueryV2DataListAppName{
	1,
	10,
	11,
	12,
	13,
	14,
	15,
	16,
	17,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
	9,
}

// NewAgentAdvCostReportListQueryV2DataListAppNameFromValue returns a pointer to a valid AgentAdvCostReportListQueryV2DataListAppName
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentAdvCostReportListQueryV2DataListAppNameFromValue(v int64) (*AgentAdvCostReportListQueryV2DataListAppName, error) {
	ev := AgentAdvCostReportListQueryV2DataListAppName(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentAdvCostReportListQueryV2DataListAppName: valid values are %v", v, AllowedAgentAdvCostReportListQueryV2DataListAppNameEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentAdvCostReportListQueryV2DataListAppName) IsValid() bool {
	for _, existing := range AllowedAgentAdvCostReportListQueryV2DataListAppNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_adv_cost_report_list_query_v2_data_list_app_name value
func (v AgentAdvCostReportListQueryV2DataListAppName) Ptr() *AgentAdvCostReportListQueryV2DataListAppName {
	return &v
}
