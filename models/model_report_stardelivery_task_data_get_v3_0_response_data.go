/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportStardeliveryTaskDataGetV30ResponseData
type ReportStardeliveryTaskDataGetV30ResponseData struct {
	// 任务列表
	List     []*ReportStardeliveryTaskDataGetV30ResponseDataListInner `json:"list,omitempty"`
	PageInfo *ReportStardeliveryTaskDataGetV30ResponseDataPageInfo    `json:"page_info,omitempty"`
}
