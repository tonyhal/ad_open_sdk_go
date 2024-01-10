/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// EventManagerShareV30DataErrorListAccountInfoAccountType
type EventManagerShareV30DataErrorListAccountInfoAccountType string

// List of event_manager_share_v3.0_data_error_list_account_info_account_type
const (
	AD_EventManagerShareV30DataErrorListAccountInfoAccountType EventManagerShareV30DataErrorListAccountInfoAccountType = "AD"
)

// All allowed values of EventManagerShareV30DataErrorListAccountInfoAccountType enum
var AllowedEventManagerShareV30DataErrorListAccountInfoAccountTypeEnumValues = []EventManagerShareV30DataErrorListAccountInfoAccountType{
	"AD",
}

// NewEventManagerShareV30DataErrorListAccountInfoAccountTypeFromValue returns a pointer to a valid EventManagerShareV30DataErrorListAccountInfoAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerShareV30DataErrorListAccountInfoAccountTypeFromValue(v string) (*EventManagerShareV30DataErrorListAccountInfoAccountType, error) {
	ev := EventManagerShareV30DataErrorListAccountInfoAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerShareV30DataErrorListAccountInfoAccountType: valid values are %v", v, AllowedEventManagerShareV30DataErrorListAccountInfoAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerShareV30DataErrorListAccountInfoAccountType) IsValid() bool {
	for _, existing := range AllowedEventManagerShareV30DataErrorListAccountInfoAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_share_v3.0_data_error_list_account_info_account_type value
func (v EventManagerShareV30DataErrorListAccountInfoAccountType) Ptr() *EventManagerShareV30DataErrorListAccountInfoAccountType {
	return &v
}
