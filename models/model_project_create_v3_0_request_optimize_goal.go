/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectCreateV30RequestOptimizeGoal
type ProjectCreateV30RequestOptimizeGoal struct {
	//
	AssetIds []int64 `json:"asset_ids,omitempty"`
	//
	ConvertId          *int64                                          `json:"convert_id,omitempty"`
	DeepExternalAction *ProjectCreateV30OptimizeGoalDeepExternalAction `json:"deep_external_action,omitempty"`
	ExternalAction     *ProjectCreateV30OptimizeGoalExternalAction     `json:"external_action,omitempty"`
	// 关键行为
	GameAddictionId *string `json:"game_addiction_id,omitempty"`
}
