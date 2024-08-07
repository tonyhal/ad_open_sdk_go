/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AsyncTaskCreateV2TaskParamsOrderType
type AsyncTaskCreateV2TaskParamsOrderType string

// List of async_task_create_v2_task_params_order_type
const (
	ASC_AsyncTaskCreateV2TaskParamsOrderType  AsyncTaskCreateV2TaskParamsOrderType = "ASC"
	DESC_AsyncTaskCreateV2TaskParamsOrderType AsyncTaskCreateV2TaskParamsOrderType = "DESC"
)

// All allowed values of AsyncTaskCreateV2TaskParamsOrderType enum
var AllowedAsyncTaskCreateV2TaskParamsOrderTypeEnumValues = []AsyncTaskCreateV2TaskParamsOrderType{
	"ASC",
	"DESC",
}

// NewAsyncTaskCreateV2TaskParamsOrderTypeFromValue returns a pointer to a valid AsyncTaskCreateV2TaskParamsOrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAsyncTaskCreateV2TaskParamsOrderTypeFromValue(v string) (*AsyncTaskCreateV2TaskParamsOrderType, error) {
	ev := AsyncTaskCreateV2TaskParamsOrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AsyncTaskCreateV2TaskParamsOrderType: valid values are %v", v, AllowedAsyncTaskCreateV2TaskParamsOrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AsyncTaskCreateV2TaskParamsOrderType) IsValid() bool {
	for _, existing := range AllowedAsyncTaskCreateV2TaskParamsOrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to async_task_create_v2_task_params_order_type value
func (v AsyncTaskCreateV2TaskParamsOrderType) Ptr() *AsyncTaskCreateV2TaskParamsOrderType {
	return &v
}
