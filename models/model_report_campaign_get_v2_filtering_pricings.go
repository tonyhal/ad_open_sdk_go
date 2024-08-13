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

// ReportCampaignGetV2FilteringPricings
type ReportCampaignGetV2FilteringPricings string

// List of report_campaign_get_v2_filtering_pricings
const (
	PRICING_CPV_ReportCampaignGetV2FilteringPricings  ReportCampaignGetV2FilteringPricings = "PRICING_CPV"
	PRICING_CPM_ReportCampaignGetV2FilteringPricings  ReportCampaignGetV2FilteringPricings = "PRICING_CPM"
	PRICING_ECPC_ReportCampaignGetV2FilteringPricings ReportCampaignGetV2FilteringPricings = "PRICING_ECPC"
	PRICING_CPC_ReportCampaignGetV2FilteringPricings  ReportCampaignGetV2FilteringPricings = "PRICING_CPC"
	PRICING_OCPM_ReportCampaignGetV2FilteringPricings ReportCampaignGetV2FilteringPricings = "PRICING_OCPM"
	PRICING_OCPC_ReportCampaignGetV2FilteringPricings ReportCampaignGetV2FilteringPricings = "PRICING_OCPC"
	PRICING_CPA_ReportCampaignGetV2FilteringPricings  ReportCampaignGetV2FilteringPricings = "PRICING_CPA"
)

// All allowed values of ReportCampaignGetV2FilteringPricings enum
var AllowedReportCampaignGetV2FilteringPricingsEnumValues = []ReportCampaignGetV2FilteringPricings{
	"PRICING_CPV",
	"PRICING_CPM",
	"PRICING_ECPC",
	"PRICING_CPC",
	"PRICING_OCPM",
	"PRICING_OCPC",
	"PRICING_CPA",
}

// NewReportCampaignGetV2FilteringPricingsFromValue returns a pointer to a valid ReportCampaignGetV2FilteringPricings
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCampaignGetV2FilteringPricingsFromValue(v string) (*ReportCampaignGetV2FilteringPricings, error) {
	ev := ReportCampaignGetV2FilteringPricings(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCampaignGetV2FilteringPricings: valid values are %v", v, AllowedReportCampaignGetV2FilteringPricingsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCampaignGetV2FilteringPricings) IsValid() bool {
	for _, existing := range AllowedReportCampaignGetV2FilteringPricingsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_campaign_get_v2_filtering_pricings value
func (v ReportCampaignGetV2FilteringPricings) Ptr() *ReportCampaignGetV2FilteringPricings {
	return &v
}
