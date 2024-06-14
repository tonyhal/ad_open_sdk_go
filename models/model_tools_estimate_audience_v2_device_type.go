/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsEstimateAudienceV2DeviceType
type ToolsEstimateAudienceV2DeviceType string

// List of tools_estimate_audience_v2_device_type
const (
	MOBILE_ToolsEstimateAudienceV2DeviceType ToolsEstimateAudienceV2DeviceType = "MOBILE"
	PAD_ToolsEstimateAudienceV2DeviceType    ToolsEstimateAudienceV2DeviceType = "PAD"
)

// All allowed values of ToolsEstimateAudienceV2DeviceType enum
var AllowedToolsEstimateAudienceV2DeviceTypeEnumValues = []ToolsEstimateAudienceV2DeviceType{
	"MOBILE",
	"PAD",
}

// NewToolsEstimateAudienceV2DeviceTypeFromValue returns a pointer to a valid ToolsEstimateAudienceV2DeviceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEstimateAudienceV2DeviceTypeFromValue(v string) (*ToolsEstimateAudienceV2DeviceType, error) {
	ev := ToolsEstimateAudienceV2DeviceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEstimateAudienceV2DeviceType: valid values are %v", v, AllowedToolsEstimateAudienceV2DeviceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEstimateAudienceV2DeviceType) IsValid() bool {
	for _, existing := range AllowedToolsEstimateAudienceV2DeviceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_estimate_audience_v2_device_type value
func (v ToolsEstimateAudienceV2DeviceType) Ptr() *ToolsEstimateAudienceV2DeviceType {
	return &v
}
