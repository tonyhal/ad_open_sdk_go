/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentQueryRiskPromotionListV2ResponseDataCursorInfo 分页信息
type AgentQueryRiskPromotionListV2ResponseDataCursorInfo struct {
	// 页码游标值
	Count *int64 `json:"count,omitempty"`
	// 页码游标值
	Cursor *int64 `json:"cursor,omitempty"`
	// 是否有下一页
	HasMore *bool `json:"has_more,omitempty"`
	// 总数
	TotalNumber *int64 `json:"total_number,omitempty"`
}
