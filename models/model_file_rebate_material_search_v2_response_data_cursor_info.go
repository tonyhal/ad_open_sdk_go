/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileRebateMaterialSearchV2ResponseDataCursorInfo 【分页方式①】
type FileRebateMaterialSearchV2ResponseDataCursorInfo struct {
	// 页码游标值
	Cursor *string `json:"cursor,omitempty"`
	// 是否有下一页
	HasMore *bool `json:"has_more,omitempty"`
}
