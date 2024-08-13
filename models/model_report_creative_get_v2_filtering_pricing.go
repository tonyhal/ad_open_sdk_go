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

// ReportCreativeGetV2FilteringPricing
type ReportCreativeGetV2FilteringPricing string

// List of report_creative_get_v2_filtering_pricing
const (
	PRICING_CPV_ReportCreativeGetV2FilteringPricing      ReportCreativeGetV2FilteringPricing = "PRICING_CPV"
	PRICING_CPM_ReportCreativeGetV2FilteringPricing      ReportCreativeGetV2FilteringPricing = "PRICING_CPM"
	PRICING_CPC_OCPM_ReportCreativeGetV2FilteringPricing ReportCreativeGetV2FilteringPricing = "PRICING_CPC_OCPM"
	PRICING_CPC_ReportCreativeGetV2FilteringPricing      ReportCreativeGetV2FilteringPricing = "PRICING_CPC"
	PRICING_OCPM_ReportCreativeGetV2FilteringPricing     ReportCreativeGetV2FilteringPricing = "PRICING_OCPM"
	PRICING_OCPC_ReportCreativeGetV2FilteringPricing     ReportCreativeGetV2FilteringPricing = "PRICING_OCPC"
	PRICING_CPA_ReportCreativeGetV2FilteringPricing      ReportCreativeGetV2FilteringPricing = "PRICING_CPA"
)

// All allowed values of ReportCreativeGetV2FilteringPricing enum
var AllowedReportCreativeGetV2FilteringPricingEnumValues = []ReportCreativeGetV2FilteringPricing{
	"PRICING_CPV",
	"PRICING_CPM",
	"PRICING_CPC_OCPM",
	"PRICING_CPC",
	"PRICING_OCPM",
	"PRICING_OCPC",
	"PRICING_CPA",
}

// NewReportCreativeGetV2FilteringPricingFromValue returns a pointer to a valid ReportCreativeGetV2FilteringPricing
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCreativeGetV2FilteringPricingFromValue(v string) (*ReportCreativeGetV2FilteringPricing, error) {
	ev := ReportCreativeGetV2FilteringPricing(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCreativeGetV2FilteringPricing: valid values are %v", v, AllowedReportCreativeGetV2FilteringPricingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCreativeGetV2FilteringPricing) IsValid() bool {
	for _, existing := range AllowedReportCreativeGetV2FilteringPricingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_creative_get_v2_filtering_pricing value
func (v ReportCreativeGetV2FilteringPricing) Ptr() *ReportCreativeGetV2FilteringPricing {
	return &v
}
