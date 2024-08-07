/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectCreateV30RequestDeliverySetting
type ProjectCreateV30RequestDeliverySetting struct {
	// cpc/cpm项目出价
	Bid      *float64                                 `json:"bid,omitempty"`
	BidSpeed *ProjectCreateV30DeliverySettingBidSpeed `json:"bid_speed,omitempty"`
	BidType  ProjectCreateV30DeliverySettingBidType   `json:"bid_type"`
	//
	Budget               *float64                                             `json:"budget,omitempty"`
	BudgetMode           ProjectCreateV30DeliverySettingBudgetMode            `json:"budget_mode"`
	BudgetOptimizeSwitch *ProjectCreateV30DeliverySettingBudgetOptimizeSwitch `json:"budget_optimize_switch,omitempty"`
	//
	CpaBid      *float64                                    `json:"cpa_bid,omitempty"`
	DeepBidType *ProjectCreateV30DeliverySettingDeepBidType `json:"deep_bid_type,omitempty"`
	//
	DeepCpabid *float64 `json:"deep_cpabid,omitempty"`
	//
	EndTime           *string                                           `json:"end_time,omitempty"`
	FilterNightSwitch *ProjectCreateV30DeliverySettingFilterNightSwitch `json:"filter_night_switch,omitempty"`
	//
	FirstRoiGoal  *float64                                      `json:"first_roi_goal,omitempty"`
	Pricing       *ProjectCreateV30DeliverySettingPricing       `json:"pricing,omitempty"`
	ProjectCustom *ProjectCreateV30DeliverySettingProjectCustom `json:"project_custom,omitempty"`
	//
	RoiGoal *float64 `json:"roi_goal,omitempty"`
	//
	ScheduleTime           *string                                                `json:"schedule_time,omitempty"`
	ScheduleType           *ProjectCreateV30DeliverySettingScheduleType           `json:"schedule_type,omitempty"`
	SearchContinueDelivery *ProjectCreateV30DeliverySettingSearchContinueDelivery `json:"search_continue_delivery,omitempty"`
	//
	ShopMultiRoiGoals []*ProjectCreateV30RequestDeliverySettingShopMultiRoiGoalsInner `json:"shop_multi_roi_goals,omitempty"`
	//
	StartTime *string `json:"start_time,omitempty"`
}
