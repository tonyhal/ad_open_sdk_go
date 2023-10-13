/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdDetailGetV10DataOptStatus
type QianchuanAdDetailGetV10DataOptStatus string

// List of qianchuan_ad_detail_get_v1.0_data_opt_status
const (
	DELETE_QianchuanAdDetailGetV10DataOptStatus        QianchuanAdDetailGetV10DataOptStatus = "DELETE"
	DISABLE_QianchuanAdDetailGetV10DataOptStatus       QianchuanAdDetailGetV10DataOptStatus = "DISABLE"
	ENABLE_QianchuanAdDetailGetV10DataOptStatus        QianchuanAdDetailGetV10DataOptStatus = "ENABLE"
	QUOTA_DISABLE_QianchuanAdDetailGetV10DataOptStatus QianchuanAdDetailGetV10DataOptStatus = "QUOTA_DISABLE"
	ROI2_DISABLE_QianchuanAdDetailGetV10DataOptStatus  QianchuanAdDetailGetV10DataOptStatus = "ROI2_DISABLE"
)

// All allowed values of QianchuanAdDetailGetV10DataOptStatus enum
var AllowedQianchuanAdDetailGetV10DataOptStatusEnumValues = []QianchuanAdDetailGetV10DataOptStatus{
	"DELETE",
	"DISABLE",
	"ENABLE",
	"QUOTA_DISABLE",
	"ROI2_DISABLE",
}

// NewQianchuanAdDetailGetV10DataOptStatusFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataOptStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataOptStatusFromValue(v string) (*QianchuanAdDetailGetV10DataOptStatus, error) {
	ev := QianchuanAdDetailGetV10DataOptStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataOptStatus: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataOptStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataOptStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataOptStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_opt_status value
func (v QianchuanAdDetailGetV10DataOptStatus) Ptr() *QianchuanAdDetailGetV10DataOptStatus {
	return &v
}
