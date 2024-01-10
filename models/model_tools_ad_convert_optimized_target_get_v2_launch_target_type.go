/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAdConvertOptimizedTargetGetV2LaunchTargetType
type ToolsAdConvertOptimizedTargetGetV2LaunchTargetType string

// List of tools_ad_convert_optimized_target_get_v2_launch_target_type
const (
	LIVE_CONVERT_ToolsAdConvertOptimizedTargetGetV2LaunchTargetType ToolsAdConvertOptimizedTargetGetV2LaunchTargetType = "LIVE_CONVERT"
	EXTERNAL_ToolsAdConvertOptimizedTargetGetV2LaunchTargetType     ToolsAdConvertOptimizedTargetGetV2LaunchTargetType = "EXTERNAL"
	APP_ToolsAdConvertOptimizedTargetGetV2LaunchTargetType          ToolsAdConvertOptimizedTargetGetV2LaunchTargetType = "APP"
)

// All allowed values of ToolsAdConvertOptimizedTargetGetV2LaunchTargetType enum
var AllowedToolsAdConvertOptimizedTargetGetV2LaunchTargetTypeEnumValues = []ToolsAdConvertOptimizedTargetGetV2LaunchTargetType{
	"LIVE_CONVERT",
	"EXTERNAL",
	"APP",
}

// NewToolsAdConvertOptimizedTargetGetV2LaunchTargetTypeFromValue returns a pointer to a valid ToolsAdConvertOptimizedTargetGetV2LaunchTargetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdConvertOptimizedTargetGetV2LaunchTargetTypeFromValue(v string) (*ToolsAdConvertOptimizedTargetGetV2LaunchTargetType, error) {
	ev := ToolsAdConvertOptimizedTargetGetV2LaunchTargetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdConvertOptimizedTargetGetV2LaunchTargetType: valid values are %v", v, AllowedToolsAdConvertOptimizedTargetGetV2LaunchTargetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdConvertOptimizedTargetGetV2LaunchTargetType) IsValid() bool {
	for _, existing := range AllowedToolsAdConvertOptimizedTargetGetV2LaunchTargetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ad_convert_optimized_target_get_v2_launch_target_type value
func (v ToolsAdConvertOptimizedTargetGetV2LaunchTargetType) Ptr() *ToolsAdConvertOptimizedTargetGetV2LaunchTargetType {
	return &v
}
