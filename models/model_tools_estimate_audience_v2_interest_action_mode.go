/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsEstimateAudienceV2InterestActionMode
type ToolsEstimateAudienceV2InterestActionMode string

// List of tools_estimate_audience_v2_interest_action_mode
const (
	RECOMMEND_ToolsEstimateAudienceV2InterestActionMode ToolsEstimateAudienceV2InterestActionMode = "RECOMMEND"
	UNLIMITED_ToolsEstimateAudienceV2InterestActionMode ToolsEstimateAudienceV2InterestActionMode = "UNLIMITED"
	CUSTOM_ToolsEstimateAudienceV2InterestActionMode    ToolsEstimateAudienceV2InterestActionMode = "CUSTOM"
)

// Ptr returns reference to tools_estimate_audience_v2_interest_action_mode value
func (v ToolsEstimateAudienceV2InterestActionMode) Ptr() *ToolsEstimateAudienceV2InterestActionMode {
	return &v
}
