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

// ReportCreativeGetV2DataListGender
type ReportCreativeGetV2DataListGender string

// List of report_creative_get_v2_data_list_gender
const (
	GENDER_UNLIMITED_ReportCreativeGetV2DataListGender ReportCreativeGetV2DataListGender = "GENDER_UNLIMITED"
	GENDER_FEMALE_ReportCreativeGetV2DataListGender    ReportCreativeGetV2DataListGender = "GENDER_FEMALE"
	GENDER_MALE_ReportCreativeGetV2DataListGender      ReportCreativeGetV2DataListGender = "GENDER_MALE"
	NONE_ReportCreativeGetV2DataListGender             ReportCreativeGetV2DataListGender = "NONE"
)

// All allowed values of ReportCreativeGetV2DataListGender enum
var AllowedReportCreativeGetV2DataListGenderEnumValues = []ReportCreativeGetV2DataListGender{
	"GENDER_UNLIMITED",
	"GENDER_FEMALE",
	"GENDER_MALE",
	"NONE",
}

// NewReportCreativeGetV2DataListGenderFromValue returns a pointer to a valid ReportCreativeGetV2DataListGender
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCreativeGetV2DataListGenderFromValue(v string) (*ReportCreativeGetV2DataListGender, error) {
	ev := ReportCreativeGetV2DataListGender(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCreativeGetV2DataListGender: valid values are %v", v, AllowedReportCreativeGetV2DataListGenderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCreativeGetV2DataListGender) IsValid() bool {
	for _, existing := range AllowedReportCreativeGetV2DataListGenderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_creative_get_v2_data_list_gender value
func (v ReportCreativeGetV2DataListGender) Ptr() *ReportCreativeGetV2DataListGender {
	return &v
}
