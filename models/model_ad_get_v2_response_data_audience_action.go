/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdGetV2ResponseDataAudienceAction
type AdGetV2ResponseDataAudienceAction struct {
	//
	ActionCategories []int64 `json:"action_categories,omitempty"`
	//
	ActionDays *int64 `json:"action_days,omitempty"`
	//
	ActionScene []*AdGetV2DataAudienceActionActionScene `json:"action_scene,omitempty"`
	//
	ActionWords []int64 `json:"action_words,omitempty"`
}
