/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanProductAnalyseListV10OrderType
type QianchuanProductAnalyseListV10OrderType string

// List of qianchuan_product_analyse_list_v1.0_order_type
const (
	ASC_QianchuanProductAnalyseListV10OrderType  QianchuanProductAnalyseListV10OrderType = "ASC"
	DESC_QianchuanProductAnalyseListV10OrderType QianchuanProductAnalyseListV10OrderType = "DESC"
)

// All allowed values of QianchuanProductAnalyseListV10OrderType enum
var AllowedQianchuanProductAnalyseListV10OrderTypeEnumValues = []QianchuanProductAnalyseListV10OrderType{
	"ASC",
	"DESC",
}

// NewQianchuanProductAnalyseListV10OrderTypeFromValue returns a pointer to a valid QianchuanProductAnalyseListV10OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanProductAnalyseListV10OrderTypeFromValue(v string) (*QianchuanProductAnalyseListV10OrderType, error) {
	ev := QianchuanProductAnalyseListV10OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanProductAnalyseListV10OrderType: valid values are %v", v, AllowedQianchuanProductAnalyseListV10OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanProductAnalyseListV10OrderType) IsValid() bool {
	for _, existing := range AllowedQianchuanProductAnalyseListV10OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_product_analyse_list_v1.0_order_type value
func (v QianchuanProductAnalyseListV10OrderType) Ptr() *QianchuanProductAnalyseListV10OrderType {
	return &v
}
