/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAdvertiserGetV2DataListAc
type ReportAdvertiserGetV2DataListAc string

// List of report_advertiser_get_v2_data_list_ac
const (
	Enum_2_G_ReportAdvertiserGetV2DataListAc ReportAdvertiserGetV2DataListAc = "2G"
	UNKNOWN_ReportAdvertiserGetV2DataListAc  ReportAdvertiserGetV2DataListAc = "unknown"
	Enum_4_G_ReportAdvertiserGetV2DataListAc ReportAdvertiserGetV2DataListAc = "4G"
	Enum_5_G_ReportAdvertiserGetV2DataListAc ReportAdvertiserGetV2DataListAc = "5G"
	Enum_3_G_ReportAdvertiserGetV2DataListAc ReportAdvertiserGetV2DataListAc = "3G"
	WIFI_ReportAdvertiserGetV2DataListAc     ReportAdvertiserGetV2DataListAc = "WIFI"
)

// Ptr returns reference to report_advertiser_get_v2_data_list_ac value
func (v ReportAdvertiserGetV2DataListAc) Ptr() *ReportAdvertiserGetV2DataListAc {
	return &v
}
