/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectBudgetUpdateV30ResponseData
type ProjectBudgetUpdateV30ResponseData struct {
	//
	Errors []*ProjectBudgetUpdateV30ResponseDataErrorsInner `json:"errors,omitempty"`
	//
	ProjectIds []int64 `json:"project_ids,omitempty"`
}
