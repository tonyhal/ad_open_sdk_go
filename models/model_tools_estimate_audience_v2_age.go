/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsEstimateAudienceV2Age
type ToolsEstimateAudienceV2Age string

// List of tools_estimate_audience_v2_age
const (
	AGE_BELOW_18_ToolsEstimateAudienceV2Age      ToolsEstimateAudienceV2Age = "AGE_BELOW_18"
	AGE_ABOVE_50_ToolsEstimateAudienceV2Age      ToolsEstimateAudienceV2Age = "AGE_ABOVE_50"
	AGE_BETWEEN_41_49_ToolsEstimateAudienceV2Age ToolsEstimateAudienceV2Age = "AGE_BETWEEN_41_49"
	AGE_BETWEEN_18_23_ToolsEstimateAudienceV2Age ToolsEstimateAudienceV2Age = "AGE_BETWEEN_18_23"
	AGE_BETWEEN_31_40_ToolsEstimateAudienceV2Age ToolsEstimateAudienceV2Age = "AGE_BETWEEN_31_40"
	AGE_BETWEEN_24_30_ToolsEstimateAudienceV2Age ToolsEstimateAudienceV2Age = "AGE_BETWEEN_24_30"
)

// All allowed values of ToolsEstimateAudienceV2Age enum
var AllowedToolsEstimateAudienceV2AgeEnumValues = []ToolsEstimateAudienceV2Age{
	"AGE_BELOW_18",
	"AGE_ABOVE_50",
	"AGE_BETWEEN_41_49",
	"AGE_BETWEEN_18_23",
	"AGE_BETWEEN_31_40",
	"AGE_BETWEEN_24_30",
}

// NewToolsEstimateAudienceV2AgeFromValue returns a pointer to a valid ToolsEstimateAudienceV2Age
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEstimateAudienceV2AgeFromValue(v string) (*ToolsEstimateAudienceV2Age, error) {
	ev := ToolsEstimateAudienceV2Age(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEstimateAudienceV2Age: valid values are %v", v, AllowedToolsEstimateAudienceV2AgeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEstimateAudienceV2Age) IsValid() bool {
	for _, existing := range AllowedToolsEstimateAudienceV2AgeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_estimate_audience_v2_age value
func (v ToolsEstimateAudienceV2Age) Ptr() *ToolsEstimateAudienceV2Age {
	return &v
}
