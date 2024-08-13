/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// EnterpriseFlowCategoryGetV10OrderType
type EnterpriseFlowCategoryGetV10OrderType string

// List of enterprise_flow_category_get_v1.0_order_type
const (
	ASC_EnterpriseFlowCategoryGetV10OrderType  EnterpriseFlowCategoryGetV10OrderType = "ASC"
	DESC_EnterpriseFlowCategoryGetV10OrderType EnterpriseFlowCategoryGetV10OrderType = "DESC"
)

// All allowed values of EnterpriseFlowCategoryGetV10OrderType enum
var AllowedEnterpriseFlowCategoryGetV10OrderTypeEnumValues = []EnterpriseFlowCategoryGetV10OrderType{
	"ASC",
	"DESC",
}

// NewEnterpriseFlowCategoryGetV10OrderTypeFromValue returns a pointer to a valid EnterpriseFlowCategoryGetV10OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnterpriseFlowCategoryGetV10OrderTypeFromValue(v string) (*EnterpriseFlowCategoryGetV10OrderType, error) {
	ev := EnterpriseFlowCategoryGetV10OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnterpriseFlowCategoryGetV10OrderType: valid values are %v", v, AllowedEnterpriseFlowCategoryGetV10OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnterpriseFlowCategoryGetV10OrderType) IsValid() bool {
	for _, existing := range AllowedEnterpriseFlowCategoryGetV10OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to enterprise_flow_category_get_v1.0_order_type value
func (v EnterpriseFlowCategoryGetV10OrderType) Ptr() *EnterpriseFlowCategoryGetV10OrderType {
	return &v
}
