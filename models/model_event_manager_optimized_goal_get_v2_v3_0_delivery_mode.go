/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// EventManagerOptimizedGoalGetV2V30DeliveryMode
type EventManagerOptimizedGoalGetV2V30DeliveryMode string

// List of event_manager_optimized_goal_get_v2_v3.0_delivery_mode
const (
	MANUAL_EventManagerOptimizedGoalGetV2V30DeliveryMode     EventManagerOptimizedGoalGetV2V30DeliveryMode = "MANUAL"
	PROCEDURAL_EventManagerOptimizedGoalGetV2V30DeliveryMode EventManagerOptimizedGoalGetV2V30DeliveryMode = "PROCEDURAL"
)

// All allowed values of EventManagerOptimizedGoalGetV2V30DeliveryMode enum
var AllowedEventManagerOptimizedGoalGetV2V30DeliveryModeEnumValues = []EventManagerOptimizedGoalGetV2V30DeliveryMode{
	"MANUAL",
	"PROCEDURAL",
}

// NewEventManagerOptimizedGoalGetV2V30DeliveryModeFromValue returns a pointer to a valid EventManagerOptimizedGoalGetV2V30DeliveryMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerOptimizedGoalGetV2V30DeliveryModeFromValue(v string) (*EventManagerOptimizedGoalGetV2V30DeliveryMode, error) {
	ev := EventManagerOptimizedGoalGetV2V30DeliveryMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerOptimizedGoalGetV2V30DeliveryMode: valid values are %v", v, AllowedEventManagerOptimizedGoalGetV2V30DeliveryModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerOptimizedGoalGetV2V30DeliveryMode) IsValid() bool {
	for _, existing := range AllowedEventManagerOptimizedGoalGetV2V30DeliveryModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_optimized_goal_get_v2_v3.0_delivery_mode value
func (v EventManagerOptimizedGoalGetV2V30DeliveryMode) Ptr() *EventManagerOptimizedGoalGetV2V30DeliveryMode {
	return &v
}
