/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAudiencePushV10Request struct for QianchuanAudiencePushV10Request
type QianchuanAudiencePushV10Request struct {
	AccountType QianchuanAudiencePushV10AccountType `json:"account_type"`
	//
	AdvIds []int64 `json:"adv_ids,omitempty"`
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	AudienceId int64 `json:"audience_id"`
	//
	IsForBrand *int32 `json:"is_for_brand,omitempty"`
}
