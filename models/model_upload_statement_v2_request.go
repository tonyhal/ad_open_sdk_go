/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// UploadStatementV2Request struct for UploadStatementV2Request
type UploadStatementV2Request struct {
	// 代理商id
	AgentId int64 `json:"agent_id"`
	// 文件
	File *FormFileInfo `json:"file"`
}
