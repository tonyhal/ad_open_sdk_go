/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueWechatPoolListV2ResponseData
type ClueWechatPoolListV2ResponseData struct {
	// 微信号列表
	Items []*ClueWechatPoolListV2ResponseDataItemsInner `json:"items,omitempty"`
	//  总数
	TotalCount *int64 `json:"total_count,omitempty"`
	// 总页数
	TotalPage *int64 `json:"total_page,omitempty"`
}
