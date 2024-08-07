/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportCampaignGetV2FilteringStatus
type ReportCampaignGetV2FilteringStatus string

// List of report_campaign_get_v2_filtering_status
const (
	CAMPAIGN_STATUS_ENABLE_ReportCampaignGetV2FilteringStatus                   ReportCampaignGetV2FilteringStatus = "CAMPAIGN_STATUS_ENABLE"
	CAMPAIGN_STATUS_ALL_ReportCampaignGetV2FilteringStatus                      ReportCampaignGetV2FilteringStatus = "CAMPAIGN_STATUS_ALL"
	CAMPAIGN_STATUS_DELETE_ReportCampaignGetV2FilteringStatus                   ReportCampaignGetV2FilteringStatus = "CAMPAIGN_STATUS_DELETE"
	CAMPAIGN_STATUS_ADVERTISER_BUDGET_EXCEED_ReportCampaignGetV2FilteringStatus ReportCampaignGetV2FilteringStatus = "CAMPAIGN_STATUS_ADVERTISER_BUDGET_EXCEED"
	CAMPAIGN_STATUS_DISABLE_ReportCampaignGetV2FilteringStatus                  ReportCampaignGetV2FilteringStatus = "CAMPAIGN_STATUS_DISABLE"
	CAMPAIGN_STATUS_NOT_DELETE_ReportCampaignGetV2FilteringStatus               ReportCampaignGetV2FilteringStatus = "CAMPAIGN_STATUS_NOT_DELETE"
)

// All allowed values of ReportCampaignGetV2FilteringStatus enum
var AllowedReportCampaignGetV2FilteringStatusEnumValues = []ReportCampaignGetV2FilteringStatus{
	"CAMPAIGN_STATUS_ENABLE",
	"CAMPAIGN_STATUS_ALL",
	"CAMPAIGN_STATUS_DELETE",
	"CAMPAIGN_STATUS_ADVERTISER_BUDGET_EXCEED",
	"CAMPAIGN_STATUS_DISABLE",
	"CAMPAIGN_STATUS_NOT_DELETE",
}

// NewReportCampaignGetV2FilteringStatusFromValue returns a pointer to a valid ReportCampaignGetV2FilteringStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCampaignGetV2FilteringStatusFromValue(v string) (*ReportCampaignGetV2FilteringStatus, error) {
	ev := ReportCampaignGetV2FilteringStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCampaignGetV2FilteringStatus: valid values are %v", v, AllowedReportCampaignGetV2FilteringStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCampaignGetV2FilteringStatus) IsValid() bool {
	for _, existing := range AllowedReportCampaignGetV2FilteringStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_campaign_get_v2_filtering_status value
func (v ReportCampaignGetV2FilteringStatus) Ptr() *ReportCampaignGetV2FilteringStatus {
	return &v
}
