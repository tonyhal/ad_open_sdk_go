/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPromotionDiagnosisSuggestionAcceptV30RequestToolsInnerParamsInnerParamValue
type ToolsPromotionDiagnosisSuggestionAcceptV30RequestToolsInnerParamsInnerParamValue struct {
	//
	BooleanParam *bool `json:"boolean_param,omitempty"`
	//
	ListParam []string `json:"list_param,omitempty"`
	//
	ObjectListParam []*map[string]string `json:"object_list_param,omitempty"`
	//
	StringParam *string `json:"string_param,omitempty"`
}
