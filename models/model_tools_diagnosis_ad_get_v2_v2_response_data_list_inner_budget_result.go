/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsDiagnosisAdGetV2V2ResponseDataListInnerBudgetResult
type ToolsDiagnosisAdGetV2V2ResponseDataListInnerBudgetResult struct {
	//
	AdDimBudget *float64 `json:"ad_dim_budget,omitempty"`
	//
	BudgetConclusionTag *string `json:"budget_conclusion_tag,omitempty"`
	//
	BudgetMode *int64 `json:"budget_mode,omitempty"`
	//
	BudgetSuggestion []string `json:"budget_suggestion,omitempty"`
	//
	StatCost *float64 `json:"stat_cost,omitempty"`
}
