/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanReportLongTransferOrderGetV10DataListIsPay
type QianchuanReportLongTransferOrderGetV10DataListIsPay string

// List of qianchuan_report_long_transfer_order_get_v1.0_data_list_is_pay
const (
	PAID_QianchuanReportLongTransferOrderGetV10DataListIsPay   QianchuanReportLongTransferOrderGetV10DataListIsPay = "PAID"
	UNPAID_QianchuanReportLongTransferOrderGetV10DataListIsPay QianchuanReportLongTransferOrderGetV10DataListIsPay = "UNPAID"
)

// All allowed values of QianchuanReportLongTransferOrderGetV10DataListIsPay enum
var AllowedQianchuanReportLongTransferOrderGetV10DataListIsPayEnumValues = []QianchuanReportLongTransferOrderGetV10DataListIsPay{
	"PAID",
	"UNPAID",
}

// NewQianchuanReportLongTransferOrderGetV10DataListIsPayFromValue returns a pointer to a valid QianchuanReportLongTransferOrderGetV10DataListIsPay
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportLongTransferOrderGetV10DataListIsPayFromValue(v string) (*QianchuanReportLongTransferOrderGetV10DataListIsPay, error) {
	ev := QianchuanReportLongTransferOrderGetV10DataListIsPay(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportLongTransferOrderGetV10DataListIsPay: valid values are %v", v, AllowedQianchuanReportLongTransferOrderGetV10DataListIsPayEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportLongTransferOrderGetV10DataListIsPay) IsValid() bool {
	for _, existing := range AllowedQianchuanReportLongTransferOrderGetV10DataListIsPayEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_long_transfer_order_get_v1.0_data_list_is_pay value
func (v QianchuanReportLongTransferOrderGetV10DataListIsPay) Ptr() *QianchuanReportLongTransferOrderGetV10DataListIsPay {
	return &v
}
