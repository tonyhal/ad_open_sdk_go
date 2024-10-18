/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCreativeGetV2DataListPricing
type ReportCreativeGetV2DataListPricing string

// List of report_creative_get_v2_data_list_pricing
const (
	PRICING_ECPC_ReportCreativeGetV2DataListPricing ReportCreativeGetV2DataListPricing = "PRICING_ECPC"
	PRICING_OCPC_ReportCreativeGetV2DataListPricing ReportCreativeGetV2DataListPricing = "PRICING_OCPC"
	PRICING_CPM_ReportCreativeGetV2DataListPricing  ReportCreativeGetV2DataListPricing = "PRICING_CPM"
	PRICING_OCPM_ReportCreativeGetV2DataListPricing ReportCreativeGetV2DataListPricing = "PRICING_OCPM"
	PRICING_CPA_ReportCreativeGetV2DataListPricing  ReportCreativeGetV2DataListPricing = "PRICING_CPA"
	PRICING_CPC_ReportCreativeGetV2DataListPricing  ReportCreativeGetV2DataListPricing = "PRICING_CPC"
	PRICING_CPV_ReportCreativeGetV2DataListPricing  ReportCreativeGetV2DataListPricing = "PRICING_CPV"
)

// Ptr returns reference to report_creative_get_v2_data_list_pricing value
func (v ReportCreativeGetV2DataListPricing) Ptr() *ReportCreativeGetV2DataListPricing {
	return &v
}
