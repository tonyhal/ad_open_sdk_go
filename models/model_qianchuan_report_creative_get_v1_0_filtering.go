/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanReportCreativeGetV10Filtering
type QianchuanReportCreativeGetV10Filtering struct {
	CampaignScene *QianchuanReportCreativeGetV10FilteringCampaignScene `json:"campaign_scene,omitempty"`
	//
	CreativeIds          []int64                                                     `json:"creative_ids,omitempty"`
	CreativeMaterialMode *QianchuanReportCreativeGetV10FilteringCreativeMaterialMode `json:"creative_material_mode,omitempty"`
	//
	ImageMode      []*QianchuanReportCreativeGetV10FilteringImageMode    `json:"image_mode,omitempty"`
	IsHighlight    *QianchuanReportCreativeGetV10FilteringIsHighlight    `json:"is_highlight,omitempty"`
	MarketingGoal  QianchuanReportCreativeGetV10FilteringMarketingGoal   `json:"marketing_goal"`
	MarketingScene *QianchuanReportCreativeGetV10FilteringMarketingScene `json:"marketing_scene,omitempty"`
	OrderPlatform  *QianchuanReportCreativeGetV10FilteringOrderPlatform  `json:"order_platform,omitempty"`
	SmartBidType   *QianchuanReportCreativeGetV10FilteringSmartBidType   `json:"smart_bid_type,omitempty"`
	Status         *QianchuanReportCreativeGetV10FilteringStatus         `json:"status,omitempty"`
}