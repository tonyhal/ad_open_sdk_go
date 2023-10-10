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

// ToolsEstimateAudienceV2Ac
type ToolsEstimateAudienceV2Ac string

// List of tools_estimate_audience_v2_ac
const (
	WIFI_ToolsEstimateAudienceV2Ac     ToolsEstimateAudienceV2Ac = "WIFI"
	Enum_5_G_ToolsEstimateAudienceV2Ac ToolsEstimateAudienceV2Ac = "5G"
	Enum_2_G_ToolsEstimateAudienceV2Ac ToolsEstimateAudienceV2Ac = "2G"
	Enum_4_G_ToolsEstimateAudienceV2Ac ToolsEstimateAudienceV2Ac = "4G"
	Enum_3_G_ToolsEstimateAudienceV2Ac ToolsEstimateAudienceV2Ac = "3G"
)

// All allowed values of ToolsEstimateAudienceV2Ac enum
var AllowedToolsEstimateAudienceV2AcEnumValues = []ToolsEstimateAudienceV2Ac{
	"WIFI",
	"5G",
	"2G",
	"4G",
	"3G",
}

// NewToolsEstimateAudienceV2AcFromValue returns a pointer to a valid ToolsEstimateAudienceV2Ac
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEstimateAudienceV2AcFromValue(v string) (*ToolsEstimateAudienceV2Ac, error) {
	ev := ToolsEstimateAudienceV2Ac(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEstimateAudienceV2Ac: valid values are %v", v, AllowedToolsEstimateAudienceV2AcEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEstimateAudienceV2Ac) IsValid() bool {
	for _, existing := range AllowedToolsEstimateAudienceV2AcEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_estimate_audience_v2_ac value
func (v ToolsEstimateAudienceV2Ac) Ptr() *ToolsEstimateAudienceV2Ac {
	return &v
}
