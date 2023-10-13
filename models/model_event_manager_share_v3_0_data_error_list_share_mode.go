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

// EventManagerShareV30DataErrorListShareMode
type EventManagerShareV30DataErrorListShareMode string

// List of event_manager_share_v3.0_data_error_list_share_mode
const (
	ALL_EventManagerShareV30DataErrorListShareMode  EventManagerShareV30DataErrorListShareMode = "ALL"
	PART_EventManagerShareV30DataErrorListShareMode EventManagerShareV30DataErrorListShareMode = "PART"
)

// All allowed values of EventManagerShareV30DataErrorListShareMode enum
var AllowedEventManagerShareV30DataErrorListShareModeEnumValues = []EventManagerShareV30DataErrorListShareMode{
	"ALL",
	"PART",
}

// NewEventManagerShareV30DataErrorListShareModeFromValue returns a pointer to a valid EventManagerShareV30DataErrorListShareMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerShareV30DataErrorListShareModeFromValue(v string) (*EventManagerShareV30DataErrorListShareMode, error) {
	ev := EventManagerShareV30DataErrorListShareMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerShareV30DataErrorListShareMode: valid values are %v", v, AllowedEventManagerShareV30DataErrorListShareModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerShareV30DataErrorListShareMode) IsValid() bool {
	for _, existing := range AllowedEventManagerShareV30DataErrorListShareModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_share_v3.0_data_error_list_share_mode value
func (v EventManagerShareV30DataErrorListShareMode) Ptr() *EventManagerShareV30DataErrorListShareMode {
	return &v
}
