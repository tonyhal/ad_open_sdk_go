/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementUploadTaskListV2Filtering
type ToolsAppManagementUploadTaskListV2Filtering struct {
	// 上传文件id，通过CreateUploadTask获得
	UploadIds []int64 `json:"upload_ids"`
}
