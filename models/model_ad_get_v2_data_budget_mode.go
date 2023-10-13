/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdGetV2DataBudgetMode
type AdGetV2DataBudgetMode string

// List of ad_get_v2_data_budget_mode
const (
	BUDGET_MODE_DAY_AdGetV2DataBudgetMode      AdGetV2DataBudgetMode = "BUDGET_MODE_DAY"
	BUDGET_MODE_TOTAL_AdGetV2DataBudgetMode    AdGetV2DataBudgetMode = "BUDGET_MODE_TOTAL"
	BUDGET_MODE_INFINITE_AdGetV2DataBudgetMode AdGetV2DataBudgetMode = "BUDGET_MODE_INFINITE"
)

// All allowed values of AdGetV2DataBudgetMode enum
var AllowedAdGetV2DataBudgetModeEnumValues = []AdGetV2DataBudgetMode{
	"BUDGET_MODE_DAY",
	"BUDGET_MODE_TOTAL",
	"BUDGET_MODE_INFINITE",
}

// NewAdGetV2DataBudgetModeFromValue returns a pointer to a valid AdGetV2DataBudgetMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataBudgetModeFromValue(v string) (*AdGetV2DataBudgetMode, error) {
	ev := AdGetV2DataBudgetMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataBudgetMode: valid values are %v", v, AllowedAdGetV2DataBudgetModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataBudgetMode) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataBudgetModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_budget_mode value
func (v AdGetV2DataBudgetMode) Ptr() *AdGetV2DataBudgetMode {
	return &v
}
