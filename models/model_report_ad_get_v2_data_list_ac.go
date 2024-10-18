/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAdGetV2DataListAc
type ReportAdGetV2DataListAc string

// List of report_ad_get_v2_data_list_ac
const (
	Enum_2_G_ReportAdGetV2DataListAc ReportAdGetV2DataListAc = "2G"
	Enum_4_G_ReportAdGetV2DataListAc ReportAdGetV2DataListAc = "4G"
	Enum_5_G_ReportAdGetV2DataListAc ReportAdGetV2DataListAc = "5G"
	UNKNOWN_ReportAdGetV2DataListAc  ReportAdGetV2DataListAc = "unknown"
	WIFI_ReportAdGetV2DataListAc     ReportAdGetV2DataListAc = "WIFI"
	Enum_3_G_ReportAdGetV2DataListAc ReportAdGetV2DataListAc = "3G"
)

// Ptr returns reference to report_ad_get_v2_data_list_ac value
func (v ReportAdGetV2DataListAc) Ptr() *ReportAdGetV2DataListAc {
	return &v
}
