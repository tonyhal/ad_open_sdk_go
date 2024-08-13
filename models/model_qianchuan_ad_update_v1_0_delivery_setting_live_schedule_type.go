/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdUpdateV10DeliverySettingLiveScheduleType
type QianchuanAdUpdateV10DeliverySettingLiveScheduleType string

// List of qianchuan_ad_update_v1.0_delivery_setting_live_schedule_type
const (
	SCHEDULE_FROM_NOW_QianchuanAdUpdateV10DeliverySettingLiveScheduleType        QianchuanAdUpdateV10DeliverySettingLiveScheduleType = "SCHEDULE_FROM_NOW"
	SCHEDULE_START_END_QianchuanAdUpdateV10DeliverySettingLiveScheduleType       QianchuanAdUpdateV10DeliverySettingLiveScheduleType = "SCHEDULE_START_END"
	SCHEDULE_TIME_FIXEDRANGE_QianchuanAdUpdateV10DeliverySettingLiveScheduleType QianchuanAdUpdateV10DeliverySettingLiveScheduleType = "SCHEDULE_TIME_FIXEDRANGE"
)

// All allowed values of QianchuanAdUpdateV10DeliverySettingLiveScheduleType enum
var AllowedQianchuanAdUpdateV10DeliverySettingLiveScheduleTypeEnumValues = []QianchuanAdUpdateV10DeliverySettingLiveScheduleType{
	"SCHEDULE_FROM_NOW",
	"SCHEDULE_START_END",
	"SCHEDULE_TIME_FIXEDRANGE",
}

// NewQianchuanAdUpdateV10DeliverySettingLiveScheduleTypeFromValue returns a pointer to a valid QianchuanAdUpdateV10DeliverySettingLiveScheduleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdUpdateV10DeliverySettingLiveScheduleTypeFromValue(v string) (*QianchuanAdUpdateV10DeliverySettingLiveScheduleType, error) {
	ev := QianchuanAdUpdateV10DeliverySettingLiveScheduleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdUpdateV10DeliverySettingLiveScheduleType: valid values are %v", v, AllowedQianchuanAdUpdateV10DeliverySettingLiveScheduleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdUpdateV10DeliverySettingLiveScheduleType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdUpdateV10DeliverySettingLiveScheduleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_update_v1.0_delivery_setting_live_schedule_type value
func (v QianchuanAdUpdateV10DeliverySettingLiveScheduleType) Ptr() *QianchuanAdUpdateV10DeliverySettingLiveScheduleType {
	return &v
}
