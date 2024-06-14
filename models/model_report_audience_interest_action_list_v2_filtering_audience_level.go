/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportAudienceInterestActionListV2FilteringAudienceLevel
type ReportAudienceInterestActionListV2FilteringAudienceLevel string

// List of report_audience_interest_action_list_v2_filtering_audience_level
const (
	THIRD_LEVEL_ReportAudienceInterestActionListV2FilteringAudienceLevel   ReportAudienceInterestActionListV2FilteringAudienceLevel = "THIRD_LEVEL"
	KEYWORD_AWEME_ReportAudienceInterestActionListV2FilteringAudienceLevel ReportAudienceInterestActionListV2FilteringAudienceLevel = "KEYWORD_AWEME"
	FOURTH_LEVEL_ReportAudienceInterestActionListV2FilteringAudienceLevel  ReportAudienceInterestActionListV2FilteringAudienceLevel = "FOURTH_LEVEL"
	SECOND_LEVEL_ReportAudienceInterestActionListV2FilteringAudienceLevel  ReportAudienceInterestActionListV2FilteringAudienceLevel = "SECOND_LEVEL"
	FIRST_LEVEL_ReportAudienceInterestActionListV2FilteringAudienceLevel   ReportAudienceInterestActionListV2FilteringAudienceLevel = "FIRST_LEVEL"
)

// All allowed values of ReportAudienceInterestActionListV2FilteringAudienceLevel enum
var AllowedReportAudienceInterestActionListV2FilteringAudienceLevelEnumValues = []ReportAudienceInterestActionListV2FilteringAudienceLevel{
	"THIRD_LEVEL",
	"KEYWORD_AWEME",
	"FOURTH_LEVEL",
	"SECOND_LEVEL",
	"FIRST_LEVEL",
}

// NewReportAudienceInterestActionListV2FilteringAudienceLevelFromValue returns a pointer to a valid ReportAudienceInterestActionListV2FilteringAudienceLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAudienceInterestActionListV2FilteringAudienceLevelFromValue(v string) (*ReportAudienceInterestActionListV2FilteringAudienceLevel, error) {
	ev := ReportAudienceInterestActionListV2FilteringAudienceLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAudienceInterestActionListV2FilteringAudienceLevel: valid values are %v", v, AllowedReportAudienceInterestActionListV2FilteringAudienceLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAudienceInterestActionListV2FilteringAudienceLevel) IsValid() bool {
	for _, existing := range AllowedReportAudienceInterestActionListV2FilteringAudienceLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_audience_interest_action_list_v2_filtering_audience_level value
func (v ReportAudienceInterestActionListV2FilteringAudienceLevel) Ptr() *ReportAudienceInterestActionListV2FilteringAudienceLevel {
	return &v
}
