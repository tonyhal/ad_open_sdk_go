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

// QianchuanAdCreateV10AudienceActionDays
type QianchuanAdCreateV10AudienceActionDays int64

// List of qianchuan_ad_create_v1.0_audience_action_days
const (
	Enum_7_QianchuanAdCreateV10AudienceActionDays   QianchuanAdCreateV10AudienceActionDays = 7
	Enum_15_QianchuanAdCreateV10AudienceActionDays  QianchuanAdCreateV10AudienceActionDays = 15
	Enum_30_QianchuanAdCreateV10AudienceActionDays  QianchuanAdCreateV10AudienceActionDays = 30
	Enum_60_QianchuanAdCreateV10AudienceActionDays  QianchuanAdCreateV10AudienceActionDays = 60
	Enum_90_QianchuanAdCreateV10AudienceActionDays  QianchuanAdCreateV10AudienceActionDays = 90
	Enum_180_QianchuanAdCreateV10AudienceActionDays QianchuanAdCreateV10AudienceActionDays = 180
	Enum_365_QianchuanAdCreateV10AudienceActionDays QianchuanAdCreateV10AudienceActionDays = 365
)

// All allowed values of QianchuanAdCreateV10AudienceActionDays enum
var AllowedQianchuanAdCreateV10AudienceActionDaysEnumValues = []QianchuanAdCreateV10AudienceActionDays{
	7,
	15,
	30,
	60,
	90,
	180,
	365,
}

// NewQianchuanAdCreateV10AudienceActionDaysFromValue returns a pointer to a valid QianchuanAdCreateV10AudienceActionDays
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10AudienceActionDaysFromValue(v int64) (*QianchuanAdCreateV10AudienceActionDays, error) {
	ev := QianchuanAdCreateV10AudienceActionDays(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10AudienceActionDays: valid values are %v", v, AllowedQianchuanAdCreateV10AudienceActionDaysEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10AudienceActionDays) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10AudienceActionDaysEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_audience_action_days value
func (v QianchuanAdCreateV10AudienceActionDays) Ptr() *QianchuanAdCreateV10AudienceActionDays {
	return &v
}
