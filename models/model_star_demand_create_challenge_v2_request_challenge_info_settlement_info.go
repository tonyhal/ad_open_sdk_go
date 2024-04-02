/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandCreateChallengeV2RequestChallengeInfoSettlementInfo 结算方式
type StarDemandCreateChallengeV2RequestChallengeInfoSettlementInfo struct {
	// 巨量引擎广告主ID 需要确保转化ID已在巨量引擎后台创建并激活
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	// Android转化ID 有效的转化ID
	AndroidConvertId *int64 `json:"android_convert_id,omitempty"`
	// 自动追加预算次数。总金额除以自动追加次数等于单次追加金额 auto_add_budget_trigger_ratio 为正时需提供，最低1，且保证单次追加预算金额不低于10000
	AutoAddBudgetTimes *int64 `json:"auto_add_budget_times,omitempty"`
	// 预算自动追加节点 当消耗到x%时，允许平台自动追加预算 非负整数。为0则不启用自动追加预算，最大100
	AutoAddBudgetTriggerRatio *int64 `json:"auto_add_budget_trigger_ratio,omitempty"`
	// 考核指标 5=视频有效播放量 4=激活总数 10=安装完成数量 31=组件点击数 32=APP唤起 201=客户自定义
	EvaluateType int64 `json:"evaluate_type"`
	// iOS转化ID 有效的转化ID
	IosConvertId *int64 `json:"ios_convert_id,omitempty"`
	// 自动追加预算总金额（单位元） auto_add_budget_trigger_ratio 为正时需提供，最低10000
	MaxBudgetAddAmount *int64 `json:"max_budget_add_amount,omitempty"`
	// 奖励规则 2000字内
	RewardRule *string `json:"reward_rule,omitempty"`
	// 奖励补充说明 200字内
	SupplementInfo *string `json:"supplement_info,omitempty"`
	// 单价（单位分） 根据考核指标的不同而不同： 为视频有效播放量时为千次曝光单价 为激活总数时为Android激活单价 为安装完成数量时为安装单价 为组件点击数时为单次点击单价 为APP唤起时为唤起单价 为客户自定义时为0
	UnitPrice int64 `json:"unit_price"`
}
