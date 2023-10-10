/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanProductAnalyseListV10TimeRange
type QianchuanProductAnalyseListV10TimeRange int64

// List of qianchuan_product_analyse_list_v1.0_time_range
const (
	Enum_15_QianchuanProductAnalyseListV10TimeRange QianchuanProductAnalyseListV10TimeRange = 15
	Enum_30_QianchuanProductAnalyseListV10TimeRange QianchuanProductAnalyseListV10TimeRange = 30
	Enum_7_QianchuanProductAnalyseListV10TimeRange  QianchuanProductAnalyseListV10TimeRange = 7
)

// All allowed values of QianchuanProductAnalyseListV10TimeRange enum
var AllowedQianchuanProductAnalyseListV10TimeRangeEnumValues = []QianchuanProductAnalyseListV10TimeRange{
	15,
	30,
	7,
}

// NewQianchuanProductAnalyseListV10TimeRangeFromValue returns a pointer to a valid QianchuanProductAnalyseListV10TimeRange
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanProductAnalyseListV10TimeRangeFromValue(v int64) (*QianchuanProductAnalyseListV10TimeRange, error) {
	ev := QianchuanProductAnalyseListV10TimeRange(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanProductAnalyseListV10TimeRange: valid values are %v", v, AllowedQianchuanProductAnalyseListV10TimeRangeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanProductAnalyseListV10TimeRange) IsValid() bool {
	for _, existing := range AllowedQianchuanProductAnalyseListV10TimeRangeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_product_analyse_list_v1.0_time_range value
func (v QianchuanProductAnalyseListV10TimeRange) Ptr() *QianchuanProductAnalyseListV10TimeRange {
	return &v
}
