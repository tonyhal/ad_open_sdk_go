/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdUpdateBudgetV2ResponseData
type AdUpdateBudgetV2ResponseData struct {
	//
	AdIds []int64 `json:"ad_ids,omitempty"`
	//
	Errors []*AdUpdateBudgetV2ResponseDataErrorsInner `json:"errors,omitempty"`
}
