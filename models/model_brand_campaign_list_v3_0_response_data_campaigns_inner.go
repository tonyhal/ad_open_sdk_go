/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCampaignListV30ResponseDataCampaignsInner struct for BrandCampaignListV30ResponseDataCampaignsInner
type BrandCampaignListV30ResponseDataCampaignsInner struct {
	// 广告主ID
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	// 广告组ID
	CampaignId *int64 `json:"campaign_id,omitempty"`
	// 广告组名称
	CampaignName   *string                                          `json:"campaign_name,omitempty"`
	CampaignStatus *BrandCampaignListV30DataCampaignsCampaignStatus `json:"campaign_status,omitempty"`
	// 排期ID
	CartId *int64 `json:"cart_id,omitempty"`
	// 广告组创建时间
	CreateTime *string `json:"create_time,omitempty"`
	// 投放结束时间 \"2020-03-01\"
	EndTime *string `json:"end_time,omitempty"`
	// 合同ID
	MainContractId *string `json:"main_contract_id,omitempty"`
	// 投放开始时间 \"2020-03-01\"
	StartTime *string `json:"start_time,omitempty"`
	// 总金额 单位元
	TotalSellPrice *string `json:"total_sell_price,omitempty"`
}