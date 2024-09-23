/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsEstimateAudienceV2AutoExtendTargets
type ToolsEstimateAudienceV2AutoExtendTargets string

// List of tools_estimate_audience_v2_auto_extend_targets
const (
	CUSTOM_AUDIENCE_ToolsEstimateAudienceV2AutoExtendTargets ToolsEstimateAudienceV2AutoExtendTargets = "CUSTOM_AUDIENCE"
	REGION_ToolsEstimateAudienceV2AutoExtendTargets          ToolsEstimateAudienceV2AutoExtendTargets = "REGION"
	INTEREST_TAG_ToolsEstimateAudienceV2AutoExtendTargets    ToolsEstimateAudienceV2AutoExtendTargets = "INTEREST_TAG"
	AD_TAG_ToolsEstimateAudienceV2AutoExtendTargets          ToolsEstimateAudienceV2AutoExtendTargets = "AD_TAG"
	GENDER_ToolsEstimateAudienceV2AutoExtendTargets          ToolsEstimateAudienceV2AutoExtendTargets = "GENDER"
	AGE_ToolsEstimateAudienceV2AutoExtendTargets             ToolsEstimateAudienceV2AutoExtendTargets = "AGE"
)

// Ptr returns reference to tools_estimate_audience_v2_auto_extend_targets value
func (v ToolsEstimateAudienceV2AutoExtendTargets) Ptr() *ToolsEstimateAudienceV2AutoExtendTargets {
	return &v
}
