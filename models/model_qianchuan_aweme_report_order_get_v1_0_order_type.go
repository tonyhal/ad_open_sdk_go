/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAwemeReportOrderGetV10OrderType
type QianchuanAwemeReportOrderGetV10OrderType string

// List of qianchuan_aweme_report_order_get_v1.0_order_type
const (
	ASC_QianchuanAwemeReportOrderGetV10OrderType  QianchuanAwemeReportOrderGetV10OrderType = "ASC"
	DESC_QianchuanAwemeReportOrderGetV10OrderType QianchuanAwemeReportOrderGetV10OrderType = "DESC"
)

// All allowed values of QianchuanAwemeReportOrderGetV10OrderType enum
var AllowedQianchuanAwemeReportOrderGetV10OrderTypeEnumValues = []QianchuanAwemeReportOrderGetV10OrderType{
	"ASC",
	"DESC",
}

// NewQianchuanAwemeReportOrderGetV10OrderTypeFromValue returns a pointer to a valid QianchuanAwemeReportOrderGetV10OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeReportOrderGetV10OrderTypeFromValue(v string) (*QianchuanAwemeReportOrderGetV10OrderType, error) {
	ev := QianchuanAwemeReportOrderGetV10OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeReportOrderGetV10OrderType: valid values are %v", v, AllowedQianchuanAwemeReportOrderGetV10OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeReportOrderGetV10OrderType) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeReportOrderGetV10OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_report_order_get_v1.0_order_type value
func (v QianchuanAwemeReportOrderGetV10OrderType) Ptr() *QianchuanAwemeReportOrderGetV10OrderType {
	return &v
}