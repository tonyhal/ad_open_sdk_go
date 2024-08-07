/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CampaignGetV2ResponseDataListInner struct for CampaignGetV2ResponseDataListInner
type CampaignGetV2ResponseDataListInner struct {
	//
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	//
	Budget                     *float64                                         `json:"budget,omitempty"`
	BudgetMode                 *CampaignGetV2DataListBudgetMode                 `json:"budget_mode,omitempty"`
	CampaignBudgetOptimization *CampaignGetV2DataListCampaignBudgetOptimization `json:"campaign_budget_optimization,omitempty"`
	//
	CampaignCreateTime *string `json:"campaign_create_time,omitempty"`
	//
	CampaignId *int64 `json:"campaign_id,omitempty"`
	//
	CampaignModifyTime *string                                  `json:"campaign_modify_time,omitempty"`
	CampaignType       *CampaignGetV2DataListCampaignType       `json:"campaign_type,omitempty"`
	DedicateType       *CampaignGetV2DataListDedicateType       `json:"dedicate_type,omitempty"`
	DeliveryMode       *CampaignGetV2DataListDeliveryMode       `json:"delivery_mode,omitempty"`
	DeliveryRelatedNum *CampaignGetV2DataListDeliveryRelatedNum `json:"delivery_related_num,omitempty"`
	//
	Id               *int64                                 `json:"id,omitempty"`
	LandingType      *CampaignGetV2DataListLandingType      `json:"landing_type,omitempty"`
	MarketingPurpose *CampaignGetV2DataListMarketingPurpose `json:"marketing_purpose,omitempty"`
	MarketingScene   *CampaignGetV2DataListMarketingScene   `json:"marketing_scene,omitempty"`
	//
	ModifyTime *string `json:"modify_time,omitempty"`
	//
	Name         *string                            `json:"name,omitempty"`
	SmartBidType *CampaignGetV2DataListSmartBidType `json:"smart_bid_type,omitempty"`
	Status       *CampaignGetV2DataListStatus       `json:"status,omitempty"`
}
