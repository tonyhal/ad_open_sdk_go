/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupUpdateBudgetV30Response struct for AdlabGroupUpdateBudgetV30Response
type AdlabGroupUpdateBudgetV30Response struct {
	//
	Code *int64                                 `json:"code,omitempty"`
	Data *AdlabGroupUpdateBudgetV30ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
