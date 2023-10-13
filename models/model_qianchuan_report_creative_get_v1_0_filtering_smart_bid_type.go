/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanReportCreativeGetV10FilteringSmartBidType
type QianchuanReportCreativeGetV10FilteringSmartBidType string

// List of qianchuan_report_creative_get_v1.0_filtering_smart_bid_type
const (
	SMART_BID_CONSERVATIVE_QianchuanReportCreativeGetV10FilteringSmartBidType QianchuanReportCreativeGetV10FilteringSmartBidType = "SMART_BID_CONSERVATIVE"
	SMART_BID_CUSTOM_QianchuanReportCreativeGetV10FilteringSmartBidType       QianchuanReportCreativeGetV10FilteringSmartBidType = "SMART_BID_CUSTOM"
)

// All allowed values of QianchuanReportCreativeGetV10FilteringSmartBidType enum
var AllowedQianchuanReportCreativeGetV10FilteringSmartBidTypeEnumValues = []QianchuanReportCreativeGetV10FilteringSmartBidType{
	"SMART_BID_CONSERVATIVE",
	"SMART_BID_CUSTOM",
}

// NewQianchuanReportCreativeGetV10FilteringSmartBidTypeFromValue returns a pointer to a valid QianchuanReportCreativeGetV10FilteringSmartBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportCreativeGetV10FilteringSmartBidTypeFromValue(v string) (*QianchuanReportCreativeGetV10FilteringSmartBidType, error) {
	ev := QianchuanReportCreativeGetV10FilteringSmartBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportCreativeGetV10FilteringSmartBidType: valid values are %v", v, AllowedQianchuanReportCreativeGetV10FilteringSmartBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportCreativeGetV10FilteringSmartBidType) IsValid() bool {
	for _, existing := range AllowedQianchuanReportCreativeGetV10FilteringSmartBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_creative_get_v1.0_filtering_smart_bid_type value
func (v QianchuanReportCreativeGetV10FilteringSmartBidType) Ptr() *QianchuanReportCreativeGetV10FilteringSmartBidType {
	return &v
}
