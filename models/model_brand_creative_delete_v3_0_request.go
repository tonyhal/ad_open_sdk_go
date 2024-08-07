/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCreativeDeleteV30Request struct for BrandCreativeDeleteV30Request
type BrandCreativeDeleteV30Request struct {
	// 计划ID
	AdId int64 `json:"ad_id"`
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 创意ID列表
	CreativeIds []int64 `json:"creative_ids"`
}
