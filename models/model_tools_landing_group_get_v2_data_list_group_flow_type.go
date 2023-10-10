/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsLandingGroupGetV2DataListGroupFlowType
type ToolsLandingGroupGetV2DataListGroupFlowType string

// List of tools_landing_group_get_v2_data_list_group_flow_type
const (
	FLOW_DISTRIBUTION_TYPE_INTELLIGENCE_ToolsLandingGroupGetV2DataListGroupFlowType ToolsLandingGroupGetV2DataListGroupFlowType = "FLOW_DISTRIBUTION_TYPE_INTELLIGENCE"
	FLOW_DISTRIBUTION_TYPE_AVERAGE_ToolsLandingGroupGetV2DataListGroupFlowType      ToolsLandingGroupGetV2DataListGroupFlowType = "FLOW_DISTRIBUTION_TYPE_AVERAGE"
)

// All allowed values of ToolsLandingGroupGetV2DataListGroupFlowType enum
var AllowedToolsLandingGroupGetV2DataListGroupFlowTypeEnumValues = []ToolsLandingGroupGetV2DataListGroupFlowType{
	"FLOW_DISTRIBUTION_TYPE_INTELLIGENCE",
	"FLOW_DISTRIBUTION_TYPE_AVERAGE",
}

// NewToolsLandingGroupGetV2DataListGroupFlowTypeFromValue returns a pointer to a valid ToolsLandingGroupGetV2DataListGroupFlowType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsLandingGroupGetV2DataListGroupFlowTypeFromValue(v string) (*ToolsLandingGroupGetV2DataListGroupFlowType, error) {
	ev := ToolsLandingGroupGetV2DataListGroupFlowType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsLandingGroupGetV2DataListGroupFlowType: valid values are %v", v, AllowedToolsLandingGroupGetV2DataListGroupFlowTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsLandingGroupGetV2DataListGroupFlowType) IsValid() bool {
	for _, existing := range AllowedToolsLandingGroupGetV2DataListGroupFlowTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_landing_group_get_v2_data_list_group_flow_type value
func (v ToolsLandingGroupGetV2DataListGroupFlowType) Ptr() *ToolsLandingGroupGetV2DataListGroupFlowType {
	return &v
}
