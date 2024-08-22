/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandAdCancelDeleteV30Request struct for BrandAdCancelDeleteV30Request
type BrandAdCancelDeleteV30Request struct {
	// 计划ID列表
	AdIds []int64 `json:"ad_ids"`
	// 广告主ID
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	// 广告组ID
	CampaignId int64 `json:"campaign_id"`
}
