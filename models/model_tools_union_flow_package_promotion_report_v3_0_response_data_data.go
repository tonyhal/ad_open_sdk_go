/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsUnionFlowPackagePromotionReportV30ResponseDataData
type ToolsUnionFlowPackagePromotionReportV30ResponseDataData struct {
	// 数据列表
	List     []*ToolsUnionFlowPackagePromotionReportV30ResponseDataDataListInner `json:"list,omitempty"`
	PageInfo *ToolsUnionFlowPackagePromotionReportV30ResponseDataDataPageInfo    `json:"page_info,omitempty"`
}
