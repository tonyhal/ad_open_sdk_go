/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileRebateMaterialSearchV2ResponseDataMaterialsInner struct for FileRebateMaterialSearchV2ResponseDataMaterialsInner
type FileRebateMaterialSearchV2ResponseDataMaterialsInner struct {
	//
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	//
	AgentSaleName *string `json:"agent_sale_name,omitempty"`
	//
	CustomerName *string                                                           `json:"customer_name,omitempty"`
	MaterialInfo *FileRebateMaterialSearchV2ResponseDataMaterialsInnerMaterialInfo `json:"material_info,omitempty"`
	//
	OperatorTag *int32 `json:"operator_tag,omitempty"`
	//
	OptimizerName *string `json:"optimizer_name,omitempty"`
	//
	RebateCalcFirstIndustryName *string `json:"rebate_calc_first_industry_name,omitempty"`
	//
	RebateCalcPolicyType *int32 `json:"rebate_calc_policy_type,omitempty"`
	//
	RebateCalcSecondIndustryName *string `json:"rebate_calc_second_industry_name,omitempty"`
	//
	RebateCalcSettlementStatsType *int32 `json:"rebate_calc_settlement_stats_type,omitempty"`
	//
	SecondAdAgentId *int64 `json:"second_ad_agent_id,omitempty"`
	//
	UpdateDate *string `json:"update_date,omitempty"`
}
