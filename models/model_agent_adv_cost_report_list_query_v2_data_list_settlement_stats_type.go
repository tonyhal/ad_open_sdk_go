/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AgentAdvCostReportListQueryV2DataListSettlementStatsType
type AgentAdvCostReportListQueryV2DataListSettlementStatsType int64

// List of agent_adv_cost_report_list_query_v2_data_list_settlement_stats_type
const (
	Enum_1_AgentAdvCostReportListQueryV2DataListSettlementStatsType AgentAdvCostReportListQueryV2DataListSettlementStatsType = 1
	Enum_2_AgentAdvCostReportListQueryV2DataListSettlementStatsType AgentAdvCostReportListQueryV2DataListSettlementStatsType = 2
	Enum_3_AgentAdvCostReportListQueryV2DataListSettlementStatsType AgentAdvCostReportListQueryV2DataListSettlementStatsType = 3
	Enum_4_AgentAdvCostReportListQueryV2DataListSettlementStatsType AgentAdvCostReportListQueryV2DataListSettlementStatsType = 4
	Enum_5_AgentAdvCostReportListQueryV2DataListSettlementStatsType AgentAdvCostReportListQueryV2DataListSettlementStatsType = 5
	Enum_6_AgentAdvCostReportListQueryV2DataListSettlementStatsType AgentAdvCostReportListQueryV2DataListSettlementStatsType = 6
	Enum_7_AgentAdvCostReportListQueryV2DataListSettlementStatsType AgentAdvCostReportListQueryV2DataListSettlementStatsType = 7
)

// All allowed values of AgentAdvCostReportListQueryV2DataListSettlementStatsType enum
var AllowedAgentAdvCostReportListQueryV2DataListSettlementStatsTypeEnumValues = []AgentAdvCostReportListQueryV2DataListSettlementStatsType{
	1,
	2,
	3,
	4,
	5,
	6,
	7,
}

// NewAgentAdvCostReportListQueryV2DataListSettlementStatsTypeFromValue returns a pointer to a valid AgentAdvCostReportListQueryV2DataListSettlementStatsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentAdvCostReportListQueryV2DataListSettlementStatsTypeFromValue(v int64) (*AgentAdvCostReportListQueryV2DataListSettlementStatsType, error) {
	ev := AgentAdvCostReportListQueryV2DataListSettlementStatsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentAdvCostReportListQueryV2DataListSettlementStatsType: valid values are %v", v, AllowedAgentAdvCostReportListQueryV2DataListSettlementStatsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentAdvCostReportListQueryV2DataListSettlementStatsType) IsValid() bool {
	for _, existing := range AllowedAgentAdvCostReportListQueryV2DataListSettlementStatsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_adv_cost_report_list_query_v2_data_list_settlement_stats_type value
func (v AgentAdvCostReportListQueryV2DataListSettlementStatsType) Ptr() *AgentAdvCostReportListQueryV2DataListSettlementStatsType {
	return &v
}