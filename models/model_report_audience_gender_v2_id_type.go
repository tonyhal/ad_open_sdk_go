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

// ReportAudienceGenderV2IdType
type ReportAudienceGenderV2IdType string

// List of report_audience_gender_v2_id_type
const (
	AUDIENCE_STAT_ID_TYPE_ADVERTISER_ReportAudienceGenderV2IdType ReportAudienceGenderV2IdType = "AUDIENCE_STAT_ID_TYPE_ADVERTISER"
	AUDIENCE_STAT_ID_TYPE_AD_ReportAudienceGenderV2IdType         ReportAudienceGenderV2IdType = "AUDIENCE_STAT_ID_TYPE_AD"
	AUDIENCE_STAT_ID_TYPE_CAMPAIGN_ReportAudienceGenderV2IdType   ReportAudienceGenderV2IdType = "AUDIENCE_STAT_ID_TYPE_CAMPAIGN"
)

// All allowed values of ReportAudienceGenderV2IdType enum
var AllowedReportAudienceGenderV2IdTypeEnumValues = []ReportAudienceGenderV2IdType{
	"AUDIENCE_STAT_ID_TYPE_ADVERTISER",
	"AUDIENCE_STAT_ID_TYPE_AD",
	"AUDIENCE_STAT_ID_TYPE_CAMPAIGN",
}

// NewReportAudienceGenderV2IdTypeFromValue returns a pointer to a valid ReportAudienceGenderV2IdType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAudienceGenderV2IdTypeFromValue(v string) (*ReportAudienceGenderV2IdType, error) {
	ev := ReportAudienceGenderV2IdType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAudienceGenderV2IdType: valid values are %v", v, AllowedReportAudienceGenderV2IdTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAudienceGenderV2IdType) IsValid() bool {
	for _, existing := range AllowedReportAudienceGenderV2IdTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_audience_gender_v2_id_type value
func (v ReportAudienceGenderV2IdType) Ptr() *ReportAudienceGenderV2IdType {
	return &v
}
