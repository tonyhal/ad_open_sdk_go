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

// ToolsInterestActionActionCategoryV2ActionDays
type ToolsInterestActionActionCategoryV2ActionDays int64

// List of tools_interest_action_action_category_v2_action_days
const (
	Enum_7_ToolsInterestActionActionCategoryV2ActionDays   ToolsInterestActionActionCategoryV2ActionDays = 7
	Enum_15_ToolsInterestActionActionCategoryV2ActionDays  ToolsInterestActionActionCategoryV2ActionDays = 15
	Enum_30_ToolsInterestActionActionCategoryV2ActionDays  ToolsInterestActionActionCategoryV2ActionDays = 30
	Enum_60_ToolsInterestActionActionCategoryV2ActionDays  ToolsInterestActionActionCategoryV2ActionDays = 60
	Enum_90_ToolsInterestActionActionCategoryV2ActionDays  ToolsInterestActionActionCategoryV2ActionDays = 90
	Enum_180_ToolsInterestActionActionCategoryV2ActionDays ToolsInterestActionActionCategoryV2ActionDays = 180
	Enum_365_ToolsInterestActionActionCategoryV2ActionDays ToolsInterestActionActionCategoryV2ActionDays = 365
)

// All allowed values of ToolsInterestActionActionCategoryV2ActionDays enum
var AllowedToolsInterestActionActionCategoryV2ActionDaysEnumValues = []ToolsInterestActionActionCategoryV2ActionDays{
	7,
	15,
	30,
	60,
	90,
	180,
	365,
}

// NewToolsInterestActionActionCategoryV2ActionDaysFromValue returns a pointer to a valid ToolsInterestActionActionCategoryV2ActionDays
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsInterestActionActionCategoryV2ActionDaysFromValue(v int64) (*ToolsInterestActionActionCategoryV2ActionDays, error) {
	ev := ToolsInterestActionActionCategoryV2ActionDays(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsInterestActionActionCategoryV2ActionDays: valid values are %v", v, AllowedToolsInterestActionActionCategoryV2ActionDaysEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsInterestActionActionCategoryV2ActionDays) IsValid() bool {
	for _, existing := range AllowedToolsInterestActionActionCategoryV2ActionDaysEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_interest_action_action_category_v2_action_days value
func (v ToolsInterestActionActionCategoryV2ActionDays) Ptr() *ToolsInterestActionActionCategoryV2ActionDays {
	return &v
}
