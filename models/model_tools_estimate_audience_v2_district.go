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

// ToolsEstimateAudienceV2District
type ToolsEstimateAudienceV2District string

// List of tools_estimate_audience_v2_district
const (
	CITY_ToolsEstimateAudienceV2District              ToolsEstimateAudienceV2District = "CITY"
	NONE_ToolsEstimateAudienceV2District              ToolsEstimateAudienceV2District = "NONE"
	COUNTY_ToolsEstimateAudienceV2District            ToolsEstimateAudienceV2District = "COUNTY"
	BUSINESS_DISTRICT_ToolsEstimateAudienceV2District ToolsEstimateAudienceV2District = "BUSINESS_DISTRICT"
	REGION_ToolsEstimateAudienceV2District            ToolsEstimateAudienceV2District = "REGION"
	OVERSEA_ToolsEstimateAudienceV2District           ToolsEstimateAudienceV2District = "OVERSEA"
)

// All allowed values of ToolsEstimateAudienceV2District enum
var AllowedToolsEstimateAudienceV2DistrictEnumValues = []ToolsEstimateAudienceV2District{
	"CITY",
	"NONE",
	"COUNTY",
	"BUSINESS_DISTRICT",
	"REGION",
	"OVERSEA",
}

// NewToolsEstimateAudienceV2DistrictFromValue returns a pointer to a valid ToolsEstimateAudienceV2District
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEstimateAudienceV2DistrictFromValue(v string) (*ToolsEstimateAudienceV2District, error) {
	ev := ToolsEstimateAudienceV2District(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEstimateAudienceV2District: valid values are %v", v, AllowedToolsEstimateAudienceV2DistrictEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEstimateAudienceV2District) IsValid() bool {
	for _, existing := range AllowedToolsEstimateAudienceV2DistrictEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_estimate_audience_v2_district value
func (v ToolsEstimateAudienceV2District) Ptr() *ToolsEstimateAudienceV2District {
	return &v
}
