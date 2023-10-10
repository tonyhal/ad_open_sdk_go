/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectListV30Filtering
type ProjectListV30Filtering struct {
	AdType           *ProjectListV30FilteringAdType           `json:"ad_type,omitempty"`
	AppPromotionType *ProjectListV30FilteringAppPromotionType `json:"app_promotion_type,omitempty"`
	DeliveryMode     *ProjectListV30FilteringDeliveryMode     `json:"delivery_mode,omitempty"`
	//
	Ids           []int64                               `json:"ids,omitempty"`
	InventoryType *ProjectListV30FilteringInventoryType `json:"inventory_type,omitempty"`
	LandingType   *ProjectListV30FilteringLandingType   `json:"landing_type,omitempty"`
	MarketingGoal *ProjectListV30FilteringMarketingGoal `json:"marketing_goal,omitempty"`
	//
	Name     *string                          `json:"name,omitempty"`
	Platform *ProjectListV30FilteringPlatform `json:"platform,omitempty"`
	Pricing  *ProjectListV30FilteringPricing  `json:"pricing,omitempty"`
	//
	ProjectCreateTime *string `json:"project_create_time,omitempty"`
	//
	ProjectModifyTime *string                              `json:"project_modify_time,omitempty"`
	Status            *ProjectListV30FilteringStatus       `json:"status,omitempty"`
	StatusFirst       *ProjectListV30FilteringStatusFirst  `json:"status_first,omitempty"`
	StatusSecond      *ProjectListV30FilteringStatusSecond `json:"status_second,omitempty"`
}
