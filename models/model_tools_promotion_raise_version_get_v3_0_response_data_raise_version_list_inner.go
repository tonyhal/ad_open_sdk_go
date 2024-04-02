/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPromotionRaiseVersionGetV30ResponseDataRaiseVersionListInner struct for ToolsPromotionRaiseVersionGetV30ResponseDataRaiseVersionListInner
type ToolsPromotionRaiseVersionGetV30ResponseDataRaiseVersionListInner struct {
	// 当前起量版本结束时间，格式：yyyy-mm-dd HH:MM
	EndTime *string `json:"end_time,omitempty"`
	// 起量版本号
	RaiseVersion *int64 `json:"raise_version,omitempty"`
	// 当前版本起量开始时间，格式：yyyy-mm-dd HH:MM
	StartTime *string `json:"start_time,omitempty"`
}
