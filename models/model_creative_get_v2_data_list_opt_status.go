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

// CreativeGetV2DataListOptStatus
type CreativeGetV2DataListOptStatus string

// List of creative_get_v2_data_list_opt_status
const (
	CREATIVE_STATUS_DISABLE_CreativeGetV2DataListOptStatus     CreativeGetV2DataListOptStatus = "CREATIVE_STATUS_DISABLE"
	CREATIVE_STATUS_ENABLE_CreativeGetV2DataListOptStatus      CreativeGetV2DataListOptStatus = "CREATIVE_STATUS_ENABLE"
	CREATIVE_STATUS_OTHER_CreativeGetV2DataListOptStatus       CreativeGetV2DataListOptStatus = "CREATIVE_STATUS_OTHER"
	CREATIVE_STATUS_QUOTA_LIMIT_CreativeGetV2DataListOptStatus CreativeGetV2DataListOptStatus = "CREATIVE_STATUS_QUOTA_LIMIT"
)

// All allowed values of CreativeGetV2DataListOptStatus enum
var AllowedCreativeGetV2DataListOptStatusEnumValues = []CreativeGetV2DataListOptStatus{
	"CREATIVE_STATUS_DISABLE",
	"CREATIVE_STATUS_ENABLE",
	"CREATIVE_STATUS_OTHER",
	"CREATIVE_STATUS_QUOTA_LIMIT",
}

// NewCreativeGetV2DataListOptStatusFromValue returns a pointer to a valid CreativeGetV2DataListOptStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeGetV2DataListOptStatusFromValue(v string) (*CreativeGetV2DataListOptStatus, error) {
	ev := CreativeGetV2DataListOptStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeGetV2DataListOptStatus: valid values are %v", v, AllowedCreativeGetV2DataListOptStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeGetV2DataListOptStatus) IsValid() bool {
	for _, existing := range AllowedCreativeGetV2DataListOptStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_get_v2_data_list_opt_status value
func (v CreativeGetV2DataListOptStatus) Ptr() *CreativeGetV2DataListOptStatus {
	return &v
}
