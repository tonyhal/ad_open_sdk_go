/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalReportPromotionGetV30Filtering
type LocalReportPromotionGetV30Filtering struct {
	CampaignType *LocalReportPromotionGetV30FilteringCampaignType `json:"campaign_type,omitempty"`
	// 优化目标
	ExternalAction     []*LocalReportPromotionGetV30FilteringExternalAction   `json:"external_action,omitempty"`
	LocalDeliveryScene *LocalReportPromotionGetV30FilteringLocalDeliveryScene `json:"local_delivery_scene,omitempty"`
	MarketingGoal      *LocalReportPromotionGetV30FilteringMarketingGoal      `json:"marketing_goal,omitempty"`
	// 广告ID，查询广告id。 注意：promotion_id是个维度条件，如果不传，查询的是本地推广告账户ID全部聚合后数据；如果仅传入字符串promotion_ids，查询的是维度明细是广告维度数据；如果传入数值，查询的是具体广告id数据
	PromotionIds []int64 `json:"promotion_ids,omitempty"`
}
