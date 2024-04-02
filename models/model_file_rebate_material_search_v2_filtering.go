/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileRebateMaterialSearchV2Filtering
type FileRebateMaterialSearchV2Filtering struct {
	// 广告主ID
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	// 客户名称
	CustomerName *string `json:"customer_name,omitempty"`
	// 【政策粒度】是否有效素材 允许值： 1是 2否 默认值：2-否
	IsValidVideoMaterial *int32 `json:"is_valid_video_material,omitempty"`
	// 素材创建日期范围结束日期 格式：yyyy-mm-dd
	MaterialCreateEndDate *string `json:"material_create_end_date,omitempty"`
	// 素材创建日期范围开始日期 格式：yyyy-mm-dd
	MaterialCreateStartDate *string `json:"material_create_start_date,omitempty"`
	// 当月最新一天素材是否在投 允许值： 1是 2否
	MaterialIsEffective *int32 `json:"material_is_effective,omitempty"`
	// 素材标签筛选项 允许值： INEFFICIENT_MATERIAL（低效素材） AD_HIGH_QUALITY_MATERIAL（AD 优质素材） FIRST_PUBLISH_MATERIAL（首发素材） - 如果不传标签筛选项，则返回广告主下所有素材 - 如果传入多个标签，取交集
	MaterialTags []string `json:"material_tags,omitempty"`
	// 运营标签 允许值： 1自运营 2走量 3收量
	OperatorTag *int32 `json:"operator_tag,omitempty"`
	// 【政策粒度】消耗范围区间上限 - 格式限制：整数
	PolicyCostMax *int64 `json:"policy_cost_max,omitempty"`
	// 【政策粒度】消耗范围区间下限 - 格式限制：整数
	PolicyCostMin *int64 `json:"policy_cost_min,omitempty"`
	// 一级结算行业
	RebateCalcFirstIndustryName *string `json:"rebate_calc_first_industry_name,omitempty"`
	// 政策类型 允许值： 1综代政策 2优代政策
	RebateCalcPolicyType *int32 `json:"rebate_calc_policy_type,omitempty"`
	// 二级结算行业
	RebateCalcSecondIndustryName *string `json:"rebate_calc_second_industry_name,omitempty"`
	// 结算行业统计类型 允许值： 1行业类目 2引流电商 4任务激励 5微信加粉
	RebateCalcSettlementStatsType *int32 `json:"rebate_calc_settlement_stats_type,omitempty"`
}
