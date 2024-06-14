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

// ToolsAdConvertOptimizedTargetGetV2DedicateType
type ToolsAdConvertOptimizedTargetGetV2DedicateType string

// List of tools_ad_convert_optimized_target_get_v2_dedicate_type
const (
	DEDICATED_ToolsAdConvertOptimizedTargetGetV2DedicateType ToolsAdConvertOptimizedTargetGetV2DedicateType = "DEDICATED"
	UNSET_ToolsAdConvertOptimizedTargetGetV2DedicateType     ToolsAdConvertOptimizedTargetGetV2DedicateType = "UNSET"
)

// All allowed values of ToolsAdConvertOptimizedTargetGetV2DedicateType enum
var AllowedToolsAdConvertOptimizedTargetGetV2DedicateTypeEnumValues = []ToolsAdConvertOptimizedTargetGetV2DedicateType{
	"DEDICATED",
	"UNSET",
}

// NewToolsAdConvertOptimizedTargetGetV2DedicateTypeFromValue returns a pointer to a valid ToolsAdConvertOptimizedTargetGetV2DedicateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdConvertOptimizedTargetGetV2DedicateTypeFromValue(v string) (*ToolsAdConvertOptimizedTargetGetV2DedicateType, error) {
	ev := ToolsAdConvertOptimizedTargetGetV2DedicateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdConvertOptimizedTargetGetV2DedicateType: valid values are %v", v, AllowedToolsAdConvertOptimizedTargetGetV2DedicateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdConvertOptimizedTargetGetV2DedicateType) IsValid() bool {
	for _, existing := range AllowedToolsAdConvertOptimizedTargetGetV2DedicateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ad_convert_optimized_target_get_v2_dedicate_type value
func (v ToolsAdConvertOptimizedTargetGetV2DedicateType) Ptr() *ToolsAdConvertOptimizedTargetGetV2DedicateType {
	return &v
}
