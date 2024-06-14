/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// EventManagerShareGetV30DataListAccountInfoAccountType
type EventManagerShareGetV30DataListAccountInfoAccountType string

// List of event_manager_share_get_v3.0_data_list_account_info_account_type
const (
	AD_EventManagerShareGetV30DataListAccountInfoAccountType EventManagerShareGetV30DataListAccountInfoAccountType = "AD"
)

// All allowed values of EventManagerShareGetV30DataListAccountInfoAccountType enum
var AllowedEventManagerShareGetV30DataListAccountInfoAccountTypeEnumValues = []EventManagerShareGetV30DataListAccountInfoAccountType{
	"AD",
}

// NewEventManagerShareGetV30DataListAccountInfoAccountTypeFromValue returns a pointer to a valid EventManagerShareGetV30DataListAccountInfoAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerShareGetV30DataListAccountInfoAccountTypeFromValue(v string) (*EventManagerShareGetV30DataListAccountInfoAccountType, error) {
	ev := EventManagerShareGetV30DataListAccountInfoAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerShareGetV30DataListAccountInfoAccountType: valid values are %v", v, AllowedEventManagerShareGetV30DataListAccountInfoAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerShareGetV30DataListAccountInfoAccountType) IsValid() bool {
	for _, existing := range AllowedEventManagerShareGetV30DataListAccountInfoAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_share_get_v3.0_data_list_account_info_account_type value
func (v EventManagerShareGetV30DataListAccountInfoAccountType) Ptr() *EventManagerShareGetV30DataListAccountInfoAccountType {
	return &v
}
