/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAccountBudgetGetV10DataBudgetMode
type QianchuanAccountBudgetGetV10DataBudgetMode string

// List of qianchuan_account_budget_get_v1.0_data_budget_mode
const (
	INFINITE_QianchuanAccountBudgetGetV10DataBudgetMode  QianchuanAccountBudgetGetV10DataBudgetMode = "INFINITE"
	SPECIFIED_QianchuanAccountBudgetGetV10DataBudgetMode QianchuanAccountBudgetGetV10DataBudgetMode = "SPECIFIED"
)

// All allowed values of QianchuanAccountBudgetGetV10DataBudgetMode enum
var AllowedQianchuanAccountBudgetGetV10DataBudgetModeEnumValues = []QianchuanAccountBudgetGetV10DataBudgetMode{
	"INFINITE",
	"SPECIFIED",
}

// NewQianchuanAccountBudgetGetV10DataBudgetModeFromValue returns a pointer to a valid QianchuanAccountBudgetGetV10DataBudgetMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAccountBudgetGetV10DataBudgetModeFromValue(v string) (*QianchuanAccountBudgetGetV10DataBudgetMode, error) {
	ev := QianchuanAccountBudgetGetV10DataBudgetMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAccountBudgetGetV10DataBudgetMode: valid values are %v", v, AllowedQianchuanAccountBudgetGetV10DataBudgetModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAccountBudgetGetV10DataBudgetMode) IsValid() bool {
	for _, existing := range AllowedQianchuanAccountBudgetGetV10DataBudgetModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_account_budget_get_v1.0_data_budget_mode value
func (v QianchuanAccountBudgetGetV10DataBudgetMode) Ptr() *QianchuanAccountBudgetGetV10DataBudgetMode {
	return &v
}
