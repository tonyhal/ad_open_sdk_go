/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BudgetGroupUpdateV30Request struct for BudgetGroupUpdateV30Request
type BudgetGroupUpdateV30Request struct {
	// 广告主id
	AdvertiserId int64 `json:"advertiser_id"`
	// 更新后的预算组日预算
	BudgetGroupBudget *float64 `json:"budget_group_budget,omitempty"`
	// 需要更新的预算组id
	BudgetGroupId int64 `json:"budget_group_id"`
	// 更新后的预算组名称
	BudgetGroupName *string `json:"budget_group_name,omitempty"`
}
