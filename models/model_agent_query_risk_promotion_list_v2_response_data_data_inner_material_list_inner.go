/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentQueryRiskPromotionListV2ResponseDataDataInnerMaterialListInner struct for AgentQueryRiskPromotionListV2ResponseDataDataInnerMaterialListInner
type AgentQueryRiskPromotionListV2ResponseDataDataInnerMaterialListInner struct {
	// 素材ID（落地页站点ID）
	Id *int64 `json:"id,omitempty"`
	// 同代理商账户下的其他关联广告ID
	RefPromotionIds []int64 `json:"ref_promotion_ids,omitempty"`
	// 素材违规原因，比如：[\"违规内容1\", \"违规内容2\"]
	RiskContent []string                                               `json:"risk_content,omitempty"`
	Type        *AgentQueryRiskPromotionListV2DataDataMaterialListType `json:"type,omitempty"`
}
