/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalProjectCreateV30Request struct for LocalProjectCreateV30Request
type LocalProjectCreateV30Request struct {
	AdType         LocalProjectCreateV30AdType           `json:"ad_type"`
	Audience       *LocalProjectCreateV30RequestAudience `json:"audience,omitempty"`
	AutoUpdatePois *LocalProjectCreateV30AutoUpdatePois  `json:"auto_update_pois,omitempty"`
	// 用于推广直播间的抖音号，可通过【获取本地推创编可用抖音号】接口获取 填写说明： 当marketing_goal=`LIVE`时有效且必填
	AwemeId *string `json:"aweme_id,omitempty"`
	// 出价（单位：分）
	Bid     *int64                       `json:"bid,omitempty"`
	BidType LocalProjectCreateV30BidType `json:"bid_type"`
	// 预算（单位：分）
	Budget          int64                                 `json:"budget"`
	BudgetMode      LocalProjectCreateV30BudgetMode       `json:"budget_mode"`
	DeliveryGoal    *LocalProjectCreateV30DeliveryGoal    `json:"delivery_goal,omitempty"`
	DeliveryPoiMode *LocalProjectCreateV30DeliveryPoiMode `json:"delivery_poi_mode,omitempty"`
	// 结束投放时间 注意：当schedule_type=“START_TO_END”时支持填写该字段
	EndTime        *string                              `json:"end_time,omitempty"`
	ExternalAction *LocalProjectCreateV30ExternalAction `json:"external_action,omitempty"`
	// 高峰日预算上调比例（注意：该字段为百分比，例如：传“40”表示高峰日时预算上调“40%”）
	HighBudgetRate *int64 `json:"high_budget_rate,omitempty"`
	// 高峰日预算设置 该字段为false时：高峰日（自然周、节假日）、高峰日预算上调比例 均不可填值
	IsSetPeakBudget *bool `json:"is_set_peak_budget,omitempty"`
	// 本地推广告账户ID
	LocalAccountId     int64                                   `json:"local_account_id"`
	LocalDeliveryScene LocalProjectCreateV30LocalDeliveryScene `json:"local_delivery_scene"`
	MarketingGoal      LocalProjectCreateV30MarketingGoal      `json:"marketing_goal"`
	// 项目名称，长度是1-50个字（两个英文字符占1个字）
	Name string `json:"name"`
	// 高峰日-节假日
	PeakHolidays []*LocalProjectCreateV30PeakHolidays `json:"peak_holidays,omitempty"`
	// 高峰日-自然周
	PeakWeekDays []*LocalProjectCreateV30PeakWeekDays `json:"peak_week_days,omitempty"`
	// 推广商品ID，可通过【获取可投商品列表】接口查询获取 填写说明： 当marketing_goal=`VIDEO_AND_IMAGE` 且 delivery_goal=`PRODUCT` 时有效且必填
	ProductId *int64 `json:"product_id,omitempty"`
	// 推广门店ID列表  填写说明： 当delivery_goal=`POI`且 delivery_poi_mode=`CUSTOM` 投放指定门店时，有效且必填
	PromotionPoiIds []int64 `json:"promotion_poi_ids,omitempty"`
	// 直播固定时长（单位：秒） 交易场景下的设置规则： 仅在营销场景设置为直播时可用
	ScheduleFixedSeconds *int64 `json:"schedule_fixed_seconds,omitempty"`
	//
	ScheduleTime *string                           `json:"schedule_time,omitempty"`
	ScheduleType LocalProjectCreateV30ScheduleType `json:"schedule_type"`
	// 开始投放时间 注意：当schedule_type=“START_TO_END”时支持填写该字段
	StartTime *string `json:"start_time,omitempty"`
}
