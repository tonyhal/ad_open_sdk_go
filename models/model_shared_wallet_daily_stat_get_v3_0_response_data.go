/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SharedWalletDailyStatGetV30ResponseData
type SharedWalletDailyStatGetV30ResponseData struct {
	PageInfo *SharedWalletDailyStatGetV30ResponseDataPageInfo `json:"page_info,omitempty"`
	// 日流水列表信息
	Results []*SharedWalletDailyStatGetV30ResponseDataResultsInner `json:"results,omitempty"`
}
