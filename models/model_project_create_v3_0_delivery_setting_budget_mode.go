/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCreateV30DeliverySettingBudgetMode
type ProjectCreateV30DeliverySettingBudgetMode string

// List of project_create_v3.0_delivery_setting_budget_mode
const (
	BUDGET_MODE_DAY_ProjectCreateV30DeliverySettingBudgetMode      ProjectCreateV30DeliverySettingBudgetMode = "BUDGET_MODE_DAY"
	BUDGET_MODE_INFINITE_ProjectCreateV30DeliverySettingBudgetMode ProjectCreateV30DeliverySettingBudgetMode = "BUDGET_MODE_INFINITE"
)

// All allowed values of ProjectCreateV30DeliverySettingBudgetMode enum
var AllowedProjectCreateV30DeliverySettingBudgetModeEnumValues = []ProjectCreateV30DeliverySettingBudgetMode{
	"BUDGET_MODE_DAY",
	"BUDGET_MODE_INFINITE",
}

// NewProjectCreateV30DeliverySettingBudgetModeFromValue returns a pointer to a valid ProjectCreateV30DeliverySettingBudgetMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30DeliverySettingBudgetModeFromValue(v string) (*ProjectCreateV30DeliverySettingBudgetMode, error) {
	ev := ProjectCreateV30DeliverySettingBudgetMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30DeliverySettingBudgetMode: valid values are %v", v, AllowedProjectCreateV30DeliverySettingBudgetModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30DeliverySettingBudgetMode) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30DeliverySettingBudgetModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_delivery_setting_budget_mode value
func (v ProjectCreateV30DeliverySettingBudgetMode) Ptr() *ProjectCreateV30DeliverySettingBudgetMode {
	return &v
}
