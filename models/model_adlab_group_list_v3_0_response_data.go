/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupListV30ResponseData
type AdlabGroupListV30ResponseData struct {
	// 项目列表
	Groups   []*AdlabGroupListV30ResponseDataGroupsInner `json:"groups,omitempty"`
	PageInfo *AdlabGroupListV30ResponseDataPageInfo      `json:"page_info,omitempty"`
}
