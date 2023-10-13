/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// EventManagerOptimizedGoalGetV2V30DataGoalsValueType
type EventManagerOptimizedGoalGetV2V30DataGoalsValueType string

// List of event_manager_optimized_goal_get_v2_v3.0_data_goals_value_type
const (
	DISABLED_EventManagerOptimizedGoalGetV2V30DataGoalsValueType            EventManagerOptimizedGoalGetV2V30DataGoalsValueType = "DISABLED"
	DISCRIMINATEBYGROUP_EventManagerOptimizedGoalGetV2V30DataGoalsValueType EventManagerOptimizedGoalGetV2V30DataGoalsValueType = "DISCRIMINATEBYGROUP"
	DYNAMICVALUE_EventManagerOptimizedGoalGetV2V30DataGoalsValueType        EventManagerOptimizedGoalGetV2V30DataGoalsValueType = "DYNAMICVALUE"
	FIXED_EventManagerOptimizedGoalGetV2V30DataGoalsValueType               EventManagerOptimizedGoalGetV2V30DataGoalsValueType = "FIXED"
	UNSET_EventManagerOptimizedGoalGetV2V30DataGoalsValueType               EventManagerOptimizedGoalGetV2V30DataGoalsValueType = "UNSET"
)

// All allowed values of EventManagerOptimizedGoalGetV2V30DataGoalsValueType enum
var AllowedEventManagerOptimizedGoalGetV2V30DataGoalsValueTypeEnumValues = []EventManagerOptimizedGoalGetV2V30DataGoalsValueType{
	"DISABLED",
	"DISCRIMINATEBYGROUP",
	"DYNAMICVALUE",
	"FIXED",
	"UNSET",
}

// NewEventManagerOptimizedGoalGetV2V30DataGoalsValueTypeFromValue returns a pointer to a valid EventManagerOptimizedGoalGetV2V30DataGoalsValueType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerOptimizedGoalGetV2V30DataGoalsValueTypeFromValue(v string) (*EventManagerOptimizedGoalGetV2V30DataGoalsValueType, error) {
	ev := EventManagerOptimizedGoalGetV2V30DataGoalsValueType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerOptimizedGoalGetV2V30DataGoalsValueType: valid values are %v", v, AllowedEventManagerOptimizedGoalGetV2V30DataGoalsValueTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerOptimizedGoalGetV2V30DataGoalsValueType) IsValid() bool {
	for _, existing := range AllowedEventManagerOptimizedGoalGetV2V30DataGoalsValueTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_optimized_goal_get_v2_v3.0_data_goals_value_type value
func (v EventManagerOptimizedGoalGetV2V30DataGoalsValueType) Ptr() *EventManagerOptimizedGoalGetV2V30DataGoalsValueType {
	return &v
}
