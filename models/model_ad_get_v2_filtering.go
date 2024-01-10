/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdGetV2Filtering
type AdGetV2Filtering struct {
	//
	AdCreateTime **string `json:"ad_create_time,omitempty"`
	//
	AdModifyTime **string `json:"ad_modify_time,omitempty"`
	//
	AdName *string `json:"ad_name,omitempty"`
	//
	CampaignId *int64 `json:"campaign_id,omitempty"`
	//
	Ids []int64 `json:"ids,omitempty"`
	//
	PricingList []string                `json:"pricing_list,omitempty"`
	Status      *AdGetV2FilteringStatus `json:"status,omitempty"`
}
