/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanReportAdGetV10Filtering
type QianchuanReportAdGetV10Filtering struct {
	//
	AdIds          []int64                                         `json:"ad_ids,omitempty"`
	CampaignScene  *QianchuanReportAdGetV10FilteringCampaignScene  `json:"campaign_scene,omitempty"`
	MarketingGoal  QianchuanReportAdGetV10FilteringMarketingGoal   `json:"marketing_goal"`
	MarketingScene *QianchuanReportAdGetV10FilteringMarketingScene `json:"marketing_scene,omitempty"`
	OrderPlatform  *QianchuanReportAdGetV10FilteringOrderPlatform  `json:"order_platform,omitempty"`
	SmartBidType   *QianchuanReportAdGetV10FilteringSmartBidType   `json:"smart_bid_type,omitempty"`
	Status         *QianchuanReportAdGetV10FilteringStatus         `json:"status,omitempty"`
}
