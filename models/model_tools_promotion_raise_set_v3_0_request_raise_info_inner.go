/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPromotionRaiseSetV30RequestRaiseInfoInner struct for ToolsPromotionRaiseSetV30RequestRaiseInfoInner
type ToolsPromotionRaiseSetV30RequestRaiseInfoInner struct {
	AppointedTime *ToolsPromotionRaiseSetV30RequestRaiseInfoInnerAppointedTime `json:"appointed_time,omitempty"`
	// 是否立即生效，仅支持广告状态为“投放中”的广告，仅支持1个方案设置“立即生效”，传入True时不支持填写appointed_time
	IsEffectiveNow bool `json:"is_effective_now"`
	// 起量预算，单位：元，允许小数点后两位 起量预算需大于等于计划出价，小于等于计划预算
	RaiseBudget float64 `json:"raise_budget"`
}
