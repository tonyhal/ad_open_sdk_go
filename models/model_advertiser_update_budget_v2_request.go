/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserUpdateBudgetV2Request struct for AdvertiserUpdateBudgetV2Request
type AdvertiserUpdateBudgetV2Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	//
	Budget     *float64                           `json:"budget,omitempty"`
	BudgetMode AdvertiserUpdateBudgetV2BudgetMode `json:"budget_mode"`
}
