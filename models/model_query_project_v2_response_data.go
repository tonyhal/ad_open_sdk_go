/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QueryProjectV2ResponseData
type QueryProjectV2ResponseData struct {
	// 页码游标值
	Cursor *int64 `json:"cursor,omitempty"`
	//
	ProjectInfoList []*QueryProjectV2ResponseDataProjectInfoListInner `json:"project_info_list,omitempty"`
	// 总量
	Total *int64 `json:"total,omitempty"`
}
