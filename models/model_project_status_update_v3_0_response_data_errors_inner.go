/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectStatusUpdateV30ResponseDataErrorsInner struct for ProjectStatusUpdateV30ResponseDataErrorsInner
type ProjectStatusUpdateV30ResponseDataErrorsInner struct {
	// 错误信息
	ErrorMessage *string `json:"error_message,omitempty"`
	// 项目id
	ProjectId *int64 `json:"project_id,omitempty"`
}