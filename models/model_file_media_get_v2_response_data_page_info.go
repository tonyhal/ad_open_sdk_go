/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileMediaGetV2ResponseDataPageInfo
type FileMediaGetV2ResponseDataPageInfo struct {
	//
	Count *int64 `json:"count,omitempty"`
	//
	Cursor *string `json:"cursor,omitempty"`
	//
	HasMore *bool `json:"has_more,omitempty"`
}
