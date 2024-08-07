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

// ToolsTaskRaiseGetV2DataReportsBudgetMode
type ToolsTaskRaiseGetV2DataReportsBudgetMode string

// List of tools_task_raise_get_v2_data_reports_budget_mode
const (
	LIMIT_ToolsTaskRaiseGetV2DataReportsBudgetMode    ToolsTaskRaiseGetV2DataReportsBudgetMode = "LIMIT"
	NO_LIMIT_ToolsTaskRaiseGetV2DataReportsBudgetMode ToolsTaskRaiseGetV2DataReportsBudgetMode = "NO_LIMIT"
)

// All allowed values of ToolsTaskRaiseGetV2DataReportsBudgetMode enum
var AllowedToolsTaskRaiseGetV2DataReportsBudgetModeEnumValues = []ToolsTaskRaiseGetV2DataReportsBudgetMode{
	"LIMIT",
	"NO_LIMIT",
}

// NewToolsTaskRaiseGetV2DataReportsBudgetModeFromValue returns a pointer to a valid ToolsTaskRaiseGetV2DataReportsBudgetMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsTaskRaiseGetV2DataReportsBudgetModeFromValue(v string) (*ToolsTaskRaiseGetV2DataReportsBudgetMode, error) {
	ev := ToolsTaskRaiseGetV2DataReportsBudgetMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsTaskRaiseGetV2DataReportsBudgetMode: valid values are %v", v, AllowedToolsTaskRaiseGetV2DataReportsBudgetModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsTaskRaiseGetV2DataReportsBudgetMode) IsValid() bool {
	for _, existing := range AllowedToolsTaskRaiseGetV2DataReportsBudgetModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_task_raise_get_v2_data_reports_budget_mode value
func (v ToolsTaskRaiseGetV2DataReportsBudgetMode) Ptr() *ToolsTaskRaiseGetV2DataReportsBudgetMode {
	return &v
}
