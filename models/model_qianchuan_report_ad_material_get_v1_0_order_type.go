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

// QianchuanReportAdMaterialGetV10OrderType
type QianchuanReportAdMaterialGetV10OrderType string

// List of qianchuan_report_ad_material_get_v1.0_order_type
const (
	ASC_QianchuanReportAdMaterialGetV10OrderType  QianchuanReportAdMaterialGetV10OrderType = "ASC"
	DESC_QianchuanReportAdMaterialGetV10OrderType QianchuanReportAdMaterialGetV10OrderType = "DESC"
)

// All allowed values of QianchuanReportAdMaterialGetV10OrderType enum
var AllowedQianchuanReportAdMaterialGetV10OrderTypeEnumValues = []QianchuanReportAdMaterialGetV10OrderType{
	"ASC",
	"DESC",
}

// NewQianchuanReportAdMaterialGetV10OrderTypeFromValue returns a pointer to a valid QianchuanReportAdMaterialGetV10OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportAdMaterialGetV10OrderTypeFromValue(v string) (*QianchuanReportAdMaterialGetV10OrderType, error) {
	ev := QianchuanReportAdMaterialGetV10OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportAdMaterialGetV10OrderType: valid values are %v", v, AllowedQianchuanReportAdMaterialGetV10OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportAdMaterialGetV10OrderType) IsValid() bool {
	for _, existing := range AllowedQianchuanReportAdMaterialGetV10OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_ad_material_get_v1.0_order_type value
func (v QianchuanReportAdMaterialGetV10OrderType) Ptr() *QianchuanReportAdMaterialGetV10OrderType {
	return &v
}