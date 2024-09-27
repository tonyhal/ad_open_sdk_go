/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsEstimateAudienceV2District
type ToolsEstimateAudienceV2District string

// List of tools_estimate_audience_v2_district
const (
	BUSINESS_DISTRICT_ToolsEstimateAudienceV2District ToolsEstimateAudienceV2District = "BUSINESS_DISTRICT"
	NONE_ToolsEstimateAudienceV2District              ToolsEstimateAudienceV2District = "NONE"
	REGION_ToolsEstimateAudienceV2District            ToolsEstimateAudienceV2District = "REGION"
	CITY_ToolsEstimateAudienceV2District              ToolsEstimateAudienceV2District = "CITY"
	COUNTY_ToolsEstimateAudienceV2District            ToolsEstimateAudienceV2District = "COUNTY"
	OVERSEA_ToolsEstimateAudienceV2District           ToolsEstimateAudienceV2District = "OVERSEA"
)

// Ptr returns reference to tools_estimate_audience_v2_district value
func (v ToolsEstimateAudienceV2District) Ptr() *ToolsEstimateAudienceV2District {
	return &v
}
