/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalProjectUpdateV30Request struct for LocalProjectUpdateV30Request
type LocalProjectUpdateV30Request struct {
	Audience *LocalProjectUpdateV30RequestAudience `json:"audience,omitempty"`
	//
	Bid *int64 `json:"bid,omitempty"`
	// 预算
	Budget *int64 `json:"budget,omitempty"`
	// 结束投放时间
	EndTime *string `json:"end_time,omitempty"`
	// 高峰日-预算上调比例（40表示上调40%）
	HighBudgetRate *int64 `json:"high_budget_rate,omitempty"`
	// 是否设置高峰日预算
	IsSetPeakBudget *bool `json:"is_set_peak_budget,omitempty"`
	//
	LocalAccountId int64 `json:"local_account_id"`
	// 项目名称
	Name *string `json:"name,omitempty"`
	// 高峰日-节假日
	PeakHolidays []*LocalProjectUpdateV30PeakHolidays `json:"peak_holidays,omitempty"`
	// 高峰日-自然周
	PeakWeekDays []*LocalProjectUpdateV30PeakWeekDays `json:"peak_week_days,omitempty"`
	// 项目ID
	ProjectId int64 `json:"project_id"`
	// 固定投放时长
	ScheduleFixedSeconds *int64 `json:"schedule_fixed_seconds,omitempty"`
	// 投放时段
	ScheduleTime *string                            `json:"schedule_time,omitempty"`
	ScheduleType *LocalProjectUpdateV30ScheduleType `json:"schedule_type,omitempty"`
}