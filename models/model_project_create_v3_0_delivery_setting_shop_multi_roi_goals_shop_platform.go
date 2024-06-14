/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCreateV30DeliverySettingShopMultiRoiGoalsShopPlatform
type ProjectCreateV30DeliverySettingShopMultiRoiGoalsShopPlatform string

// List of project_create_v3.0_delivery_setting_shop_multi_roi_goals_shop_platform
const (
	JD_ProjectCreateV30DeliverySettingShopMultiRoiGoalsShopPlatform    ProjectCreateV30DeliverySettingShopMultiRoiGoalsShopPlatform = "JD"
	OTHER_ProjectCreateV30DeliverySettingShopMultiRoiGoalsShopPlatform ProjectCreateV30DeliverySettingShopMultiRoiGoalsShopPlatform = "OTHER"
	PDD_ProjectCreateV30DeliverySettingShopMultiRoiGoalsShopPlatform   ProjectCreateV30DeliverySettingShopMultiRoiGoalsShopPlatform = "PDD"
	TB_ProjectCreateV30DeliverySettingShopMultiRoiGoalsShopPlatform    ProjectCreateV30DeliverySettingShopMultiRoiGoalsShopPlatform = "TB"
)

// All allowed values of ProjectCreateV30DeliverySettingShopMultiRoiGoalsShopPlatform enum
var AllowedProjectCreateV30DeliverySettingShopMultiRoiGoalsShopPlatformEnumValues = []ProjectCreateV30DeliverySettingShopMultiRoiGoalsShopPlatform{
	"JD",
	"OTHER",
	"PDD",
	"TB",
}

// NewProjectCreateV30DeliverySettingShopMultiRoiGoalsShopPlatformFromValue returns a pointer to a valid ProjectCreateV30DeliverySettingShopMultiRoiGoalsShopPlatform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30DeliverySettingShopMultiRoiGoalsShopPlatformFromValue(v string) (*ProjectCreateV30DeliverySettingShopMultiRoiGoalsShopPlatform, error) {
	ev := ProjectCreateV30DeliverySettingShopMultiRoiGoalsShopPlatform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30DeliverySettingShopMultiRoiGoalsShopPlatform: valid values are %v", v, AllowedProjectCreateV30DeliverySettingShopMultiRoiGoalsShopPlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30DeliverySettingShopMultiRoiGoalsShopPlatform) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30DeliverySettingShopMultiRoiGoalsShopPlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_delivery_setting_shop_multi_roi_goals_shop_platform value
func (v ProjectCreateV30DeliverySettingShopMultiRoiGoalsShopPlatform) Ptr() *ProjectCreateV30DeliverySettingShopMultiRoiGoalsShopPlatform {
	return &v
}
