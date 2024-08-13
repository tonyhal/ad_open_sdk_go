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

// QianchuanAdDetailGetV10DataDeliverySettingBudgetMode
type QianchuanAdDetailGetV10DataDeliverySettingBudgetMode string

// List of qianchuan_ad_detail_get_v1.0_data_delivery_setting_budget_mode
const (
	BUDGET_MODE_DAY_QianchuanAdDetailGetV10DataDeliverySettingBudgetMode             QianchuanAdDetailGetV10DataDeliverySettingBudgetMode = "BUDGET_MODE_DAY"
	BUDGET_MODE_INFINITE_QianchuanAdDetailGetV10DataDeliverySettingBudgetMode        QianchuanAdDetailGetV10DataDeliverySettingBudgetMode = "BUDGET_MODE_INFINITE"
	BUDGET_MODE_SEVEN_DAY_CYCLE_QianchuanAdDetailGetV10DataDeliverySettingBudgetMode QianchuanAdDetailGetV10DataDeliverySettingBudgetMode = "BUDGET_MODE_SEVEN_DAY_CYCLE"
	BUDGET_MODE_TOTAL_QianchuanAdDetailGetV10DataDeliverySettingBudgetMode           QianchuanAdDetailGetV10DataDeliverySettingBudgetMode = "BUDGET_MODE_TOTAL"
)

// All allowed values of QianchuanAdDetailGetV10DataDeliverySettingBudgetMode enum
var AllowedQianchuanAdDetailGetV10DataDeliverySettingBudgetModeEnumValues = []QianchuanAdDetailGetV10DataDeliverySettingBudgetMode{
	"BUDGET_MODE_DAY",
	"BUDGET_MODE_INFINITE",
	"BUDGET_MODE_SEVEN_DAY_CYCLE",
	"BUDGET_MODE_TOTAL",
}

// NewQianchuanAdDetailGetV10DataDeliverySettingBudgetModeFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataDeliverySettingBudgetMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataDeliverySettingBudgetModeFromValue(v string) (*QianchuanAdDetailGetV10DataDeliverySettingBudgetMode, error) {
	ev := QianchuanAdDetailGetV10DataDeliverySettingBudgetMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataDeliverySettingBudgetMode: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataDeliverySettingBudgetModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataDeliverySettingBudgetMode) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataDeliverySettingBudgetModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_delivery_setting_budget_mode value
func (v QianchuanAdDetailGetV10DataDeliverySettingBudgetMode) Ptr() *QianchuanAdDetailGetV10DataDeliverySettingBudgetMode {
	return &v
}
