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

// EventManagerOptimizedGoalGetV2V30AppType
type EventManagerOptimizedGoalGetV2V30AppType string

// List of event_manager_optimized_goal_get_v2_v3.0_app_type
const (
	ANDROID_EventManagerOptimizedGoalGetV2V30AppType EventManagerOptimizedGoalGetV2V30AppType = "ANDROID"
	DEFAULT_EventManagerOptimizedGoalGetV2V30AppType EventManagerOptimizedGoalGetV2V30AppType = "DEFAULT"
	IOS_EventManagerOptimizedGoalGetV2V30AppType     EventManagerOptimizedGoalGetV2V30AppType = "IOS"
)

// All allowed values of EventManagerOptimizedGoalGetV2V30AppType enum
var AllowedEventManagerOptimizedGoalGetV2V30AppTypeEnumValues = []EventManagerOptimizedGoalGetV2V30AppType{
	"ANDROID",
	"DEFAULT",
	"IOS",
}

// NewEventManagerOptimizedGoalGetV2V30AppTypeFromValue returns a pointer to a valid EventManagerOptimizedGoalGetV2V30AppType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerOptimizedGoalGetV2V30AppTypeFromValue(v string) (*EventManagerOptimizedGoalGetV2V30AppType, error) {
	ev := EventManagerOptimizedGoalGetV2V30AppType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerOptimizedGoalGetV2V30AppType: valid values are %v", v, AllowedEventManagerOptimizedGoalGetV2V30AppTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerOptimizedGoalGetV2V30AppType) IsValid() bool {
	for _, existing := range AllowedEventManagerOptimizedGoalGetV2V30AppTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_optimized_goal_get_v2_v3.0_app_type value
func (v EventManagerOptimizedGoalGetV2V30AppType) Ptr() *EventManagerOptimizedGoalGetV2V30AppType {
	return &v
}
