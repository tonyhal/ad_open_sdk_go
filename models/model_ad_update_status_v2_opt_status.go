/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdUpdateStatusV2OptStatus
type AdUpdateStatusV2OptStatus string

// List of ad_update_status_v2_opt_status
const (
	DISABLE_AdUpdateStatusV2OptStatus AdUpdateStatusV2OptStatus = "disable"
	ENABLE_AdUpdateStatusV2OptStatus  AdUpdateStatusV2OptStatus = "enable"
	DELETE_AdUpdateStatusV2OptStatus  AdUpdateStatusV2OptStatus = "delete"
)

// All allowed values of AdUpdateStatusV2OptStatus enum
var AllowedAdUpdateStatusV2OptStatusEnumValues = []AdUpdateStatusV2OptStatus{
	"disable",
	"enable",
	"delete",
}

// NewAdUpdateStatusV2OptStatusFromValue returns a pointer to a valid AdUpdateStatusV2OptStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdUpdateStatusV2OptStatusFromValue(v string) (*AdUpdateStatusV2OptStatus, error) {
	ev := AdUpdateStatusV2OptStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdUpdateStatusV2OptStatus: valid values are %v", v, AllowedAdUpdateStatusV2OptStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdUpdateStatusV2OptStatus) IsValid() bool {
	for _, existing := range AllowedAdUpdateStatusV2OptStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_update_status_v2_opt_status value
func (v AdUpdateStatusV2OptStatus) Ptr() *AdUpdateStatusV2OptStatus {
	return &v
}
