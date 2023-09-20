/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdQuotaGetV10ResponseDataQuotaSearchStageInfo 当前所在投放阶段信息
type QianchuanAdQuotaGetV10ResponseDataQuotaSearchStageInfo struct {
	// 投放初期截止日期
	EndDay *string `json:"end_day,omitempty"`
	// 是否在投放初期,0：否,1：是
	IsPrimary *int64 `json:"is_primary,omitempty"`
	// 投放初期起始日期
	StartDay *string `json:"start_day,omitempty"`
}
