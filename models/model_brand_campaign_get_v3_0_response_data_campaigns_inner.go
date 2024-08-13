/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCampaignGetV30ResponseDataCampaignsInner struct for BrandCampaignGetV30ResponseDataCampaignsInner
type BrandCampaignGetV30ResponseDataCampaignsInner struct {
	// 广告组总金额 该金额乘以10的5次方  避免出现小数
	AdsBudget *int64 `json:"ads_budget,omitempty"`
	// 广告主ID
	AdvertiserId       *int64                                                           `json:"advertiser_id,omitempty"`
	AuthenticationInfo *BrandCampaignGetV30ResponseDataCampaignsInnerAuthenticationInfo `json:"authentication_info,omitempty"`
	CampaignCategory   *BrandCampaignGetV30DataCampaignsCampaignCategory                `json:"campaign_category,omitempty"`
	// 广告组ID
	CampaignId *int64 `json:"campaign_id,omitempty"`
	// 广告组名称
	CampaignName   *string                                         `json:"campaign_name,omitempty"`
	CampaignStatus *BrandCampaignGetV30DataCampaignsCampaignStatus `json:"campaign_status,omitempty"`
	// 排期ID
	CartId       *int64                                        `json:"cart_id,omitempty"`
	ContractType *BrandCampaignGetV30DataCampaignsContractType `json:"contract_type,omitempty"`
	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`
	// 广告组投放截止时间
	EndTime          *string                                           `json:"end_time,omitempty"`
	MarketingPurpose *BrandCampaignGetV30DataCampaignsMarketingPurpose `json:"marketing_purpose,omitempty"`
	// 修改时间
	ModifyTime *string `json:"modify_time,omitempty"`
	// 广告组投放起始时间
	StartTime *string `json:"start_time,omitempty"`
}
