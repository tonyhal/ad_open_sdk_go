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

// ToolsEstimateAudienceV2RetargetingType
type ToolsEstimateAudienceV2RetargetingType string

// List of tools_estimate_audience_v2_retargeting_type
const (
	NONE_ToolsEstimateAudienceV2RetargetingType                ToolsEstimateAudienceV2RetargetingType = "NONE"
	RETARGETING_EXCLUDE_ToolsEstimateAudienceV2RetargetingType ToolsEstimateAudienceV2RetargetingType = "RETARGETING_EXCLUDE"
	RETARGETING_INCLUDE_ToolsEstimateAudienceV2RetargetingType ToolsEstimateAudienceV2RetargetingType = "RETARGETING_INCLUDE"
	RETARGETING_NONE_ToolsEstimateAudienceV2RetargetingType    ToolsEstimateAudienceV2RetargetingType = "RETARGETING_NONE"
)

// All allowed values of ToolsEstimateAudienceV2RetargetingType enum
var AllowedToolsEstimateAudienceV2RetargetingTypeEnumValues = []ToolsEstimateAudienceV2RetargetingType{
	"NONE",
	"RETARGETING_EXCLUDE",
	"RETARGETING_INCLUDE",
	"RETARGETING_NONE",
}

// NewToolsEstimateAudienceV2RetargetingTypeFromValue returns a pointer to a valid ToolsEstimateAudienceV2RetargetingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEstimateAudienceV2RetargetingTypeFromValue(v string) (*ToolsEstimateAudienceV2RetargetingType, error) {
	ev := ToolsEstimateAudienceV2RetargetingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEstimateAudienceV2RetargetingType: valid values are %v", v, AllowedToolsEstimateAudienceV2RetargetingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEstimateAudienceV2RetargetingType) IsValid() bool {
	for _, existing := range AllowedToolsEstimateAudienceV2RetargetingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_estimate_audience_v2_retargeting_type value
func (v ToolsEstimateAudienceV2RetargetingType) Ptr() *ToolsEstimateAudienceV2RetargetingType {
	return &v
}
