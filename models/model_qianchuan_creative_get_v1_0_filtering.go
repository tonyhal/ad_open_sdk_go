/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanCreativeGetV10Filtering
type QianchuanCreativeGetV10Filtering struct {
	//
	AdIds []int64 `json:"ad_ids,omitempty"`
	//
	CampaignId *int64 `json:"campaign_id,omitempty"`
	//
	CreativeCreateEndDate *string `json:"creative_create_end_date,omitempty"`
	//
	CreativeCreateStartDate *string `json:"creative_create_start_date,omitempty"`
	//
	CreativeId           *int64                                                `json:"creative_id,omitempty"`
	CreativeMaterialMode *QianchuanCreativeGetV10FilteringCreativeMaterialMode `json:"creative_material_mode,omitempty"`
	//
	CreativeModifyTime *string                                         `json:"creative_modify_time,omitempty"`
	MarketingGoal      QianchuanCreativeGetV10FilteringMarketingGoal   `json:"marketing_goal"`
	MarketingScene     *QianchuanCreativeGetV10FilteringMarketingScene `json:"marketing_scene,omitempty"`
	Status             *QianchuanCreativeGetV10FilteringStatus         `json:"status,omitempty"`
}
