/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserBudgetGetV2DataListBudgetMode
type AdvertiserBudgetGetV2DataListBudgetMode string

// List of advertiser_budget_get_v2_data_list_budget_mode
const (
	BUDGET_MODE_DAY_AdvertiserBudgetGetV2DataListBudgetMode      AdvertiserBudgetGetV2DataListBudgetMode = "BUDGET_MODE_DAY"
	BUDGET_MODE_INFINITE_AdvertiserBudgetGetV2DataListBudgetMode AdvertiserBudgetGetV2DataListBudgetMode = "BUDGET_MODE_INFINITE"
	BUDGET_MODE_TOTAL_AdvertiserBudgetGetV2DataListBudgetMode    AdvertiserBudgetGetV2DataListBudgetMode = "BUDGET_MODE_TOTAL"
)

// Ptr returns reference to advertiser_budget_get_v2_data_list_budget_mode value
func (v AdvertiserBudgetGetV2DataListBudgetMode) Ptr() *AdvertiserBudgetGetV2DataListBudgetMode {
	return &v
}
