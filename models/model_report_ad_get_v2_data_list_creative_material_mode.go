/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAdGetV2DataListCreativeMaterialMode
type ReportAdGetV2DataListCreativeMaterialMode string

// List of report_ad_get_v2_data_list_creative_material_mode
const (
	STATIC_ASSEMBLE_ReportAdGetV2DataListCreativeMaterialMode ReportAdGetV2DataListCreativeMaterialMode = "STATIC_ASSEMBLE"
	INTERVENE_ReportAdGetV2DataListCreativeMaterialMode       ReportAdGetV2DataListCreativeMaterialMode = "INTERVENE"
	CTR_ReportAdGetV2DataListCreativeMaterialMode             ReportAdGetV2DataListCreativeMaterialMode = "CTR"
)

// Ptr returns reference to report_ad_get_v2_data_list_creative_material_mode value
func (v ReportAdGetV2DataListCreativeMaterialMode) Ptr() *ReportAdGetV2DataListCreativeMaterialMode {
	return &v
}
