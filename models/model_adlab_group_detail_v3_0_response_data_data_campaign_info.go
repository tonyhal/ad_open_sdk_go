/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupDetailV30ResponseDataDataCampaignInfo 广告组维度信息
type AdlabGroupDetailV30ResponseDataDataCampaignInfo struct {
	// 项目预算
	Budget             float64                                                    `json:"budget"`
	BudgetMode         AdlabGroupDetailV30DataDataCampaignInfoBudgetMode          `json:"budget_mode"`
	DeliveryRelatedNum *AdlabGroupDetailV30DataDataCampaignInfoDeliveryRelatedNum `json:"delivery_related_num,omitempty"`
	LandingType        AdlabGroupDetailV30DataDataCampaignInfoLandingType         `json:"landing_type"`
	MarketingPurpose   AdlabGroupDetailV30DataDataCampaignInfoMarketingPurpose    `json:"marketing_purpose"`
	// 项目名称
	Name *string `json:"name,omitempty"`
}
