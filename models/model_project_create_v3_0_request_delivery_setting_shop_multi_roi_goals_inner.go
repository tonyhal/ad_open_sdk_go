/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectCreateV30RequestDeliverySettingShopMultiRoiGoalsInner struct for ProjectCreateV30RequestDeliverySettingShopMultiRoiGoalsInner
type ProjectCreateV30RequestDeliverySettingShopMultiRoiGoalsInner struct {
	//
	RoiGoal      *float64                                                      `json:"roi_goal,omitempty"`
	ShopPlatform *ProjectCreateV30DeliverySettingShopMultiRoiGoalsShopPlatform `json:"shop_platform,omitempty"`
}
