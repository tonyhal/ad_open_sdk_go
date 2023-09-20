/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAbTestInfoGetV2ResponseDataTestReportsInnerDayStatsInner struct for ToolsAbTestInfoGetV2ResponseDataTestReportsInnerDayStatsInner
type ToolsAbTestInfoGetV2ResponseDataTestReportsInnerDayStatsInner struct {
	// 数据统计日期
	Date    *string                                                               `json:"date,omitempty"`
	Metrics *ToolsAbTestInfoGetV2ResponseDataTestReportsInnerDayStatsInnerMetrics `json:"metrics,omitempty"`
	// 实验对象ID
	ObjectId *int64 `json:"object_id,omitempty"`
}
