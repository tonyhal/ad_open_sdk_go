/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanOrientationPackageGetV10ResponseDataPageInfo
type QianchuanOrientationPackageGetV10ResponseDataPageInfo struct {
	//
	Page int32 `json:"page"`
	//
	PageSize int32 `json:"page_size"`
	//
	TotalNumber *int32 `json:"total_number,omitempty"`
	//
	TotalPage *int32 `json:"total_page,omitempty"`
}
