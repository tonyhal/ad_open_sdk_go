/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionScheduleTimeUpdateV30RequestDataInner struct for PromotionScheduleTimeUpdateV30RequestDataInner
type PromotionScheduleTimeUpdateV30RequestDataInner struct {
	// 广告ID
	PromotionId int64 `json:"promotion_id"`
	// 更新后的投放时段
	ScheduleTime string `json:"schedule_time"`
}
