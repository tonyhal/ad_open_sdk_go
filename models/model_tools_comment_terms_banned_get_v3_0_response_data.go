/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsCommentTermsBannedGetV30ResponseData
type ToolsCommentTermsBannedGetV30ResponseData struct {
	PageInfo ToolsCommentTermsBannedGetV30ResponseDataPageInfo `json:"page_info"`
	// 屏蔽词列表
	Terms []string `json:"terms"`
}
