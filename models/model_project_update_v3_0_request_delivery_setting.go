/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectUpdateV30RequestDeliverySetting
type ProjectUpdateV30RequestDeliverySetting struct {
	//
	Bid *float32 `json:"bid,omitempty"`
	//
	Budget     *float32                                   `json:"budget,omitempty"`
	BudgetMode *ProjectUpdateV30DeliverySettingBudgetMode `json:"budget_mode,omitempty"`
	//
	CpaBid *float32 `json:"cpa_bid,omitempty"`
	//
	DeepCpabid *float32 `json:"deep_cpabid,omitempty"`
	//
	EndTime           *string                                           `json:"end_time,omitempty"`
	FilterNightSwitch *ProjectUpdateV30DeliverySettingFilterNightSwitch `json:"filter_night_switch,omitempty"`
	//
	RoiGoal *float32 `json:"roi_goal,omitempty"`
	//
	ScheduleTime *string                                      `json:"schedule_time,omitempty"`
	ScheduleType *ProjectUpdateV30DeliverySettingScheduleType `json:"schedule_type,omitempty"`
}