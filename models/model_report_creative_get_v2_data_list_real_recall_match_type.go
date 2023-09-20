/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportCreativeGetV2DataListRealRecallMatchType
type ReportCreativeGetV2DataListRealRecallMatchType string

// List of report_creative_get_v2_data_list_real_recall_match_type
const (
	PHRASE_ReportCreativeGetV2DataListRealRecallMatchType    ReportCreativeGetV2DataListRealRecallMatchType = "PHRASE"
	PRECISION_ReportCreativeGetV2DataListRealRecallMatchType ReportCreativeGetV2DataListRealRecallMatchType = "PRECISION"
	EXTENSIVE_ReportCreativeGetV2DataListRealRecallMatchType ReportCreativeGetV2DataListRealRecallMatchType = "EXTENSIVE"
)

// All allowed values of ReportCreativeGetV2DataListRealRecallMatchType enum
var AllowedReportCreativeGetV2DataListRealRecallMatchTypeEnumValues = []ReportCreativeGetV2DataListRealRecallMatchType{
	"PHRASE",
	"PRECISION",
	"EXTENSIVE",
}

// NewReportCreativeGetV2DataListRealRecallMatchTypeFromValue returns a pointer to a valid ReportCreativeGetV2DataListRealRecallMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCreativeGetV2DataListRealRecallMatchTypeFromValue(v string) (*ReportCreativeGetV2DataListRealRecallMatchType, error) {
	ev := ReportCreativeGetV2DataListRealRecallMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCreativeGetV2DataListRealRecallMatchType: valid values are %v", v, AllowedReportCreativeGetV2DataListRealRecallMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCreativeGetV2DataListRealRecallMatchType) IsValid() bool {
	for _, existing := range AllowedReportCreativeGetV2DataListRealRecallMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_creative_get_v2_data_list_real_recall_match_type value
func (v ReportCreativeGetV2DataListRealRecallMatchType) Ptr() *ReportCreativeGetV2DataListRealRecallMatchType {
	return &v
}