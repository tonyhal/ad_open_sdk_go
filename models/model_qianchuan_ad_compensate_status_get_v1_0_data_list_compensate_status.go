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

// QianchuanAdCompensateStatusGetV10DataListCompensateStatus
type QianchuanAdCompensateStatusGetV10DataListCompensateStatus string

// List of qianchuan_ad_compensate_status_get_v1.0_data_list_compensate_status
const (
	IN_EFFECT_QianchuanAdCompensateStatusGetV10DataListCompensateStatus  QianchuanAdCompensateStatusGetV10DataListCompensateStatus = "IN_EFFECT"
	INVALID_QianchuanAdCompensateStatusGetV10DataListCompensateStatus    QianchuanAdCompensateStatusGetV10DataListCompensateStatus = "INVALID"
	CONFIRMING_QianchuanAdCompensateStatusGetV10DataListCompensateStatus QianchuanAdCompensateStatusGetV10DataListCompensateStatus = "CONFIRMING"
	PAID_QianchuanAdCompensateStatusGetV10DataListCompensateStatus       QianchuanAdCompensateStatusGetV10DataListCompensateStatus = "PAID"
	ENDED_QianchuanAdCompensateStatusGetV10DataListCompensateStatus      QianchuanAdCompensateStatusGetV10DataListCompensateStatus = "ENDED"
	DEFAULT_QianchuanAdCompensateStatusGetV10DataListCompensateStatus    QianchuanAdCompensateStatusGetV10DataListCompensateStatus = "DEFAULT"
)

// All allowed values of QianchuanAdCompensateStatusGetV10DataListCompensateStatus enum
var AllowedQianchuanAdCompensateStatusGetV10DataListCompensateStatusEnumValues = []QianchuanAdCompensateStatusGetV10DataListCompensateStatus{
	"IN_EFFECT",
	"INVALID",
	"CONFIRMING",
	"PAID",
	"ENDED",
	"DEFAULT",
}

// NewQianchuanAdCompensateStatusGetV10DataListCompensateStatusFromValue returns a pointer to a valid QianchuanAdCompensateStatusGetV10DataListCompensateStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCompensateStatusGetV10DataListCompensateStatusFromValue(v string) (*QianchuanAdCompensateStatusGetV10DataListCompensateStatus, error) {
	ev := QianchuanAdCompensateStatusGetV10DataListCompensateStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCompensateStatusGetV10DataListCompensateStatus: valid values are %v", v, AllowedQianchuanAdCompensateStatusGetV10DataListCompensateStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCompensateStatusGetV10DataListCompensateStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCompensateStatusGetV10DataListCompensateStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_compensate_status_get_v1.0_data_list_compensate_status value
func (v QianchuanAdCompensateStatusGetV10DataListCompensateStatus) Ptr() *QianchuanAdCompensateStatusGetV10DataListCompensateStatus {
	return &v
}
