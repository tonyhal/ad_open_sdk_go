/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsTaskRaiseOptimizationIdsGetV2PlatformVersion
type ToolsTaskRaiseOptimizationIdsGetV2PlatformVersion string

// List of tools_task_raise_optimization_ids_get_v2_platform_version
const (
	V2_ToolsTaskRaiseOptimizationIdsGetV2PlatformVersion ToolsTaskRaiseOptimizationIdsGetV2PlatformVersion = "V2"
)

// All allowed values of ToolsTaskRaiseOptimizationIdsGetV2PlatformVersion enum
var AllowedToolsTaskRaiseOptimizationIdsGetV2PlatformVersionEnumValues = []ToolsTaskRaiseOptimizationIdsGetV2PlatformVersion{
	"V2",
}

// NewToolsTaskRaiseOptimizationIdsGetV2PlatformVersionFromValue returns a pointer to a valid ToolsTaskRaiseOptimizationIdsGetV2PlatformVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsTaskRaiseOptimizationIdsGetV2PlatformVersionFromValue(v string) (*ToolsTaskRaiseOptimizationIdsGetV2PlatformVersion, error) {
	ev := ToolsTaskRaiseOptimizationIdsGetV2PlatformVersion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsTaskRaiseOptimizationIdsGetV2PlatformVersion: valid values are %v", v, AllowedToolsTaskRaiseOptimizationIdsGetV2PlatformVersionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsTaskRaiseOptimizationIdsGetV2PlatformVersion) IsValid() bool {
	for _, existing := range AllowedToolsTaskRaiseOptimizationIdsGetV2PlatformVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_task_raise_optimization_ids_get_v2_platform_version value
func (v ToolsTaskRaiseOptimizationIdsGetV2PlatformVersion) Ptr() *ToolsTaskRaiseOptimizationIdsGetV2PlatformVersion {
	return &v
}
