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

// ToolsIndustryGetV2Type
type ToolsIndustryGetV2Type string

// List of tools_industry_get_v2_type
const (
	ADVERTISER_ToolsIndustryGetV2Type    ToolsIndustryGetV2Type = "ADVERTISER"
	AGENT_ToolsIndustryGetV2Type         ToolsIndustryGetV2Type = "AGENT"
	QUALIFICATION_ToolsIndustryGetV2Type ToolsIndustryGetV2Type = "QUALIFICATION"
)

// All allowed values of ToolsIndustryGetV2Type enum
var AllowedToolsIndustryGetV2TypeEnumValues = []ToolsIndustryGetV2Type{
	"ADVERTISER",
	"AGENT",
	"QUALIFICATION",
}

// NewToolsIndustryGetV2TypeFromValue returns a pointer to a valid ToolsIndustryGetV2Type
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsIndustryGetV2TypeFromValue(v string) (*ToolsIndustryGetV2Type, error) {
	ev := ToolsIndustryGetV2Type(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsIndustryGetV2Type: valid values are %v", v, AllowedToolsIndustryGetV2TypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsIndustryGetV2Type) IsValid() bool {
	for _, existing := range AllowedToolsIndustryGetV2TypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_industry_get_v2_type value
func (v ToolsIndustryGetV2Type) Ptr() *ToolsIndustryGetV2Type {
	return &v
}
