/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EventManagerOptimizedGoalGetV2V30ResponseData
type EventManagerOptimizedGoalGetV2V30ResponseData struct {
	//
	AssetIds []int64 `json:"asset_ids,omitempty"`
	//
	Goals []*EventManagerOptimizedGoalGetV2V30ResponseDataGoalsInner `json:"goals,omitempty"`
}
