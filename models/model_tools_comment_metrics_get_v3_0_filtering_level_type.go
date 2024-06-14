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

// ToolsCommentMetricsGetV30FilteringLevelType
type ToolsCommentMetricsGetV30FilteringLevelType string

// List of tools_comment_metrics_get_v3.0_filtering_level_type
const (
	LEVEL_ALL_ToolsCommentMetricsGetV30FilteringLevelType ToolsCommentMetricsGetV30FilteringLevelType = "LEVEL_ALL"
	LEVEL_ONE_ToolsCommentMetricsGetV30FilteringLevelType ToolsCommentMetricsGetV30FilteringLevelType = "LEVEL_ONE"
	LEVEL_TWO_ToolsCommentMetricsGetV30FilteringLevelType ToolsCommentMetricsGetV30FilteringLevelType = "LEVEL_TWO"
)

// All allowed values of ToolsCommentMetricsGetV30FilteringLevelType enum
var AllowedToolsCommentMetricsGetV30FilteringLevelTypeEnumValues = []ToolsCommentMetricsGetV30FilteringLevelType{
	"LEVEL_ALL",
	"LEVEL_ONE",
	"LEVEL_TWO",
}

// NewToolsCommentMetricsGetV30FilteringLevelTypeFromValue returns a pointer to a valid ToolsCommentMetricsGetV30FilteringLevelType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsCommentMetricsGetV30FilteringLevelTypeFromValue(v string) (*ToolsCommentMetricsGetV30FilteringLevelType, error) {
	ev := ToolsCommentMetricsGetV30FilteringLevelType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsCommentMetricsGetV30FilteringLevelType: valid values are %v", v, AllowedToolsCommentMetricsGetV30FilteringLevelTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsCommentMetricsGetV30FilteringLevelType) IsValid() bool {
	for _, existing := range AllowedToolsCommentMetricsGetV30FilteringLevelTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_comment_metrics_get_v3.0_filtering_level_type value
func (v ToolsCommentMetricsGetV30FilteringLevelType) Ptr() *ToolsCommentMetricsGetV30FilteringLevelType {
	return &v
}
