/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAdGetV2DataListPricing
type ReportAdGetV2DataListPricing string

// List of report_ad_get_v2_data_list_pricing
const (
	PRICING_CPA_ReportAdGetV2DataListPricing  ReportAdGetV2DataListPricing = "PRICING_CPA"
	PRICING_OCPC_ReportAdGetV2DataListPricing ReportAdGetV2DataListPricing = "PRICING_OCPC"
	PRICING_ECPC_ReportAdGetV2DataListPricing ReportAdGetV2DataListPricing = "PRICING_ECPC"
	PRICING_OCPM_ReportAdGetV2DataListPricing ReportAdGetV2DataListPricing = "PRICING_OCPM"
	PRICING_CPC_ReportAdGetV2DataListPricing  ReportAdGetV2DataListPricing = "PRICING_CPC"
	PRICING_CPM_ReportAdGetV2DataListPricing  ReportAdGetV2DataListPricing = "PRICING_CPM"
	PRICING_CPV_ReportAdGetV2DataListPricing  ReportAdGetV2DataListPricing = "PRICING_CPV"
)

// Ptr returns reference to report_ad_get_v2_data_list_pricing value
func (v ReportAdGetV2DataListPricing) Ptr() *ReportAdGetV2DataListPricing {
	return &v
}
