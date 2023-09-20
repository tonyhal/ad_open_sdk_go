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

// EventManagerOptimizedGoalGetV2V30DataGoalsAssetTypes
type EventManagerOptimizedGoalGetV2V30DataGoalsAssetTypes string

// List of event_manager_optimized_goal_get_v2_v3.0_data_goals_asset_types
const (
	APP_EventManagerOptimizedGoalGetV2V30DataGoalsAssetTypes          EventManagerOptimizedGoalGetV2V30DataGoalsAssetTypes = "APP"
	MINI_PROGRAM_EventManagerOptimizedGoalGetV2V30DataGoalsAssetTypes EventManagerOptimizedGoalGetV2V30DataGoalsAssetTypes = "MINI_PROGRAM"
	ORANGE_EventManagerOptimizedGoalGetV2V30DataGoalsAssetTypes       EventManagerOptimizedGoalGetV2V30DataGoalsAssetTypes = "ORANGE"
	QUICK_APP_EventManagerOptimizedGoalGetV2V30DataGoalsAssetTypes    EventManagerOptimizedGoalGetV2V30DataGoalsAssetTypes = "QUICK_APP"
	THIRDPARTY_EventManagerOptimizedGoalGetV2V30DataGoalsAssetTypes   EventManagerOptimizedGoalGetV2V30DataGoalsAssetTypes = "THIRDPARTY"
)

// All allowed values of EventManagerOptimizedGoalGetV2V30DataGoalsAssetTypes enum
var AllowedEventManagerOptimizedGoalGetV2V30DataGoalsAssetTypesEnumValues = []EventManagerOptimizedGoalGetV2V30DataGoalsAssetTypes{
	"APP",
	"MINI_PROGRAM",
	"ORANGE",
	"QUICK_APP",
	"THIRDPARTY",
}

// NewEventManagerOptimizedGoalGetV2V30DataGoalsAssetTypesFromValue returns a pointer to a valid EventManagerOptimizedGoalGetV2V30DataGoalsAssetTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerOptimizedGoalGetV2V30DataGoalsAssetTypesFromValue(v string) (*EventManagerOptimizedGoalGetV2V30DataGoalsAssetTypes, error) {
	ev := EventManagerOptimizedGoalGetV2V30DataGoalsAssetTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerOptimizedGoalGetV2V30DataGoalsAssetTypes: valid values are %v", v, AllowedEventManagerOptimizedGoalGetV2V30DataGoalsAssetTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerOptimizedGoalGetV2V30DataGoalsAssetTypes) IsValid() bool {
	for _, existing := range AllowedEventManagerOptimizedGoalGetV2V30DataGoalsAssetTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_optimized_goal_get_v2_v3.0_data_goals_asset_types value
func (v EventManagerOptimizedGoalGetV2V30DataGoalsAssetTypes) Ptr() *EventManagerOptimizedGoalGetV2V30DataGoalsAssetTypes {
	return &v
}
