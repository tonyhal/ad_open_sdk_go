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

// EventManagerShareGetV30DataListAllAccountType
type EventManagerShareGetV30DataListAllAccountType string

// List of event_manager_share_get_v3.0_data_list_all_account_type
const (
	AD_EventManagerShareGetV30DataListAllAccountType EventManagerShareGetV30DataListAllAccountType = "AD"
)

// All allowed values of EventManagerShareGetV30DataListAllAccountType enum
var AllowedEventManagerShareGetV30DataListAllAccountTypeEnumValues = []EventManagerShareGetV30DataListAllAccountType{
	"AD",
}

// NewEventManagerShareGetV30DataListAllAccountTypeFromValue returns a pointer to a valid EventManagerShareGetV30DataListAllAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerShareGetV30DataListAllAccountTypeFromValue(v string) (*EventManagerShareGetV30DataListAllAccountType, error) {
	ev := EventManagerShareGetV30DataListAllAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerShareGetV30DataListAllAccountType: valid values are %v", v, AllowedEventManagerShareGetV30DataListAllAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerShareGetV30DataListAllAccountType) IsValid() bool {
	for _, existing := range AllowedEventManagerShareGetV30DataListAllAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_share_get_v3.0_data_list_all_account_type value
func (v EventManagerShareGetV30DataListAllAccountType) Ptr() *EventManagerShareGetV30DataListAllAccountType {
	return &v
}
