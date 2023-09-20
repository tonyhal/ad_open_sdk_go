/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsInterestActionId2wordV2TagType
type ToolsInterestActionId2wordV2TagType string

// List of tools_interest_action_id2word_v2_tag_type
const (
	CATEGORY_ToolsInterestActionId2wordV2TagType ToolsInterestActionId2wordV2TagType = "CATEGORY"
	KEYWORD_ToolsInterestActionId2wordV2TagType  ToolsInterestActionId2wordV2TagType = "KEYWORD"
)

// All allowed values of ToolsInterestActionId2wordV2TagType enum
var AllowedToolsInterestActionId2wordV2TagTypeEnumValues = []ToolsInterestActionId2wordV2TagType{
	"CATEGORY",
	"KEYWORD",
}

// NewToolsInterestActionId2wordV2TagTypeFromValue returns a pointer to a valid ToolsInterestActionId2wordV2TagType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsInterestActionId2wordV2TagTypeFromValue(v string) (*ToolsInterestActionId2wordV2TagType, error) {
	ev := ToolsInterestActionId2wordV2TagType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsInterestActionId2wordV2TagType: valid values are %v", v, AllowedToolsInterestActionId2wordV2TagTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsInterestActionId2wordV2TagType) IsValid() bool {
	for _, existing := range AllowedToolsInterestActionId2wordV2TagTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_interest_action_id2word_v2_tag_type value
func (v ToolsInterestActionId2wordV2TagType) Ptr() *ToolsInterestActionId2wordV2TagType {
	return &v
}