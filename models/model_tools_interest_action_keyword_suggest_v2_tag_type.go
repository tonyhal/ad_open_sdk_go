/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsInterestActionKeywordSuggestV2TagType
type ToolsInterestActionKeywordSuggestV2TagType string

// List of tools_interest_action_keyword_suggest_v2_tag_type
const (
	CATEGORY_ToolsInterestActionKeywordSuggestV2TagType ToolsInterestActionKeywordSuggestV2TagType = "CATEGORY"
	KEYWORD_ToolsInterestActionKeywordSuggestV2TagType  ToolsInterestActionKeywordSuggestV2TagType = "KEYWORD"
)

// All allowed values of ToolsInterestActionKeywordSuggestV2TagType enum
var AllowedToolsInterestActionKeywordSuggestV2TagTypeEnumValues = []ToolsInterestActionKeywordSuggestV2TagType{
	"CATEGORY",
	"KEYWORD",
}

// NewToolsInterestActionKeywordSuggestV2TagTypeFromValue returns a pointer to a valid ToolsInterestActionKeywordSuggestV2TagType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsInterestActionKeywordSuggestV2TagTypeFromValue(v string) (*ToolsInterestActionKeywordSuggestV2TagType, error) {
	ev := ToolsInterestActionKeywordSuggestV2TagType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsInterestActionKeywordSuggestV2TagType: valid values are %v", v, AllowedToolsInterestActionKeywordSuggestV2TagTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsInterestActionKeywordSuggestV2TagType) IsValid() bool {
	for _, existing := range AllowedToolsInterestActionKeywordSuggestV2TagTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_interest_action_keyword_suggest_v2_tag_type value
func (v ToolsInterestActionKeywordSuggestV2TagType) Ptr() *ToolsInterestActionKeywordSuggestV2TagType {
	return &v
}
