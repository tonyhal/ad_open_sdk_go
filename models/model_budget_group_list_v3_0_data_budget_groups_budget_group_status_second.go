/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusSecond
type BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusSecond string

// List of budget_group_list_v3.0_data_budget_groups_budget_group_status_second
const (
	ACCOUNT_EXCEEDED_BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusSecond BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusSecond = "ACCOUNT_EXCEEDED"
	GROUP_EXCEEDED_BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusSecond   BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusSecond = "GROUP_EXCEEDED"
)

// Ptr returns reference to budget_group_list_v3.0_data_budget_groups_budget_group_status_second value
func (v BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusSecond) Ptr() *BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusSecond {
	return &v
}
