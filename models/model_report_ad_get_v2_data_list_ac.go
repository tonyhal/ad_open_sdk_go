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

// ReportAdGetV2DataListAc
type ReportAdGetV2DataListAc string

// List of report_ad_get_v2_data_list_ac
const (
	Enum_5_G_ReportAdGetV2DataListAc ReportAdGetV2DataListAc = "5G"
	Enum_3_G_ReportAdGetV2DataListAc ReportAdGetV2DataListAc = "3G"
	UNKNOWN_ReportAdGetV2DataListAc  ReportAdGetV2DataListAc = "unknown"
	WIFI_ReportAdGetV2DataListAc     ReportAdGetV2DataListAc = "WIFI"
	Enum_4_G_ReportAdGetV2DataListAc ReportAdGetV2DataListAc = "4G"
	Enum_2_G_ReportAdGetV2DataListAc ReportAdGetV2DataListAc = "2G"
)

// All allowed values of ReportAdGetV2DataListAc enum
var AllowedReportAdGetV2DataListAcEnumValues = []ReportAdGetV2DataListAc{
	"5G",
	"3G",
	"unknown",
	"WIFI",
	"4G",
	"2G",
}

// NewReportAdGetV2DataListAcFromValue returns a pointer to a valid ReportAdGetV2DataListAc
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdGetV2DataListAcFromValue(v string) (*ReportAdGetV2DataListAc, error) {
	ev := ReportAdGetV2DataListAc(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdGetV2DataListAc: valid values are %v", v, AllowedReportAdGetV2DataListAcEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdGetV2DataListAc) IsValid() bool {
	for _, existing := range AllowedReportAdGetV2DataListAcEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_ad_get_v2_data_list_ac value
func (v ReportAdGetV2DataListAc) Ptr() *ReportAdGetV2DataListAc {
	return &v
}
