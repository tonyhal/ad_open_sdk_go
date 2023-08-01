/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DouplusOrderListV30FilterStatus
type DouplusOrderListV30FilterStatus string

// List of douplus_order_list_v3.0_filter_status
const (
	AUDITING_DouplusOrderListV30FilterStatus      DouplusOrderListV30FilterStatus = "AUDITING"
	AUDIT_PAUSE_DouplusOrderListV30FilterStatus   DouplusOrderListV30FilterStatus = "AUDIT_PAUSE"
	DELIVERIED_DouplusOrderListV30FilterStatus    DouplusOrderListV30FilterStatus = "DELIVERIED"
	DELIVERING_DouplusOrderListV30FilterStatus    DouplusOrderListV30FilterStatus = "DELIVERING"
	OFFLINE_AUDIT_DouplusOrderListV30FilterStatus DouplusOrderListV30FilterStatus = "OFFLINE_AUDIT"
	TIME_NO_REACH_DouplusOrderListV30FilterStatus DouplusOrderListV30FilterStatus = "TIME_NO_REACH"
	UNDELIVERIED_DouplusOrderListV30FilterStatus  DouplusOrderListV30FilterStatus = "UNDELIVERIED"
	UNPAID_DouplusOrderListV30FilterStatus        DouplusOrderListV30FilterStatus = "UNPAID"
	WAIT_TO_START_DouplusOrderListV30FilterStatus DouplusOrderListV30FilterStatus = "WAIT_TO_START"
)

// All allowed values of DouplusOrderListV30FilterStatus enum
var AllowedDouplusOrderListV30FilterStatusEnumValues = []DouplusOrderListV30FilterStatus{
	"AUDITING",
	"AUDIT_PAUSE",
	"DELIVERIED",
	"DELIVERING",
	"OFFLINE_AUDIT",
	"TIME_NO_REACH",
	"UNDELIVERIED",
	"UNPAID",
	"WAIT_TO_START",
}

// NewDouplusOrderListV30FilterStatusFromValue returns a pointer to a valid DouplusOrderListV30FilterStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDouplusOrderListV30FilterStatusFromValue(v string) (*DouplusOrderListV30FilterStatus, error) {
	ev := DouplusOrderListV30FilterStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DouplusOrderListV30FilterStatus: valid values are %v", v, AllowedDouplusOrderListV30FilterStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DouplusOrderListV30FilterStatus) IsValid() bool {
	for _, existing := range AllowedDouplusOrderListV30FilterStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to douplus_order_list_v3.0_filter_status value
func (v DouplusOrderListV30FilterStatus) Ptr() *DouplusOrderListV30FilterStatus {
	return &v
}