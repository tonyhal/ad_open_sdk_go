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

// EventManagerOptimizedGoalGetV2V30MarketingGoal
type EventManagerOptimizedGoalGetV2V30MarketingGoal string

// List of event_manager_optimized_goal_get_v2_v3.0_marketing_goal
const (
	LIVE_EventManagerOptimizedGoalGetV2V30MarketingGoal            EventManagerOptimizedGoalGetV2V30MarketingGoal = "LIVE"
	VIDEO_AND_IMAGE_EventManagerOptimizedGoalGetV2V30MarketingGoal EventManagerOptimizedGoalGetV2V30MarketingGoal = "VIDEO_AND_IMAGE"
)

// All allowed values of EventManagerOptimizedGoalGetV2V30MarketingGoal enum
var AllowedEventManagerOptimizedGoalGetV2V30MarketingGoalEnumValues = []EventManagerOptimizedGoalGetV2V30MarketingGoal{
	"LIVE",
	"VIDEO_AND_IMAGE",
}

// NewEventManagerOptimizedGoalGetV2V30MarketingGoalFromValue returns a pointer to a valid EventManagerOptimizedGoalGetV2V30MarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerOptimizedGoalGetV2V30MarketingGoalFromValue(v string) (*EventManagerOptimizedGoalGetV2V30MarketingGoal, error) {
	ev := EventManagerOptimizedGoalGetV2V30MarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerOptimizedGoalGetV2V30MarketingGoal: valid values are %v", v, AllowedEventManagerOptimizedGoalGetV2V30MarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerOptimizedGoalGetV2V30MarketingGoal) IsValid() bool {
	for _, existing := range AllowedEventManagerOptimizedGoalGetV2V30MarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_optimized_goal_get_v2_v3.0_marketing_goal value
func (v EventManagerOptimizedGoalGetV2V30MarketingGoal) Ptr() *EventManagerOptimizedGoalGetV2V30MarketingGoal {
	return &v
}
