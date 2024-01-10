/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsAssetTypes
type ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsAssetTypes string

// List of tools_event_convert_optimized_goal_get_v3.0_data_goals_deepGoals_asset_types
const (
	MINI_PROGRAME_ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsAssetTypes   ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsAssetTypes = "MINI_PROGRAME"
	TETRIS_EXTERNAL_ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsAssetTypes ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsAssetTypes = "TETRIS_EXTERNAL"
)

// All allowed values of ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsAssetTypes enum
var AllowedToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsAssetTypesEnumValues = []ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsAssetTypes{
	"MINI_PROGRAME",
	"TETRIS_EXTERNAL",
}

// NewToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsAssetTypesFromValue returns a pointer to a valid ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsAssetTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsAssetTypesFromValue(v string) (*ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsAssetTypes, error) {
	ev := ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsAssetTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsAssetTypes: valid values are %v", v, AllowedToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsAssetTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsAssetTypes) IsValid() bool {
	for _, existing := range AllowedToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsAssetTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_event_convert_optimized_goal_get_v3.0_data_goals_deepGoals_asset_types value
func (v ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsAssetTypes) Ptr() *ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsAssetTypes {
	return &v
}
