/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsEstimateAudienceV2DeviceType
type ToolsEstimateAudienceV2DeviceType string

// List of tools_estimate_audience_v2_device_type
const (
	PAD_ToolsEstimateAudienceV2DeviceType    ToolsEstimateAudienceV2DeviceType = "PAD"
	MOBILE_ToolsEstimateAudienceV2DeviceType ToolsEstimateAudienceV2DeviceType = "MOBILE"
)

// Ptr returns reference to tools_estimate_audience_v2_device_type value
func (v ToolsEstimateAudienceV2DeviceType) Ptr() *ToolsEstimateAudienceV2DeviceType {
	return &v
}
