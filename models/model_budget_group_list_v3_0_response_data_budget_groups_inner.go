/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BudgetGroupListV30ResponseDataBudgetGroupsInner struct for BudgetGroupListV30ResponseDataBudgetGroupsInner
type BudgetGroupListV30ResponseDataBudgetGroupsInner struct {
	AdType *BudgetGroupListV30DataBudgetGroupsAdType `json:"ad_type,omitempty"`
	// 预算组预算
	BudgetGroupBudget *float64 `json:"budget_group_budget,omitempty"`
	// 预算组id
	BudgetGroupId *int64 `json:"budget_group_id,omitempty"`
	// 预算组名称
	BudgetGroupName        *string                                                   `json:"budget_group_name,omitempty"`
	BudgetGroupStatusFirst *BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusFirst `json:"budget_group_status_first,omitempty"`
	// 二级状态
	BudgetGroupStatusSecond []*BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusSecond `json:"budget_group_status_second,omitempty"`
	// 预算组消耗
	Cost         *float64                                        `json:"cost,omitempty"`
	DeliveryType *BudgetGroupListV30DataBudgetGroupsDeliveryType `json:"delivery_type,omitempty"`
	// 关联的（非删除）项目个数
	NumProjects *int64 `json:"num_projects,omitempty"`
}
