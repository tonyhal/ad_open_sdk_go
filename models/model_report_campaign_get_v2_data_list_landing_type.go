/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportCampaignGetV2DataListLandingType
type ReportCampaignGetV2DataListLandingType string

// List of report_campaign_get_v2_data_list_landing_type
const (
	LINK_ReportCampaignGetV2DataListLandingType      ReportCampaignGetV2DataListLandingType = "LINK"
	APP_ReportCampaignGetV2DataListLandingType       ReportCampaignGetV2DataListLandingType = "APP"
	STORE_ReportCampaignGetV2DataListLandingType     ReportCampaignGetV2DataListLandingType = "STORE"
	LIVE_ReportCampaignGetV2DataListLandingType      ReportCampaignGetV2DataListLandingType = "LIVE"
	ARTICLE_ReportCampaignGetV2DataListLandingType   ReportCampaignGetV2DataListLandingType = "ARTICLE"
	QUICK_APP_ReportCampaignGetV2DataListLandingType ReportCampaignGetV2DataListLandingType = "QUICK_APP"
	AWEME_ReportCampaignGetV2DataListLandingType     ReportCampaignGetV2DataListLandingType = "AWEME"
	SHOP_ReportCampaignGetV2DataListLandingType      ReportCampaignGetV2DataListLandingType = "SHOP"
	DPA_ReportCampaignGetV2DataListLandingType       ReportCampaignGetV2DataListLandingType = "DPA"
	GOODS_ReportCampaignGetV2DataListLandingType     ReportCampaignGetV2DataListLandingType = "GOODS"
)

// All allowed values of ReportCampaignGetV2DataListLandingType enum
var AllowedReportCampaignGetV2DataListLandingTypeEnumValues = []ReportCampaignGetV2DataListLandingType{
	"LINK",
	"APP",
	"STORE",
	"LIVE",
	"ARTICLE",
	"QUICK_APP",
	"AWEME",
	"SHOP",
	"DPA",
	"GOODS",
}

// NewReportCampaignGetV2DataListLandingTypeFromValue returns a pointer to a valid ReportCampaignGetV2DataListLandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCampaignGetV2DataListLandingTypeFromValue(v string) (*ReportCampaignGetV2DataListLandingType, error) {
	ev := ReportCampaignGetV2DataListLandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCampaignGetV2DataListLandingType: valid values are %v", v, AllowedReportCampaignGetV2DataListLandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCampaignGetV2DataListLandingType) IsValid() bool {
	for _, existing := range AllowedReportCampaignGetV2DataListLandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_campaign_get_v2_data_list_landing_type value
func (v ReportCampaignGetV2DataListLandingType) Ptr() *ReportCampaignGetV2DataListLandingType {
	return &v
}
