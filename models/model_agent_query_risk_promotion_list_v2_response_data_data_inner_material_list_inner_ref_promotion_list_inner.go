/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentQueryRiskPromotionListV2ResponseDataDataInnerMaterialListInnerRefPromotionListInner struct for AgentQueryRiskPromotionListV2ResponseDataDataInnerMaterialListInnerRefPromotionListInner
type AgentQueryRiskPromotionListV2ResponseDataDataInnerMaterialListInnerRefPromotionListInner struct {
	// 同代理商账户下的其他关联广告ID所属的广告主账户ID
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	// 同代理商账户下的其他关联广告ID所属的代理商账户ID
	AgentId *int64 `json:"agent_id,omitempty"`
	// 同代理商账户下的其他关联广告ID
	PromotionId *int64 `json:"promotion_id,omitempty"`
}
