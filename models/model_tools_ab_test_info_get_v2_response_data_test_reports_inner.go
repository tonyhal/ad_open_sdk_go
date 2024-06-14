/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAbTestInfoGetV2ResponseDataTestReportsInner struct for ToolsAbTestInfoGetV2ResponseDataTestReportsInner
type ToolsAbTestInfoGetV2ResponseDataTestReportsInner struct {
	// 实验对象分日统计数据
	DayStats []*ToolsAbTestInfoGetV2ResponseDataTestReportsInnerDayStatsInner `json:"day_stats,omitempty"`
	Metrics  *ToolsAbTestInfoGetV2ResponseDataTestReportsInnerMetrics         `json:"metrics,omitempty"`
	// 实验对象ID
	ObjectId *int64 `json:"object_id,omitempty"`
}
