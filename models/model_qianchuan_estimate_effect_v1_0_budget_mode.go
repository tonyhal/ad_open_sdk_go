/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanEstimateEffectV10BudgetMode
type QianchuanEstimateEffectV10BudgetMode string

// List of qianchuan_estimate_effect_v1.0_budget_mode
const (
	BUDGET_MODE_DAY_QianchuanEstimateEffectV10BudgetMode   QianchuanEstimateEffectV10BudgetMode = "BUDGET_MODE_DAY"
	BUDGET_MODE_TOTAL_QianchuanEstimateEffectV10BudgetMode QianchuanEstimateEffectV10BudgetMode = "BUDGET_MODE_TOTAL"
)

// All allowed values of QianchuanEstimateEffectV10BudgetMode enum
var AllowedQianchuanEstimateEffectV10BudgetModeEnumValues = []QianchuanEstimateEffectV10BudgetMode{
	"BUDGET_MODE_DAY",
	"BUDGET_MODE_TOTAL",
}

// NewQianchuanEstimateEffectV10BudgetModeFromValue returns a pointer to a valid QianchuanEstimateEffectV10BudgetMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanEstimateEffectV10BudgetModeFromValue(v string) (*QianchuanEstimateEffectV10BudgetMode, error) {
	ev := QianchuanEstimateEffectV10BudgetMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanEstimateEffectV10BudgetMode: valid values are %v", v, AllowedQianchuanEstimateEffectV10BudgetModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanEstimateEffectV10BudgetMode) IsValid() bool {
	for _, existing := range AllowedQianchuanEstimateEffectV10BudgetModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_estimate_effect_v1.0_budget_mode value
func (v QianchuanEstimateEffectV10BudgetMode) Ptr() *QianchuanEstimateEffectV10BudgetMode {
	return &v
}
