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

// LocalReportPromotionGetV30OrderType
type LocalReportPromotionGetV30OrderType string

// List of local_report_promotion_get_v3.0_order_type
const (
	ASC_LocalReportPromotionGetV30OrderType  LocalReportPromotionGetV30OrderType = "ASC"
	DESC_LocalReportPromotionGetV30OrderType LocalReportPromotionGetV30OrderType = "DESC"
)

// All allowed values of LocalReportPromotionGetV30OrderType enum
var AllowedLocalReportPromotionGetV30OrderTypeEnumValues = []LocalReportPromotionGetV30OrderType{
	"ASC",
	"DESC",
}

// NewLocalReportPromotionGetV30OrderTypeFromValue returns a pointer to a valid LocalReportPromotionGetV30OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLocalReportPromotionGetV30OrderTypeFromValue(v string) (*LocalReportPromotionGetV30OrderType, error) {
	ev := LocalReportPromotionGetV30OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LocalReportPromotionGetV30OrderType: valid values are %v", v, AllowedLocalReportPromotionGetV30OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LocalReportPromotionGetV30OrderType) IsValid() bool {
	for _, existing := range AllowedLocalReportPromotionGetV30OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to local_report_promotion_get_v3.0_order_type value
func (v LocalReportPromotionGetV30OrderType) Ptr() *LocalReportPromotionGetV30OrderType {
	return &v
}