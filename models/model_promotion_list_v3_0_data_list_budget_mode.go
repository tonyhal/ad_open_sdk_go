/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionListV30DataListBudgetMode
type PromotionListV30DataListBudgetMode string

// List of promotion_list_v3.0_data_list_budget_mode
const (
	BUDGET_MODE_DAY_PromotionListV30DataListBudgetMode   PromotionListV30DataListBudgetMode = "BUDGET_MODE_DAY"
	BUDGET_MODE_TOTAL_PromotionListV30DataListBudgetMode PromotionListV30DataListBudgetMode = "BUDGET_MODE_TOTAL"
)

// All allowed values of PromotionListV30DataListBudgetMode enum
var AllowedPromotionListV30DataListBudgetModeEnumValues = []PromotionListV30DataListBudgetMode{
	"BUDGET_MODE_DAY",
	"BUDGET_MODE_TOTAL",
}

// NewPromotionListV30DataListBudgetModeFromValue returns a pointer to a valid PromotionListV30DataListBudgetMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListBudgetModeFromValue(v string) (*PromotionListV30DataListBudgetMode, error) {
	ev := PromotionListV30DataListBudgetMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListBudgetMode: valid values are %v", v, AllowedPromotionListV30DataListBudgetModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListBudgetMode) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListBudgetModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_budget_mode value
func (v PromotionListV30DataListBudgetMode) Ptr() *PromotionListV30DataListBudgetMode {
	return &v
}
