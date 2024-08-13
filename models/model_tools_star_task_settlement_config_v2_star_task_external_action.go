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

// ToolsStarTaskSettlementConfigV2StarTaskExternalAction
type ToolsStarTaskSettlementConfigV2StarTaskExternalAction string

// List of tools_star_task_settlement_config_v2_star_task_external_action
const (
	AD_CONVERT_TYPE_ACTIVE_ToolsStarTaskSettlementConfigV2StarTaskExternalAction          ToolsStarTaskSettlementConfigV2StarTaskExternalAction = "AD_CONVERT_TYPE_ACTIVE"
	AD_CONVERT_TYPE_ACTIVE_REGISTER_ToolsStarTaskSettlementConfigV2StarTaskExternalAction ToolsStarTaskSettlementConfigV2StarTaskExternalAction = "AD_CONVERT_TYPE_ACTIVE_REGISTER"
	AD_CONVERT_TYPE_PAY_ToolsStarTaskSettlementConfigV2StarTaskExternalAction             ToolsStarTaskSettlementConfigV2StarTaskExternalAction = "AD_CONVERT_TYPE_PAY"
)

// All allowed values of ToolsStarTaskSettlementConfigV2StarTaskExternalAction enum
var AllowedToolsStarTaskSettlementConfigV2StarTaskExternalActionEnumValues = []ToolsStarTaskSettlementConfigV2StarTaskExternalAction{
	"AD_CONVERT_TYPE_ACTIVE",
	"AD_CONVERT_TYPE_ACTIVE_REGISTER",
	"AD_CONVERT_TYPE_PAY",
}

// NewToolsStarTaskSettlementConfigV2StarTaskExternalActionFromValue returns a pointer to a valid ToolsStarTaskSettlementConfigV2StarTaskExternalAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsStarTaskSettlementConfigV2StarTaskExternalActionFromValue(v string) (*ToolsStarTaskSettlementConfigV2StarTaskExternalAction, error) {
	ev := ToolsStarTaskSettlementConfigV2StarTaskExternalAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsStarTaskSettlementConfigV2StarTaskExternalAction: valid values are %v", v, AllowedToolsStarTaskSettlementConfigV2StarTaskExternalActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsStarTaskSettlementConfigV2StarTaskExternalAction) IsValid() bool {
	for _, existing := range AllowedToolsStarTaskSettlementConfigV2StarTaskExternalActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_star_task_settlement_config_v2_star_task_external_action value
func (v ToolsStarTaskSettlementConfigV2StarTaskExternalAction) Ptr() *ToolsStarTaskSettlementConfigV2StarTaskExternalAction {
	return &v
}
