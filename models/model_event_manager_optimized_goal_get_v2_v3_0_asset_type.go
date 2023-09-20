/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// EventManagerOptimizedGoalGetV2V30AssetType
type EventManagerOptimizedGoalGetV2V30AssetType string

// List of event_manager_optimized_goal_get_v2_v3.0_asset_type
const (
	APP_EventManagerOptimizedGoalGetV2V30AssetType          EventManagerOptimizedGoalGetV2V30AssetType = "APP"
	ENTERPRISE_EventManagerOptimizedGoalGetV2V30AssetType   EventManagerOptimizedGoalGetV2V30AssetType = "ENTERPRISE"
	MINI_PROGRAM_EventManagerOptimizedGoalGetV2V30AssetType EventManagerOptimizedGoalGetV2V30AssetType = "MINI_PROGRAM"
	ORANGE_EventManagerOptimizedGoalGetV2V30AssetType       EventManagerOptimizedGoalGetV2V30AssetType = "ORANGE"
	QUICK_APP_EventManagerOptimizedGoalGetV2V30AssetType    EventManagerOptimizedGoalGetV2V30AssetType = "QUICK_APP"
	THIRDPARTY_EventManagerOptimizedGoalGetV2V30AssetType   EventManagerOptimizedGoalGetV2V30AssetType = "THIRDPARTY"
	WECHAT_APP_EventManagerOptimizedGoalGetV2V30AssetType   EventManagerOptimizedGoalGetV2V30AssetType = "WECHAT_APP"
)

// All allowed values of EventManagerOptimizedGoalGetV2V30AssetType enum
var AllowedEventManagerOptimizedGoalGetV2V30AssetTypeEnumValues = []EventManagerOptimizedGoalGetV2V30AssetType{
	"APP",
	"ENTERPRISE",
	"MINI_PROGRAM",
	"ORANGE",
	"QUICK_APP",
	"THIRDPARTY",
	"WECHAT_APP",
}

// NewEventManagerOptimizedGoalGetV2V30AssetTypeFromValue returns a pointer to a valid EventManagerOptimizedGoalGetV2V30AssetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerOptimizedGoalGetV2V30AssetTypeFromValue(v string) (*EventManagerOptimizedGoalGetV2V30AssetType, error) {
	ev := EventManagerOptimizedGoalGetV2V30AssetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerOptimizedGoalGetV2V30AssetType: valid values are %v", v, AllowedEventManagerOptimizedGoalGetV2V30AssetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerOptimizedGoalGetV2V30AssetType) IsValid() bool {
	for _, existing := range AllowedEventManagerOptimizedGoalGetV2V30AssetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_optimized_goal_get_v2_v3.0_asset_type value
func (v EventManagerOptimizedGoalGetV2V30AssetType) Ptr() *EventManagerOptimizedGoalGetV2V30AssetType {
	return &v
}
