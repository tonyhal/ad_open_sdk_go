/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdCreateV10DeliverySettingVideoScheduleType
type QianchuanAdCreateV10DeliverySettingVideoScheduleType string

// List of qianchuan_ad_create_v1.0_delivery_setting_video_schedule_type
const (
	SCHEDULE_FROM_NOW_QianchuanAdCreateV10DeliverySettingVideoScheduleType  QianchuanAdCreateV10DeliverySettingVideoScheduleType = "SCHEDULE_FROM_NOW"
	SCHEDULE_START_END_QianchuanAdCreateV10DeliverySettingVideoScheduleType QianchuanAdCreateV10DeliverySettingVideoScheduleType = "SCHEDULE_START_END"
)

// Ptr returns reference to qianchuan_ad_create_v1.0_delivery_setting_video_schedule_type value
func (v QianchuanAdCreateV10DeliverySettingVideoScheduleType) Ptr() *QianchuanAdCreateV10DeliverySettingVideoScheduleType {
	return &v
}
