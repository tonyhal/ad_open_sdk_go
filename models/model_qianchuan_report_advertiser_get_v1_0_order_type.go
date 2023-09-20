/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanReportAdvertiserGetV10OrderType
type QianchuanReportAdvertiserGetV10OrderType string

// List of qianchuan_report_advertiser_get_v1.0_order_type
const (
	ASC_QianchuanReportAdvertiserGetV10OrderType  QianchuanReportAdvertiserGetV10OrderType = "ASC"
	DESC_QianchuanReportAdvertiserGetV10OrderType QianchuanReportAdvertiserGetV10OrderType = "DESC"
)

// All allowed values of QianchuanReportAdvertiserGetV10OrderType enum
var AllowedQianchuanReportAdvertiserGetV10OrderTypeEnumValues = []QianchuanReportAdvertiserGetV10OrderType{
	"ASC",
	"DESC",
}

// NewQianchuanReportAdvertiserGetV10OrderTypeFromValue returns a pointer to a valid QianchuanReportAdvertiserGetV10OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportAdvertiserGetV10OrderTypeFromValue(v string) (*QianchuanReportAdvertiserGetV10OrderType, error) {
	ev := QianchuanReportAdvertiserGetV10OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportAdvertiserGetV10OrderType: valid values are %v", v, AllowedQianchuanReportAdvertiserGetV10OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportAdvertiserGetV10OrderType) IsValid() bool {
	for _, existing := range AllowedQianchuanReportAdvertiserGetV10OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_advertiser_get_v1.0_order_type value
func (v QianchuanReportAdvertiserGetV10OrderType) Ptr() *QianchuanReportAdvertiserGetV10OrderType {
	return &v
}
