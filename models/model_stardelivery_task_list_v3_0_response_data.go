/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StardeliveryTaskListV30ResponseData
type StardeliveryTaskListV30ResponseData struct {
	// 任务列表
	List     []*StardeliveryTaskListV30ResponseDataListInner `json:"list,omitempty"`
	PageInfo *StardeliveryTaskListV30ResponseDataPageInfo    `json:"page_info,omitempty"`
}
