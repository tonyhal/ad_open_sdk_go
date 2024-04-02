/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsTaskRaiseCreateV2PlatformVersion
type ToolsTaskRaiseCreateV2PlatformVersion string

// List of tools_task_raise_create_v2_platform_version
const (
	V2_ToolsTaskRaiseCreateV2PlatformVersion ToolsTaskRaiseCreateV2PlatformVersion = "V2"
)

// All allowed values of ToolsTaskRaiseCreateV2PlatformVersion enum
var AllowedToolsTaskRaiseCreateV2PlatformVersionEnumValues = []ToolsTaskRaiseCreateV2PlatformVersion{
	"V2",
}

// NewToolsTaskRaiseCreateV2PlatformVersionFromValue returns a pointer to a valid ToolsTaskRaiseCreateV2PlatformVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsTaskRaiseCreateV2PlatformVersionFromValue(v string) (*ToolsTaskRaiseCreateV2PlatformVersion, error) {
	ev := ToolsTaskRaiseCreateV2PlatformVersion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsTaskRaiseCreateV2PlatformVersion: valid values are %v", v, AllowedToolsTaskRaiseCreateV2PlatformVersionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsTaskRaiseCreateV2PlatformVersion) IsValid() bool {
	for _, existing := range AllowedToolsTaskRaiseCreateV2PlatformVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_task_raise_create_v2_platform_version value
func (v ToolsTaskRaiseCreateV2PlatformVersion) Ptr() *ToolsTaskRaiseCreateV2PlatformVersion {
	return &v
}
