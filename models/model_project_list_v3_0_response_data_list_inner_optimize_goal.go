/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectListV30ResponseDataListInnerOptimizeGoal
type ProjectListV30ResponseDataListInnerOptimizeGoal struct {
	//
	AssetIds []int64 `json:"asset_ids,omitempty"`
	//
	ConvertId          *int64                                                `json:"convert_id,omitempty"`
	DeepExternalAction *ProjectListV30DataListOptimizeGoalDeepExternalAction `json:"deep_external_action,omitempty"`
	ExternalAction     *ProjectListV30DataListOptimizeGoalExternalAction     `json:"external_action,omitempty"`
	//
	PaidSwitch *int64 `json:"paid_switch,omitempty"`
}
