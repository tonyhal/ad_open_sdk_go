/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCreativeGetV2FilteringCreativeMaterialModes
type ReportCreativeGetV2FilteringCreativeMaterialModes string

// List of report_creative_get_v2_filtering_creative_material_modes
const (
	STATIC_ASSEMBLE_ReportCreativeGetV2FilteringCreativeMaterialModes ReportCreativeGetV2FilteringCreativeMaterialModes = "STATIC_ASSEMBLE"
	INTERVENE_ReportCreativeGetV2FilteringCreativeMaterialModes       ReportCreativeGetV2FilteringCreativeMaterialModes = "INTERVENE"
	CTR_ReportCreativeGetV2FilteringCreativeMaterialModes             ReportCreativeGetV2FilteringCreativeMaterialModes = "CTR"
)

// Ptr returns reference to report_creative_get_v2_filtering_creative_material_modes value
func (v ReportCreativeGetV2FilteringCreativeMaterialModes) Ptr() *ReportCreativeGetV2FilteringCreativeMaterialModes {
	return &v
}
