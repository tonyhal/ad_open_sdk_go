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

// EventManagerShareV30ShareMode
type EventManagerShareV30ShareMode string

// List of event_manager_share_v3.0_share_mode
const (
	ALL_EventManagerShareV30ShareMode  EventManagerShareV30ShareMode = "ALL"
	PART_EventManagerShareV30ShareMode EventManagerShareV30ShareMode = "PART"
)

// All allowed values of EventManagerShareV30ShareMode enum
var AllowedEventManagerShareV30ShareModeEnumValues = []EventManagerShareV30ShareMode{
	"ALL",
	"PART",
}

// NewEventManagerShareV30ShareModeFromValue returns a pointer to a valid EventManagerShareV30ShareMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerShareV30ShareModeFromValue(v string) (*EventManagerShareV30ShareMode, error) {
	ev := EventManagerShareV30ShareMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerShareV30ShareMode: valid values are %v", v, AllowedEventManagerShareV30ShareModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerShareV30ShareMode) IsValid() bool {
	for _, existing := range AllowedEventManagerShareV30ShareModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_share_v3.0_share_mode value
func (v EventManagerShareV30ShareMode) Ptr() *EventManagerShareV30ShareMode {
	return &v
}
