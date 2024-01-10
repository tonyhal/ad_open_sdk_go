/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanReportSearchWordGetV10FilteringRangeCalculation
type QianchuanReportSearchWordGetV10FilteringRangeCalculation string

// List of qianchuan_report_search_word_get_v1.0_filtering_range_calculation
const (
	LOWER_QianchuanReportSearchWordGetV10FilteringRangeCalculation QianchuanReportSearchWordGetV10FilteringRangeCalculation = "LOWER"
	UPPER_QianchuanReportSearchWordGetV10FilteringRangeCalculation QianchuanReportSearchWordGetV10FilteringRangeCalculation = "UPPER"
)

// All allowed values of QianchuanReportSearchWordGetV10FilteringRangeCalculation enum
var AllowedQianchuanReportSearchWordGetV10FilteringRangeCalculationEnumValues = []QianchuanReportSearchWordGetV10FilteringRangeCalculation{
	"LOWER",
	"UPPER",
}

// NewQianchuanReportSearchWordGetV10FilteringRangeCalculationFromValue returns a pointer to a valid QianchuanReportSearchWordGetV10FilteringRangeCalculation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportSearchWordGetV10FilteringRangeCalculationFromValue(v string) (*QianchuanReportSearchWordGetV10FilteringRangeCalculation, error) {
	ev := QianchuanReportSearchWordGetV10FilteringRangeCalculation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportSearchWordGetV10FilteringRangeCalculation: valid values are %v", v, AllowedQianchuanReportSearchWordGetV10FilteringRangeCalculationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportSearchWordGetV10FilteringRangeCalculation) IsValid() bool {
	for _, existing := range AllowedQianchuanReportSearchWordGetV10FilteringRangeCalculationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_search_word_get_v1.0_filtering_range_calculation value
func (v QianchuanReportSearchWordGetV10FilteringRangeCalculation) Ptr() *QianchuanReportSearchWordGetV10FilteringRangeCalculation {
	return &v
}
