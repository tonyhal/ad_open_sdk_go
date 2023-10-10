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

// ToolsAppManagementBookingRecordsGetV2HostType
type ToolsAppManagementBookingRecordsGetV2HostType string

// List of tools_app_management_booking_records_get_v2_host_type
const (
	HOST_IN_ToolsAppManagementBookingRecordsGetV2HostType  ToolsAppManagementBookingRecordsGetV2HostType = "HOST_IN"
	HOST_OUT_ToolsAppManagementBookingRecordsGetV2HostType ToolsAppManagementBookingRecordsGetV2HostType = "HOST_OUT"
)

// All allowed values of ToolsAppManagementBookingRecordsGetV2HostType enum
var AllowedToolsAppManagementBookingRecordsGetV2HostTypeEnumValues = []ToolsAppManagementBookingRecordsGetV2HostType{
	"HOST_IN",
	"HOST_OUT",
}

// NewToolsAppManagementBookingRecordsGetV2HostTypeFromValue returns a pointer to a valid ToolsAppManagementBookingRecordsGetV2HostType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementBookingRecordsGetV2HostTypeFromValue(v string) (*ToolsAppManagementBookingRecordsGetV2HostType, error) {
	ev := ToolsAppManagementBookingRecordsGetV2HostType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementBookingRecordsGetV2HostType: valid values are %v", v, AllowedToolsAppManagementBookingRecordsGetV2HostTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementBookingRecordsGetV2HostType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementBookingRecordsGetV2HostTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_booking_records_get_v2_host_type value
func (v ToolsAppManagementBookingRecordsGetV2HostType) Ptr() *ToolsAppManagementBookingRecordsGetV2HostType {
	return &v
}
