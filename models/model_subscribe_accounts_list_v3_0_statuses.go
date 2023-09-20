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

// SubscribeAccountsListV30Statuses
type SubscribeAccountsListV30Statuses string

// List of subscribe_accounts_list_v3.0_statuses
const (
	OK_SubscribeAccountsListV30Statuses           SubscribeAccountsListV30Statuses = "OK"
	PENDING_SubscribeAccountsListV30Statuses      SubscribeAccountsListV30Statuses = "PENDING"
	UNAUTHORIZED_SubscribeAccountsListV30Statuses SubscribeAccountsListV30Statuses = "UNAUTHORIZED"
	UNKNOWN_SubscribeAccountsListV30Statuses      SubscribeAccountsListV30Statuses = "UNKNOWN"
)

// All allowed values of SubscribeAccountsListV30Statuses enum
var AllowedSubscribeAccountsListV30StatusesEnumValues = []SubscribeAccountsListV30Statuses{
	"OK",
	"PENDING",
	"UNAUTHORIZED",
	"UNKNOWN",
}

// NewSubscribeAccountsListV30StatusesFromValue returns a pointer to a valid SubscribeAccountsListV30Statuses
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSubscribeAccountsListV30StatusesFromValue(v string) (*SubscribeAccountsListV30Statuses, error) {
	ev := SubscribeAccountsListV30Statuses(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SubscribeAccountsListV30Statuses: valid values are %v", v, AllowedSubscribeAccountsListV30StatusesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SubscribeAccountsListV30Statuses) IsValid() bool {
	for _, existing := range AllowedSubscribeAccountsListV30StatusesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to subscribe_accounts_list_v3.0_statuses value
func (v SubscribeAccountsListV30Statuses) Ptr() *SubscribeAccountsListV30Statuses {
	return &v
}
