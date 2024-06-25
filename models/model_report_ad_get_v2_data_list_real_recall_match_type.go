/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportAdGetV2DataListRealRecallMatchType
type ReportAdGetV2DataListRealRecallMatchType string

// List of report_ad_get_v2_data_list_real_recall_match_type
const (
	PHRASE_ReportAdGetV2DataListRealRecallMatchType    ReportAdGetV2DataListRealRecallMatchType = "PHRASE"
	EXTENSIVE_ReportAdGetV2DataListRealRecallMatchType ReportAdGetV2DataListRealRecallMatchType = "EXTENSIVE"
	PRECISION_ReportAdGetV2DataListRealRecallMatchType ReportAdGetV2DataListRealRecallMatchType = "PRECISION"
)

// All allowed values of ReportAdGetV2DataListRealRecallMatchType enum
var AllowedReportAdGetV2DataListRealRecallMatchTypeEnumValues = []ReportAdGetV2DataListRealRecallMatchType{
	"PHRASE",
	"EXTENSIVE",
	"PRECISION",
}

// NewReportAdGetV2DataListRealRecallMatchTypeFromValue returns a pointer to a valid ReportAdGetV2DataListRealRecallMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdGetV2DataListRealRecallMatchTypeFromValue(v string) (*ReportAdGetV2DataListRealRecallMatchType, error) {
	ev := ReportAdGetV2DataListRealRecallMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdGetV2DataListRealRecallMatchType: valid values are %v", v, AllowedReportAdGetV2DataListRealRecallMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdGetV2DataListRealRecallMatchType) IsValid() bool {
	for _, existing := range AllowedReportAdGetV2DataListRealRecallMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_ad_get_v2_data_list_real_recall_match_type value
func (v ReportAdGetV2DataListRealRecallMatchType) Ptr() *ReportAdGetV2DataListRealRecallMatchType {
	return &v
}
