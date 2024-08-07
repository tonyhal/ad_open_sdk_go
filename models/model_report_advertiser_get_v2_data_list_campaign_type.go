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

// ReportAdvertiserGetV2DataListCampaignType
type ReportAdvertiserGetV2DataListCampaignType string

// List of report_advertiser_get_v2_data_list_campaign_type
const (
	SEARCH_ReportAdvertiserGetV2DataListCampaignType ReportAdvertiserGetV2DataListCampaignType = "SEARCH"
	FEED_ReportAdvertiserGetV2DataListCampaignType   ReportAdvertiserGetV2DataListCampaignType = "FEED"
)

// All allowed values of ReportAdvertiserGetV2DataListCampaignType enum
var AllowedReportAdvertiserGetV2DataListCampaignTypeEnumValues = []ReportAdvertiserGetV2DataListCampaignType{
	"SEARCH",
	"FEED",
}

// NewReportAdvertiserGetV2DataListCampaignTypeFromValue returns a pointer to a valid ReportAdvertiserGetV2DataListCampaignType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdvertiserGetV2DataListCampaignTypeFromValue(v string) (*ReportAdvertiserGetV2DataListCampaignType, error) {
	ev := ReportAdvertiserGetV2DataListCampaignType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdvertiserGetV2DataListCampaignType: valid values are %v", v, AllowedReportAdvertiserGetV2DataListCampaignTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdvertiserGetV2DataListCampaignType) IsValid() bool {
	for _, existing := range AllowedReportAdvertiserGetV2DataListCampaignTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_advertiser_get_v2_data_list_campaign_type value
func (v ReportAdvertiserGetV2DataListCampaignType) Ptr() *ReportAdvertiserGetV2DataListCampaignType {
	return &v
}
