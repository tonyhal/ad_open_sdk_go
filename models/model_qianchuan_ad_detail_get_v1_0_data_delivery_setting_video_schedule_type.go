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

// QianchuanAdDetailGetV10DataDeliverySettingVideoScheduleType
type QianchuanAdDetailGetV10DataDeliverySettingVideoScheduleType string

// List of qianchuan_ad_detail_get_v1.0_data_delivery_setting_video_schedule_type
const (
	SCHEDULE_FROM_NOW_QianchuanAdDetailGetV10DataDeliverySettingVideoScheduleType  QianchuanAdDetailGetV10DataDeliverySettingVideoScheduleType = "SCHEDULE_FROM_NOW"
	SCHEDULE_START_END_QianchuanAdDetailGetV10DataDeliverySettingVideoScheduleType QianchuanAdDetailGetV10DataDeliverySettingVideoScheduleType = "SCHEDULE_START_END"
)

// All allowed values of QianchuanAdDetailGetV10DataDeliverySettingVideoScheduleType enum
var AllowedQianchuanAdDetailGetV10DataDeliverySettingVideoScheduleTypeEnumValues = []QianchuanAdDetailGetV10DataDeliverySettingVideoScheduleType{
	"SCHEDULE_FROM_NOW",
	"SCHEDULE_START_END",
}

// NewQianchuanAdDetailGetV10DataDeliverySettingVideoScheduleTypeFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataDeliverySettingVideoScheduleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataDeliverySettingVideoScheduleTypeFromValue(v string) (*QianchuanAdDetailGetV10DataDeliverySettingVideoScheduleType, error) {
	ev := QianchuanAdDetailGetV10DataDeliverySettingVideoScheduleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataDeliverySettingVideoScheduleType: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataDeliverySettingVideoScheduleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataDeliverySettingVideoScheduleType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataDeliverySettingVideoScheduleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_delivery_setting_video_schedule_type value
func (v QianchuanAdDetailGetV10DataDeliverySettingVideoScheduleType) Ptr() *QianchuanAdDetailGetV10DataDeliverySettingVideoScheduleType {
	return &v
}
