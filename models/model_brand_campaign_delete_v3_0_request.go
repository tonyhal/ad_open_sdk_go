/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCampaignDeleteV30Request struct for BrandCampaignDeleteV30Request
type BrandCampaignDeleteV30Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 广告组ID
	CampaignId int64 `json:"campaign_id"`
}
