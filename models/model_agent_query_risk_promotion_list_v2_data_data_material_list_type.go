/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentQueryRiskPromotionListV2DataDataMaterialListType
type AgentQueryRiskPromotionListV2DataDataMaterialListType string

// List of agent_query_risk_promotion_list_v2_data_data_material_list_type
const (
	IMAGE_AgentQueryRiskPromotionListV2DataDataMaterialListType AgentQueryRiskPromotionListV2DataDataMaterialListType = "IMAGE"
	VIDEO_AgentQueryRiskPromotionListV2DataDataMaterialListType AgentQueryRiskPromotionListV2DataDataMaterialListType = "VIDEO"
	SITE_AgentQueryRiskPromotionListV2DataDataMaterialListType  AgentQueryRiskPromotionListV2DataDataMaterialListType = "SITE"
)

// Ptr returns reference to agent_query_risk_promotion_list_v2_data_data_material_list_type value
func (v AgentQueryRiskPromotionListV2DataDataMaterialListType) Ptr() *AgentQueryRiskPromotionListV2DataDataMaterialListType {
	return &v
}
