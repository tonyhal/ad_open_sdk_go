/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// EventManagerEventConfigsGetV2SortType
type EventManagerEventConfigsGetV2SortType string

// List of event_manager_event_configs_get_v2_sort_type
const (
	DESC_EventManagerEventConfigsGetV2SortType EventManagerEventConfigsGetV2SortType = "DESC"
	ASC_EventManagerEventConfigsGetV2SortType  EventManagerEventConfigsGetV2SortType = "ASC"
)

// All allowed values of EventManagerEventConfigsGetV2SortType enum
var AllowedEventManagerEventConfigsGetV2SortTypeEnumValues = []EventManagerEventConfigsGetV2SortType{
	"DESC",
	"ASC",
}

// NewEventManagerEventConfigsGetV2SortTypeFromValue returns a pointer to a valid EventManagerEventConfigsGetV2SortType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerEventConfigsGetV2SortTypeFromValue(v string) (*EventManagerEventConfigsGetV2SortType, error) {
	ev := EventManagerEventConfigsGetV2SortType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerEventConfigsGetV2SortType: valid values are %v", v, AllowedEventManagerEventConfigsGetV2SortTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerEventConfigsGetV2SortType) IsValid() bool {
	for _, existing := range AllowedEventManagerEventConfigsGetV2SortTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_event_configs_get_v2_sort_type value
func (v EventManagerEventConfigsGetV2SortType) Ptr() *EventManagerEventConfigsGetV2SortType {
	return &v
}
