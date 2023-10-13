/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdStatusUpdateV10OptStatus
type QianchuanAdStatusUpdateV10OptStatus string

// List of qianchuan_ad_status_update_v1.0_opt_status
const (
	DISABLE_QianchuanAdStatusUpdateV10OptStatus QianchuanAdStatusUpdateV10OptStatus = "DISABLE"
	ENABLE_QianchuanAdStatusUpdateV10OptStatus  QianchuanAdStatusUpdateV10OptStatus = "ENABLE"
	DELETE_QianchuanAdStatusUpdateV10OptStatus  QianchuanAdStatusUpdateV10OptStatus = "DELETE"
	REVIVE_QianchuanAdStatusUpdateV10OptStatus  QianchuanAdStatusUpdateV10OptStatus = "REVIVE"
)

// All allowed values of QianchuanAdStatusUpdateV10OptStatus enum
var AllowedQianchuanAdStatusUpdateV10OptStatusEnumValues = []QianchuanAdStatusUpdateV10OptStatus{
	"DISABLE",
	"ENABLE",
	"DELETE",
	"REVIVE",
}

// NewQianchuanAdStatusUpdateV10OptStatusFromValue returns a pointer to a valid QianchuanAdStatusUpdateV10OptStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdStatusUpdateV10OptStatusFromValue(v string) (*QianchuanAdStatusUpdateV10OptStatus, error) {
	ev := QianchuanAdStatusUpdateV10OptStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdStatusUpdateV10OptStatus: valid values are %v", v, AllowedQianchuanAdStatusUpdateV10OptStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdStatusUpdateV10OptStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanAdStatusUpdateV10OptStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_status_update_v1.0_opt_status value
func (v QianchuanAdStatusUpdateV10OptStatus) Ptr() *QianchuanAdStatusUpdateV10OptStatus {
	return &v
}
