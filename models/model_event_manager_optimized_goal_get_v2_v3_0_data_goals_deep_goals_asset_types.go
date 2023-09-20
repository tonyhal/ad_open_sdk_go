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

// EventManagerOptimizedGoalGetV2V30DataGoalsDeepGoalsAssetTypes
type EventManagerOptimizedGoalGetV2V30DataGoalsDeepGoalsAssetTypes string

// List of event_manager_optimized_goal_get_v2_v3.0_data_goals_deep_goals_asset_types
const (
	APP_EventManagerOptimizedGoalGetV2V30DataGoalsDeepGoalsAssetTypes          EventManagerOptimizedGoalGetV2V30DataGoalsDeepGoalsAssetTypes = "APP"
	MINI_PROGRAM_EventManagerOptimizedGoalGetV2V30DataGoalsDeepGoalsAssetTypes EventManagerOptimizedGoalGetV2V30DataGoalsDeepGoalsAssetTypes = "MINI_PROGRAM"
	ORANGE_EventManagerOptimizedGoalGetV2V30DataGoalsDeepGoalsAssetTypes       EventManagerOptimizedGoalGetV2V30DataGoalsDeepGoalsAssetTypes = "ORANGE"
	QUICK_APP_EventManagerOptimizedGoalGetV2V30DataGoalsDeepGoalsAssetTypes    EventManagerOptimizedGoalGetV2V30DataGoalsDeepGoalsAssetTypes = "QUICK_APP"
	THIRDPARTY_EventManagerOptimizedGoalGetV2V30DataGoalsDeepGoalsAssetTypes   EventManagerOptimizedGoalGetV2V30DataGoalsDeepGoalsAssetTypes = "THIRDPARTY"
)

// All allowed values of EventManagerOptimizedGoalGetV2V30DataGoalsDeepGoalsAssetTypes enum
var AllowedEventManagerOptimizedGoalGetV2V30DataGoalsDeepGoalsAssetTypesEnumValues = []EventManagerOptimizedGoalGetV2V30DataGoalsDeepGoalsAssetTypes{
	"APP",
	"MINI_PROGRAM",
	"ORANGE",
	"QUICK_APP",
	"THIRDPARTY",
}

// NewEventManagerOptimizedGoalGetV2V30DataGoalsDeepGoalsAssetTypesFromValue returns a pointer to a valid EventManagerOptimizedGoalGetV2V30DataGoalsDeepGoalsAssetTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerOptimizedGoalGetV2V30DataGoalsDeepGoalsAssetTypesFromValue(v string) (*EventManagerOptimizedGoalGetV2V30DataGoalsDeepGoalsAssetTypes, error) {
	ev := EventManagerOptimizedGoalGetV2V30DataGoalsDeepGoalsAssetTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerOptimizedGoalGetV2V30DataGoalsDeepGoalsAssetTypes: valid values are %v", v, AllowedEventManagerOptimizedGoalGetV2V30DataGoalsDeepGoalsAssetTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerOptimizedGoalGetV2V30DataGoalsDeepGoalsAssetTypes) IsValid() bool {
	for _, existing := range AllowedEventManagerOptimizedGoalGetV2V30DataGoalsDeepGoalsAssetTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_optimized_goal_get_v2_v3.0_data_goals_deep_goals_asset_types value
func (v EventManagerOptimizedGoalGetV2V30DataGoalsDeepGoalsAssetTypes) Ptr() *EventManagerOptimizedGoalGetV2V30DataGoalsDeepGoalsAssetTypes {
	return &v
}
