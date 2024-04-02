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

// ReportAudienceInterestActionListV2FilteringInterestActionType
type ReportAudienceInterestActionListV2FilteringInterestActionType string

// List of report_audience_interest_action_list_v2_filtering_interest_action_type
const (
	ACTION_ReportAudienceInterestActionListV2FilteringInterestActionType   ReportAudienceInterestActionListV2FilteringInterestActionType = "ACTION"
	INTEREST_ReportAudienceInterestActionListV2FilteringInterestActionType ReportAudienceInterestActionListV2FilteringInterestActionType = "INTEREST"
)

// All allowed values of ReportAudienceInterestActionListV2FilteringInterestActionType enum
var AllowedReportAudienceInterestActionListV2FilteringInterestActionTypeEnumValues = []ReportAudienceInterestActionListV2FilteringInterestActionType{
	"ACTION",
	"INTEREST",
}

// NewReportAudienceInterestActionListV2FilteringInterestActionTypeFromValue returns a pointer to a valid ReportAudienceInterestActionListV2FilteringInterestActionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAudienceInterestActionListV2FilteringInterestActionTypeFromValue(v string) (*ReportAudienceInterestActionListV2FilteringInterestActionType, error) {
	ev := ReportAudienceInterestActionListV2FilteringInterestActionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAudienceInterestActionListV2FilteringInterestActionType: valid values are %v", v, AllowedReportAudienceInterestActionListV2FilteringInterestActionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAudienceInterestActionListV2FilteringInterestActionType) IsValid() bool {
	for _, existing := range AllowedReportAudienceInterestActionListV2FilteringInterestActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_audience_interest_action_list_v2_filtering_interest_action_type value
func (v ReportAudienceInterestActionListV2FilteringInterestActionType) Ptr() *ReportAudienceInterestActionListV2FilteringInterestActionType {
	return &v
}
