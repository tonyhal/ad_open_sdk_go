/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectRoigoalUpdateV30RequestDataInner struct for ProjectRoigoalUpdateV30RequestDataInner
type ProjectRoigoalUpdateV30RequestDataInner struct {
	//
	ProjectId int64 `json:"project_id"`
	//
	RoiGoal *float64 `json:"roi_goal,omitempty"`
	//
	ShopMultiRoiGoals []*ProjectRoigoalUpdateV30RequestDataInnerShopMultiRoiGoalsInner `json:"shop_multi_roi_goals,omitempty"`
}
