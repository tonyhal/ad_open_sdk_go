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

// QianchuanUniAwemeAdCreateV10DeliverySettingLiveScheduleType
type QianchuanUniAwemeAdCreateV10DeliverySettingLiveScheduleType string

// List of qianchuan_uni_aweme_ad_create_v1.0_delivery_setting_live_schedule_type
const (
	SCHEDULE_FROM_NOW_QianchuanUniAwemeAdCreateV10DeliverySettingLiveScheduleType  QianchuanUniAwemeAdCreateV10DeliverySettingLiveScheduleType = "SCHEDULE_FROM_NOW"
	SCHEDULE_START_END_QianchuanUniAwemeAdCreateV10DeliverySettingLiveScheduleType QianchuanUniAwemeAdCreateV10DeliverySettingLiveScheduleType = "SCHEDULE_START_END"
)

// All allowed values of QianchuanUniAwemeAdCreateV10DeliverySettingLiveScheduleType enum
var AllowedQianchuanUniAwemeAdCreateV10DeliverySettingLiveScheduleTypeEnumValues = []QianchuanUniAwemeAdCreateV10DeliverySettingLiveScheduleType{
	"SCHEDULE_FROM_NOW",
	"SCHEDULE_START_END",
}

// NewQianchuanUniAwemeAdCreateV10DeliverySettingLiveScheduleTypeFromValue returns a pointer to a valid QianchuanUniAwemeAdCreateV10DeliverySettingLiveScheduleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanUniAwemeAdCreateV10DeliverySettingLiveScheduleTypeFromValue(v string) (*QianchuanUniAwemeAdCreateV10DeliverySettingLiveScheduleType, error) {
	ev := QianchuanUniAwemeAdCreateV10DeliverySettingLiveScheduleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanUniAwemeAdCreateV10DeliverySettingLiveScheduleType: valid values are %v", v, AllowedQianchuanUniAwemeAdCreateV10DeliverySettingLiveScheduleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanUniAwemeAdCreateV10DeliverySettingLiveScheduleType) IsValid() bool {
	for _, existing := range AllowedQianchuanUniAwemeAdCreateV10DeliverySettingLiveScheduleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_uni_aweme_ad_create_v1.0_delivery_setting_live_schedule_type value
func (v QianchuanUniAwemeAdCreateV10DeliverySettingLiveScheduleType) Ptr() *QianchuanUniAwemeAdCreateV10DeliverySettingLiveScheduleType {
	return &v
}