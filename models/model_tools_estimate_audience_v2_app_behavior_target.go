/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsEstimateAudienceV2AppBehaviorTarget
type ToolsEstimateAudienceV2AppBehaviorTarget string

// List of tools_estimate_audience_v2_app_behavior_target
const (
	APP_ToolsEstimateAudienceV2AppBehaviorTarget      ToolsEstimateAudienceV2AppBehaviorTarget = "APP"
	CATEGORY_ToolsEstimateAudienceV2AppBehaviorTarget ToolsEstimateAudienceV2AppBehaviorTarget = "CATEGORY"
	NONE_ToolsEstimateAudienceV2AppBehaviorTarget     ToolsEstimateAudienceV2AppBehaviorTarget = "NONE"
)

// All allowed values of ToolsEstimateAudienceV2AppBehaviorTarget enum
var AllowedToolsEstimateAudienceV2AppBehaviorTargetEnumValues = []ToolsEstimateAudienceV2AppBehaviorTarget{
	"APP",
	"CATEGORY",
	"NONE",
}

// NewToolsEstimateAudienceV2AppBehaviorTargetFromValue returns a pointer to a valid ToolsEstimateAudienceV2AppBehaviorTarget
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEstimateAudienceV2AppBehaviorTargetFromValue(v string) (*ToolsEstimateAudienceV2AppBehaviorTarget, error) {
	ev := ToolsEstimateAudienceV2AppBehaviorTarget(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEstimateAudienceV2AppBehaviorTarget: valid values are %v", v, AllowedToolsEstimateAudienceV2AppBehaviorTargetEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEstimateAudienceV2AppBehaviorTarget) IsValid() bool {
	for _, existing := range AllowedToolsEstimateAudienceV2AppBehaviorTargetEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_estimate_audience_v2_app_behavior_target value
func (v ToolsEstimateAudienceV2AppBehaviorTarget) Ptr() *ToolsEstimateAudienceV2AppBehaviorTarget {
	return &v
}
