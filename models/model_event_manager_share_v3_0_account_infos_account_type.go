/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// EventManagerShareV30AccountInfosAccountType
type EventManagerShareV30AccountInfosAccountType string

// List of event_manager_share_v3.0_account_infos_account_type
const (
	AD_EventManagerShareV30AccountInfosAccountType EventManagerShareV30AccountInfosAccountType = "AD"
)

// All allowed values of EventManagerShareV30AccountInfosAccountType enum
var AllowedEventManagerShareV30AccountInfosAccountTypeEnumValues = []EventManagerShareV30AccountInfosAccountType{
	"AD",
}

// NewEventManagerShareV30AccountInfosAccountTypeFromValue returns a pointer to a valid EventManagerShareV30AccountInfosAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerShareV30AccountInfosAccountTypeFromValue(v string) (*EventManagerShareV30AccountInfosAccountType, error) {
	ev := EventManagerShareV30AccountInfosAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerShareV30AccountInfosAccountType: valid values are %v", v, AllowedEventManagerShareV30AccountInfosAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerShareV30AccountInfosAccountType) IsValid() bool {
	for _, existing := range AllowedEventManagerShareV30AccountInfosAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_share_v3.0_account_infos_account_type value
func (v EventManagerShareV30AccountInfosAccountType) Ptr() *EventManagerShareV30AccountInfosAccountType {
	return &v
}
