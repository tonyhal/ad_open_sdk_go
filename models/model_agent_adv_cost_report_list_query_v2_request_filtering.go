/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentAdvCostReportListQueryV2RequestFiltering 过滤条件
type AgentAdvCostReportListQueryV2RequestFiltering struct {
	// 广告主 id
	AdvertiserIds []int64 `json:"advertiser_ids,omitempty"`
	// 代理商客户id
	AgentCustomerId *int64                                         `json:"agent_customer_id,omitempty"`
	AppName         *AgentAdvCostReportListQueryV2FilteringAppName `json:"app_name,omitempty"`
	// 广告主所属公司名称，若选填该字段，限制最小长度为1，最大长度为223。支持模糊查询。
	CompanyName   *string                                              `json:"company_name,omitempty"`
	CostSource    *AgentAdvCostReportListQueryV2FilteringCostSource    `json:"cost_source,omitempty"`
	EcommerceType *AgentAdvCostReportListQueryV2FilteringEcommerceType `json:"ecommerce_type,omitempty"`
	// 一级行业名称。可从【获取行业列表】接口获取。
	FirstIndustry   *string                                                `json:"first_industry,omitempty"`
	PricingCategory *AgentAdvCostReportListQueryV2FilteringPricingCategory `json:"pricing_category,omitempty"`
	// 项目名称，若选填该字段，限制最小长度为1，最大长度为223。支持模糊查询
	ProjectName *string `json:"project_name,omitempty"`
	// 项目编号
	ProjectSerial *string                                              `json:"project_serial,omitempty"`
	PromotionType *AgentAdvCostReportListQueryV2FilteringPromotionType `json:"promotion_type,omitempty"`
	// 代理商子账户id
	SecondAdAgentId *int64 `json:"second_ad_agent_id,omitempty"`
	// 二级行业名称。可从【获取行业列表】接口获取。
	SecondIndustry *string `json:"second_industry,omitempty"`
}
