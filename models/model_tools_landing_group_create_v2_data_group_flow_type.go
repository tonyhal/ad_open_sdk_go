/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsLandingGroupCreateV2DataGroupFlowType
type ToolsLandingGroupCreateV2DataGroupFlowType string

// List of tools_landing_group_create_v2_data_group_flow_type
const (
	FLOW_DISTRIBUTION_TYPE_AVERAGE_ToolsLandingGroupCreateV2DataGroupFlowType      ToolsLandingGroupCreateV2DataGroupFlowType = "FLOW_DISTRIBUTION_TYPE_AVERAGE"
	FLOW_DISTRIBUTION_TYPE_INTELLIGENCE_ToolsLandingGroupCreateV2DataGroupFlowType ToolsLandingGroupCreateV2DataGroupFlowType = "FLOW_DISTRIBUTION_TYPE_INTELLIGENCE"
)

// All allowed values of ToolsLandingGroupCreateV2DataGroupFlowType enum
var AllowedToolsLandingGroupCreateV2DataGroupFlowTypeEnumValues = []ToolsLandingGroupCreateV2DataGroupFlowType{
	"FLOW_DISTRIBUTION_TYPE_AVERAGE",
	"FLOW_DISTRIBUTION_TYPE_INTELLIGENCE",
}

// NewToolsLandingGroupCreateV2DataGroupFlowTypeFromValue returns a pointer to a valid ToolsLandingGroupCreateV2DataGroupFlowType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsLandingGroupCreateV2DataGroupFlowTypeFromValue(v string) (*ToolsLandingGroupCreateV2DataGroupFlowType, error) {
	ev := ToolsLandingGroupCreateV2DataGroupFlowType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsLandingGroupCreateV2DataGroupFlowType: valid values are %v", v, AllowedToolsLandingGroupCreateV2DataGroupFlowTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsLandingGroupCreateV2DataGroupFlowType) IsValid() bool {
	for _, existing := range AllowedToolsLandingGroupCreateV2DataGroupFlowTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_landing_group_create_v2_data_group_flow_type value
func (v ToolsLandingGroupCreateV2DataGroupFlowType) Ptr() *ToolsLandingGroupCreateV2DataGroupFlowType {
	return &v
}
