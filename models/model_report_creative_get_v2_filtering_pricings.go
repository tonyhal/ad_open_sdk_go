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

// ReportCreativeGetV2FilteringPricings
type ReportCreativeGetV2FilteringPricings string

// List of report_creative_get_v2_filtering_pricings
const (
	PRICING_CPC_ReportCreativeGetV2FilteringPricings  ReportCreativeGetV2FilteringPricings = "PRICING_CPC"
	PRICING_CPA_ReportCreativeGetV2FilteringPricings  ReportCreativeGetV2FilteringPricings = "PRICING_CPA"
	PRICING_ECPC_ReportCreativeGetV2FilteringPricings ReportCreativeGetV2FilteringPricings = "PRICING_ECPC"
	PRICING_CPV_ReportCreativeGetV2FilteringPricings  ReportCreativeGetV2FilteringPricings = "PRICING_CPV"
	PRICING_OCPM_ReportCreativeGetV2FilteringPricings ReportCreativeGetV2FilteringPricings = "PRICING_OCPM"
	PRICING_OCPC_ReportCreativeGetV2FilteringPricings ReportCreativeGetV2FilteringPricings = "PRICING_OCPC"
	PRICING_CPM_ReportCreativeGetV2FilteringPricings  ReportCreativeGetV2FilteringPricings = "PRICING_CPM"
)

// All allowed values of ReportCreativeGetV2FilteringPricings enum
var AllowedReportCreativeGetV2FilteringPricingsEnumValues = []ReportCreativeGetV2FilteringPricings{
	"PRICING_CPC",
	"PRICING_CPA",
	"PRICING_ECPC",
	"PRICING_CPV",
	"PRICING_OCPM",
	"PRICING_OCPC",
	"PRICING_CPM",
}

// NewReportCreativeGetV2FilteringPricingsFromValue returns a pointer to a valid ReportCreativeGetV2FilteringPricings
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCreativeGetV2FilteringPricingsFromValue(v string) (*ReportCreativeGetV2FilteringPricings, error) {
	ev := ReportCreativeGetV2FilteringPricings(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCreativeGetV2FilteringPricings: valid values are %v", v, AllowedReportCreativeGetV2FilteringPricingsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCreativeGetV2FilteringPricings) IsValid() bool {
	for _, existing := range AllowedReportCreativeGetV2FilteringPricingsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_creative_get_v2_filtering_pricings value
func (v ReportCreativeGetV2FilteringPricings) Ptr() *ReportCreativeGetV2FilteringPricings {
	return &v
}
