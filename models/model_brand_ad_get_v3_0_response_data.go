/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandAdGetV30ResponseData
type BrandAdGetV30ResponseData struct {
	// 广告计划详情
	Ads []*BrandAdGetV30ResponseDataAdsInner `json:"ads,omitempty"`
	// 总计划数目
	TotalCount *int64 `json:"total_count,omitempty"`
}
