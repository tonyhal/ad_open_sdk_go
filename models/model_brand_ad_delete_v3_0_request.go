/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandAdDeleteV30Request struct for BrandAdDeleteV30Request
type BrandAdDeleteV30Request struct {
	// 广告计划ID列表
	AdIds []int64 `json:"ad_ids"`
	// 广告主ID
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	// 广告组ID
	CampaignId int64 `json:"campaign_id"`
}
