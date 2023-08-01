/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectCreateV30RequestDeliverySetting
type ProjectCreateV30RequestDeliverySetting struct {
	// cpc/cpm项目出价
	Bid      *float32                                 `json:"bid,omitempty"`
	BidSpeed *ProjectCreateV30DeliverySettingBidSpeed `json:"bid_speed,omitempty"`
	BidType  ProjectCreateV30DeliverySettingBidType   `json:"bid_type"`
	//
	Budget               *float32                                             `json:"budget,omitempty"`
	BudgetMode           ProjectCreateV30DeliverySettingBudgetMode            `json:"budget_mode"`
	BudgetOptimizeSwitch *ProjectCreateV30DeliverySettingBudgetOptimizeSwitch `json:"budget_optimize_switch,omitempty"`
	//
	CpaBid      *float32                                    `json:"cpa_bid,omitempty"`
	DeepBidType *ProjectCreateV30DeliverySettingDeepBidType `json:"deep_bid_type,omitempty"`
	//
	DeepCpabid *float32 `json:"deep_cpabid,omitempty"`
	//
	EndTime           *string                                           `json:"end_time,omitempty"`
	FilterNightSwitch *ProjectCreateV30DeliverySettingFilterNightSwitch `json:"filter_night_switch,omitempty"`
	Pricing           *ProjectCreateV30DeliverySettingPricing           `json:"pricing,omitempty"`
	ProjectCustom     *ProjectCreateV30DeliverySettingProjectCustom     `json:"project_custom,omitempty"`
	//
	RoiGoal *float32 `json:"roi_goal,omitempty"`
	//
	ScheduleTime *string                                     `json:"schedule_time,omitempty"`
	ScheduleType ProjectCreateV30DeliverySettingScheduleType `json:"schedule_type"`
	//
	StartTime *string `json:"start_time,omitempty"`
}