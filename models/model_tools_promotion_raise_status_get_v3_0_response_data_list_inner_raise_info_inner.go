/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPromotionRaiseStatusGetV30ResponseDataListInnerRaiseInfoInner struct for ToolsPromotionRaiseStatusGetV30ResponseDataListInnerRaiseInfoInner
type ToolsPromotionRaiseStatusGetV30ResponseDataListInnerRaiseInfoInner struct {
	AppointTime *ToolsPromotionRaiseStatusGetV30ResponseDataListInnerRaiseInfoInnerAppointTime `json:"appoint_time,omitempty"`
	// 是否立即生效
	IsEffectiveNow *bool `json:"is_effective_now,omitempty"`
	// 起量预算
	RaiseBudget *float64 `json:"raise_budget,omitempty"`
}
