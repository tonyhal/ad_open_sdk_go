/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AsyncTaskCreateV2RequestTaskParams
type AsyncTaskCreateV2RequestTaskParams struct {
	//
	EndDate string `json:"end_date"`
	//
	Fields    []string                                     `json:"fields,omitempty"`
	Filtering *AsyncTaskCreateV2RequestTaskParamsFiltering `json:"filtering,omitempty"`
	//
	GroupBy []string `json:"group_by"`
	//
	OrderField *string                               `json:"order_field,omitempty"`
	OrderType  *AsyncTaskCreateV2TaskParamsOrderType `json:"order_type,omitempty"`
	//
	StartDate string `json:"start_date"`
}
