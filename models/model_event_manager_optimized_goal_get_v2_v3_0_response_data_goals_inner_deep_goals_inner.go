/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EventManagerOptimizedGoalGetV2V30ResponseDataGoalsInnerDeepGoalsInner struct for EventManagerOptimizedGoalGetV2V30ResponseDataGoalsInnerDeepGoalsInner
type EventManagerOptimizedGoalGetV2V30ResponseDataGoalsInnerDeepGoalsInner struct {
	//
	AssetTypes         []*EventManagerOptimizedGoalGetV2V30DataGoalsDeepGoalsAssetTypes       `json:"asset_types,omitempty"`
	DeepExternalAction *EventManagerOptimizedGoalGetV2V30DataGoalsDeepGoalsDeepExternalAction `json:"deep_external_action,omitempty"`
	//
	HistoryBack *bool `json:"history_back,omitempty"`
	//
	OptimizationName *string `json:"optimization_name,omitempty"`
	//
	TwentyFourHourBack *bool `json:"twenty_four_hour_back,omitempty"`
}
