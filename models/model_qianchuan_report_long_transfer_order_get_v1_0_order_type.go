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

// QianchuanReportLongTransferOrderGetV10OrderType
type QianchuanReportLongTransferOrderGetV10OrderType string

// List of qianchuan_report_long_transfer_order_get_v1.0_order_type
const (
	ASC_QianchuanReportLongTransferOrderGetV10OrderType  QianchuanReportLongTransferOrderGetV10OrderType = "ASC"
	DESC_QianchuanReportLongTransferOrderGetV10OrderType QianchuanReportLongTransferOrderGetV10OrderType = "DESC"
)

// All allowed values of QianchuanReportLongTransferOrderGetV10OrderType enum
var AllowedQianchuanReportLongTransferOrderGetV10OrderTypeEnumValues = []QianchuanReportLongTransferOrderGetV10OrderType{
	"ASC",
	"DESC",
}

// NewQianchuanReportLongTransferOrderGetV10OrderTypeFromValue returns a pointer to a valid QianchuanReportLongTransferOrderGetV10OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportLongTransferOrderGetV10OrderTypeFromValue(v string) (*QianchuanReportLongTransferOrderGetV10OrderType, error) {
	ev := QianchuanReportLongTransferOrderGetV10OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportLongTransferOrderGetV10OrderType: valid values are %v", v, AllowedQianchuanReportLongTransferOrderGetV10OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportLongTransferOrderGetV10OrderType) IsValid() bool {
	for _, existing := range AllowedQianchuanReportLongTransferOrderGetV10OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_long_transfer_order_get_v1.0_order_type value
func (v QianchuanReportLongTransferOrderGetV10OrderType) Ptr() *QianchuanReportLongTransferOrderGetV10OrderType {
	return &v
}
