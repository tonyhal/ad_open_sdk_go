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

// QianchuanReportUniPromotionGetV10LabAdType
type QianchuanReportUniPromotionGetV10LabAdType string

// List of qianchuan_report_uni_promotion_get_v1.0_lab_ad_type
const (
	LAB_AD_QianchuanReportUniPromotionGetV10LabAdType QianchuanReportUniPromotionGetV10LabAdType = "LAB_AD"
)

// All allowed values of QianchuanReportUniPromotionGetV10LabAdType enum
var AllowedQianchuanReportUniPromotionGetV10LabAdTypeEnumValues = []QianchuanReportUniPromotionGetV10LabAdType{
	"LAB_AD",
}

// NewQianchuanReportUniPromotionGetV10LabAdTypeFromValue returns a pointer to a valid QianchuanReportUniPromotionGetV10LabAdType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportUniPromotionGetV10LabAdTypeFromValue(v string) (*QianchuanReportUniPromotionGetV10LabAdType, error) {
	ev := QianchuanReportUniPromotionGetV10LabAdType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportUniPromotionGetV10LabAdType: valid values are %v", v, AllowedQianchuanReportUniPromotionGetV10LabAdTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportUniPromotionGetV10LabAdType) IsValid() bool {
	for _, existing := range AllowedQianchuanReportUniPromotionGetV10LabAdTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_uni_promotion_get_v1.0_lab_ad_type value
func (v QianchuanReportUniPromotionGetV10LabAdType) Ptr() *QianchuanReportUniPromotionGetV10LabAdType {
	return &v
}
