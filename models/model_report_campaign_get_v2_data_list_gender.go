/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportCampaignGetV2DataListGender
type ReportCampaignGetV2DataListGender string

// List of report_campaign_get_v2_data_list_gender
const (
	NONE_ReportCampaignGetV2DataListGender             ReportCampaignGetV2DataListGender = "NONE"
	GENDER_UNLIMITED_ReportCampaignGetV2DataListGender ReportCampaignGetV2DataListGender = "GENDER_UNLIMITED"
	GENDER_FEMALE_ReportCampaignGetV2DataListGender    ReportCampaignGetV2DataListGender = "GENDER_FEMALE"
	GENDER_MALE_ReportCampaignGetV2DataListGender      ReportCampaignGetV2DataListGender = "GENDER_MALE"
)

// All allowed values of ReportCampaignGetV2DataListGender enum
var AllowedReportCampaignGetV2DataListGenderEnumValues = []ReportCampaignGetV2DataListGender{
	"NONE",
	"GENDER_UNLIMITED",
	"GENDER_FEMALE",
	"GENDER_MALE",
}

// NewReportCampaignGetV2DataListGenderFromValue returns a pointer to a valid ReportCampaignGetV2DataListGender
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCampaignGetV2DataListGenderFromValue(v string) (*ReportCampaignGetV2DataListGender, error) {
	ev := ReportCampaignGetV2DataListGender(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCampaignGetV2DataListGender: valid values are %v", v, AllowedReportCampaignGetV2DataListGenderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCampaignGetV2DataListGender) IsValid() bool {
	for _, existing := range AllowedReportCampaignGetV2DataListGenderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_campaign_get_v2_data_list_gender value
func (v ReportCampaignGetV2DataListGender) Ptr() *ReportCampaignGetV2DataListGender {
	return &v
}
