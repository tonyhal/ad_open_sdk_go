/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsEstimateAudienceV2ActivateType
type ToolsEstimateAudienceV2ActivateType string

// List of tools_estimate_audience_v2_activate_type
const (
	WITH_IN_A_MONTH_ToolsEstimateAudienceV2ActivateType         ToolsEstimateAudienceV2ActivateType = "WITH_IN_A_MONTH"
	UNLIMITED_ToolsEstimateAudienceV2ActivateType               ToolsEstimateAudienceV2ActivateType = "UNLIMITED"
	THREE_MONTH_EAILIER_ToolsEstimateAudienceV2ActivateType     ToolsEstimateAudienceV2ActivateType = "THREE_MONTH_EAILIER"
	ONE_MONTH_2_THREE_MONTH_ToolsEstimateAudienceV2ActivateType ToolsEstimateAudienceV2ActivateType = "ONE_MONTH_2_THREE_MONTH"
)

// All allowed values of ToolsEstimateAudienceV2ActivateType enum
var AllowedToolsEstimateAudienceV2ActivateTypeEnumValues = []ToolsEstimateAudienceV2ActivateType{
	"WITH_IN_A_MONTH",
	"UNLIMITED",
	"THREE_MONTH_EAILIER",
	"ONE_MONTH_2_THREE_MONTH",
}

// NewToolsEstimateAudienceV2ActivateTypeFromValue returns a pointer to a valid ToolsEstimateAudienceV2ActivateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEstimateAudienceV2ActivateTypeFromValue(v string) (*ToolsEstimateAudienceV2ActivateType, error) {
	ev := ToolsEstimateAudienceV2ActivateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEstimateAudienceV2ActivateType: valid values are %v", v, AllowedToolsEstimateAudienceV2ActivateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEstimateAudienceV2ActivateType) IsValid() bool {
	for _, existing := range AllowedToolsEstimateAudienceV2ActivateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_estimate_audience_v2_activate_type value
func (v ToolsEstimateAudienceV2ActivateType) Ptr() *ToolsEstimateAudienceV2ActivateType {
	return &v
}
