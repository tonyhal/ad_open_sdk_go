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

// ReportAdvertiserGetV2DataListRealRecallMatchType
type ReportAdvertiserGetV2DataListRealRecallMatchType string

// List of report_advertiser_get_v2_data_list_real_recall_match_type
const (
	PHRASE_ReportAdvertiserGetV2DataListRealRecallMatchType    ReportAdvertiserGetV2DataListRealRecallMatchType = "PHRASE"
	EXTENSIVE_ReportAdvertiserGetV2DataListRealRecallMatchType ReportAdvertiserGetV2DataListRealRecallMatchType = "EXTENSIVE"
	PRECISION_ReportAdvertiserGetV2DataListRealRecallMatchType ReportAdvertiserGetV2DataListRealRecallMatchType = "PRECISION"
)

// All allowed values of ReportAdvertiserGetV2DataListRealRecallMatchType enum
var AllowedReportAdvertiserGetV2DataListRealRecallMatchTypeEnumValues = []ReportAdvertiserGetV2DataListRealRecallMatchType{
	"PHRASE",
	"EXTENSIVE",
	"PRECISION",
}

// NewReportAdvertiserGetV2DataListRealRecallMatchTypeFromValue returns a pointer to a valid ReportAdvertiserGetV2DataListRealRecallMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdvertiserGetV2DataListRealRecallMatchTypeFromValue(v string) (*ReportAdvertiserGetV2DataListRealRecallMatchType, error) {
	ev := ReportAdvertiserGetV2DataListRealRecallMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdvertiserGetV2DataListRealRecallMatchType: valid values are %v", v, AllowedReportAdvertiserGetV2DataListRealRecallMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdvertiserGetV2DataListRealRecallMatchType) IsValid() bool {
	for _, existing := range AllowedReportAdvertiserGetV2DataListRealRecallMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_advertiser_get_v2_data_list_real_recall_match_type value
func (v ReportAdvertiserGetV2DataListRealRecallMatchType) Ptr() *ReportAdvertiserGetV2DataListRealRecallMatchType {
	return &v
}
