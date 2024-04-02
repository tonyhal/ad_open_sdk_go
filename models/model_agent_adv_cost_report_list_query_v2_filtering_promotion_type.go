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

// AgentAdvCostReportListQueryV2FilteringPromotionType
type AgentAdvCostReportListQueryV2FilteringPromotionType int64

// List of agent_adv_cost_report_list_query_v2_filtering_promotion_type
const (
	Enum_1_AgentAdvCostReportListQueryV2FilteringPromotionType AgentAdvCostReportListQueryV2FilteringPromotionType = 1
	Enum_2_AgentAdvCostReportListQueryV2FilteringPromotionType AgentAdvCostReportListQueryV2FilteringPromotionType = 2
)

// All allowed values of AgentAdvCostReportListQueryV2FilteringPromotionType enum
var AllowedAgentAdvCostReportListQueryV2FilteringPromotionTypeEnumValues = []AgentAdvCostReportListQueryV2FilteringPromotionType{
	1,
	2,
}

// NewAgentAdvCostReportListQueryV2FilteringPromotionTypeFromValue returns a pointer to a valid AgentAdvCostReportListQueryV2FilteringPromotionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentAdvCostReportListQueryV2FilteringPromotionTypeFromValue(v int64) (*AgentAdvCostReportListQueryV2FilteringPromotionType, error) {
	ev := AgentAdvCostReportListQueryV2FilteringPromotionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentAdvCostReportListQueryV2FilteringPromotionType: valid values are %v", v, AllowedAgentAdvCostReportListQueryV2FilteringPromotionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentAdvCostReportListQueryV2FilteringPromotionType) IsValid() bool {
	for _, existing := range AllowedAgentAdvCostReportListQueryV2FilteringPromotionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_adv_cost_report_list_query_v2_filtering_promotion_type value
func (v AgentAdvCostReportListQueryV2FilteringPromotionType) Ptr() *AgentAdvCostReportListQueryV2FilteringPromotionType {
	return &v
}
