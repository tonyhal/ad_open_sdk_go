/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsTaskRaiseGetV2ResponseDataReportsInner struct for ToolsTaskRaiseGetV2ResponseDataReportsInner
type ToolsTaskRaiseGetV2ResponseDataReportsInner struct {
	//
	AllocatedBudget *float64                                  `json:"allocated_budget,omitempty"`
	BudgetMode      *ToolsTaskRaiseGetV2DataReportsBudgetMode `json:"budget_mode,omitempty"`
	//
	Duration *int64 `json:"duration,omitempty"`
	//
	IncreasedCost *float64 `json:"increased_cost,omitempty"`
	//
	Name      *string                                  `json:"name,omitempty"`
	RaiseMode *ToolsTaskRaiseGetV2DataReportsRaiseMode `json:"raise_mode,omitempty"`
	//
	ReportId *int64 `json:"report_id,omitempty"`
	//
	StartTime *string                               `json:"start_time,omitempty"`
	Status    *ToolsTaskRaiseGetV2DataReportsStatus `json:"status,omitempty"`
}
