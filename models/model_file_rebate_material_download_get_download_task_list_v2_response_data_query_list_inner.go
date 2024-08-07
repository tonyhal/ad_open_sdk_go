/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileRebateMaterialDownloadGetDownloadTaskListV2ResponseDataQueryListInner struct for FileRebateMaterialDownloadGetDownloadTaskListV2ResponseDataQueryListInner
type FileRebateMaterialDownloadGetDownloadTaskListV2ResponseDataQueryListInner struct {
	// 查询ID
	QueryId *string `json:"query_id,omitempty"`
	// 查询状态 1- 初始化，2-运行中，3-成功，4-失败
	QueryStatus *int64 `json:"query_status,omitempty"`
	// 任务ID列表
	TaskList []*FileRebateMaterialDownloadGetDownloadTaskListV2ResponseDataQueryListInnerTaskListInner `json:"task_list,omitempty"`
}
