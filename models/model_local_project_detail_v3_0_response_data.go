/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalProjectDetailV30ResponseData
type LocalProjectDetailV30ResponseData struct {
	AdType         *LocalProjectDetailV30DataAdType           `json:"ad_type,omitempty"`
	Audience       *LocalProjectDetailV30ResponseDataAudience `json:"audience,omitempty"`
	AutoUpdatePois *LocalProjectDetailV30DataAutoUpdatePois   `json:"auto_update_pois,omitempty"`
	// 抖音号id
	AwemeId *string `json:"aweme_id,omitempty"`
	// 出价，单位为分
	Bid     *int64                            `json:"bid,omitempty"`
	BidType *LocalProjectDetailV30DataBidType `json:"bid_type,omitempty"`
	// 项目预算，单位为分
	Budget          *int64                                    `json:"budget,omitempty"`
	BudgetMode      *LocalProjectDetailV30DataBudgetMode      `json:"budget_mode,omitempty"`
	DeliveryPoiMode *LocalProjectDetailV30DataDeliveryPoiMode `json:"delivery_poi_mode,omitempty"`
	// 投放结束时间
	EndTime        *string                                  `json:"end_time,omitempty"`
	ExternalAction *LocalProjectDetailV30DataExternalAction `json:"external_action,omitempty"`
	// 上调高峰日预算比例，单位为百分比
	HighBudgetRate *int64 `json:"high_budget_rate,omitempty"`
	// 是否设置了高峰日预算
	IsSetPeakBudget *bool `json:"is_set_peak_budget,omitempty"`
	// 广告账户id
	LocalAccountId     *int64                                       `json:"local_account_id,omitempty"`
	LocalDeliveryScene *LocalProjectDetailV30DataLocalDeliveryScene `json:"local_delivery_scene,omitempty"`
	MarketingGoal      *LocalProjectDetailV30DataMarketingGoal      `json:"marketing_goal,omitempty"`
	// 多门店id
	MultiPoiId *int64 `json:"multi_poi_id,omitempty"`
	// 项目名称
	Name *string `json:"name,omitempty"`
	// 高峰日-节假日
	PeakHolidays []*LocalProjectDetailV30DataPeakHolidays `json:"peak_holidays,omitempty"`
	// 高峰日-自然周
	PeakWeekDays []*LocalProjectDetailV30DataPeakWeekDays `json:"peak_week_days,omitempty"`
	// 单门店id
	PoiId *int64 `json:"poi_id,omitempty"`
	// 商品id
	ProductId *int64 `json:"product_id,omitempty"`
	// 项目id
	ProjectId *int64 `json:"project_id,omitempty"`
	// 固定投放时长，单位s
	ScheduleFixedSeconds *int64 `json:"schedule_fixed_seconds,omitempty"`
	// 投放时段
	ScheduleTime *string                                `json:"schedule_time,omitempty"`
	ScheduleType *LocalProjectDetailV30DataScheduleType `json:"schedule_type,omitempty"`
	// 投放起始时间
	StartTime *string `json:"start_time,omitempty"`
}