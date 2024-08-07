/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsIndustryGetV2ResponseDataListInner struct for ToolsIndustryGetV2ResponseDataListInner
type ToolsIndustryGetV2ResponseDataListInner struct {
	//
	Description *string `json:"description,omitempty"`
	//
	FirstIndustryId *int64 `json:"first_industry_id,omitempty"`
	//
	FirstIndustryName *string `json:"first_industry_name,omitempty"`
	//
	IndustryId *int64 `json:"industry_id,omitempty"`
	//
	IndustryName *string `json:"industry_name,omitempty"`
	//
	Level *int64 `json:"level,omitempty"`
	//
	SecondIndustryId *int64 `json:"second_industry_id,omitempty"`
	//
	SecondIndustryName *string `json:"second_industry_name,omitempty"`
	//
	ThirdIndustryId *int64 `json:"third_industry_id,omitempty"`
	//
	ThirdIndustryName *string `json:"third_industry_name,omitempty"`
}
