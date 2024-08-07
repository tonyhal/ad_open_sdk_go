/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsStarTaskSettlementConfigV2ResponseData
type ToolsStarTaskSettlementConfigV2ResponseData struct {
	// 任务最小可设置基础素材费单价（元），每条视频最低需要支付的价格,影响达人接单积极性 - 不同行业+素材类目组合，最小可设置素材费单价不同
	MinStarMaterialBid *float64 `json:"min_star_material_bid,omitempty"`
	// 任务最小可设置预算（元） - 最小可设置预算为固定值
	MinStarTaskBudget *float64 `json:"min_star_task_budget,omitempty"`
	// 素材分成费比例（单位是千分之一分，比如接口返回是20= 千分之20，实际是2%） - 不同的行业+素材类目组合，分成比例可能不同 - 限时免费期间，分成比例为0，后续变更请关注巨量运营/销售同学的通知
	StarAdCostDivideRatio *int64 `json:"star_ad_cost_divide_ratio,omitempty"`
	// 任务建议出价金额（元） - 不同行业和优化目标组合下，建议出价金额不同
	SuggestionStarTaskBid *float64 `json:"suggestion_star_task_bid,omitempty"`
}
