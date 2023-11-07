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

// ToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListSuggestionsToolType
type ToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListSuggestionsToolType string

// List of tools_promotion_diagnosis_suggestion_get_v3.0_data_suggestion_list_scene_list_suggestions_tool_type
const (
	ACTION_ToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListSuggestionsToolType ToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListSuggestionsToolType = "ACTION"
	TEXT_ToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListSuggestionsToolType   ToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListSuggestionsToolType = "TEXT"
)

// All allowed values of ToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListSuggestionsToolType enum
var AllowedToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListSuggestionsToolTypeEnumValues = []ToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListSuggestionsToolType{
	"ACTION",
	"TEXT",
}

// NewToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListSuggestionsToolTypeFromValue returns a pointer to a valid ToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListSuggestionsToolType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListSuggestionsToolTypeFromValue(v string) (*ToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListSuggestionsToolType, error) {
	ev := ToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListSuggestionsToolType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListSuggestionsToolType: valid values are %v", v, AllowedToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListSuggestionsToolTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListSuggestionsToolType) IsValid() bool {
	for _, existing := range AllowedToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListSuggestionsToolTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_promotion_diagnosis_suggestion_get_v3.0_data_suggestion_list_scene_list_suggestions_tool_type value
func (v ToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListSuggestionsToolType) Ptr() *ToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListSuggestionsToolType {
	return &v
}
