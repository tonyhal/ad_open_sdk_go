/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsEstimateAudienceV2LocationType
type ToolsEstimateAudienceV2LocationType string

// List of tools_estimate_audience_v2_location_type
const (
	TRAVEL_ToolsEstimateAudienceV2LocationType  ToolsEstimateAudienceV2LocationType = "TRAVEL"
	HOME_ToolsEstimateAudienceV2LocationType    ToolsEstimateAudienceV2LocationType = "HOME"
	ALL_ToolsEstimateAudienceV2LocationType     ToolsEstimateAudienceV2LocationType = "ALL"
	CURRENT_ToolsEstimateAudienceV2LocationType ToolsEstimateAudienceV2LocationType = "CURRENT"
)

// All allowed values of ToolsEstimateAudienceV2LocationType enum
var AllowedToolsEstimateAudienceV2LocationTypeEnumValues = []ToolsEstimateAudienceV2LocationType{
	"TRAVEL",
	"HOME",
	"ALL",
	"CURRENT",
}

// NewToolsEstimateAudienceV2LocationTypeFromValue returns a pointer to a valid ToolsEstimateAudienceV2LocationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEstimateAudienceV2LocationTypeFromValue(v string) (*ToolsEstimateAudienceV2LocationType, error) {
	ev := ToolsEstimateAudienceV2LocationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEstimateAudienceV2LocationType: valid values are %v", v, AllowedToolsEstimateAudienceV2LocationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEstimateAudienceV2LocationType) IsValid() bool {
	for _, existing := range AllowedToolsEstimateAudienceV2LocationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_estimate_audience_v2_location_type value
func (v ToolsEstimateAudienceV2LocationType) Ptr() *ToolsEstimateAudienceV2LocationType {
	return &v
}
