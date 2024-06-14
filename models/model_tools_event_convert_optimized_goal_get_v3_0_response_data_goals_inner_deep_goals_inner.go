/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsEventConvertOptimizedGoalGetV30ResponseDataGoalsInnerDeepGoalsInner struct for ToolsEventConvertOptimizedGoalGetV30ResponseDataGoalsInnerDeepGoalsInner
type ToolsEventConvertOptimizedGoalGetV30ResponseDataGoalsInnerDeepGoalsInner struct {
	//
	AssetTypes         []*ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsAssetTypes       `json:"asset_types,omitempty"`
	DeepExternalAction *ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsDeepExternalAction `json:"deep_external_action,omitempty"`
	//
	HistoryBack *bool `json:"history_back,omitempty"`
	//
	OptimizationName *string `json:"optimization_name,omitempty"`
	//
	TwentyFourHourBack *bool                                                            `json:"twenty_four_hour_back,omitempty"`
	ValueType          *ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueType `json:"value_type,omitempty"`
}
