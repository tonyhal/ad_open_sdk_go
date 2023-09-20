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

// ToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListSuggestionsToolType
type ToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListSuggestionsToolType string

// List of tools_diagnosis_suggestion_get_v2_data_suggestion_list_scene_list_suggestions_tool_type
const (
	ACTION_ToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListSuggestionsToolType ToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListSuggestionsToolType = "Action"
	TEXT_ToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListSuggestionsToolType   ToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListSuggestionsToolType = "Text"
)

// All allowed values of ToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListSuggestionsToolType enum
var AllowedToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListSuggestionsToolTypeEnumValues = []ToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListSuggestionsToolType{
	"Action",
	"Text",
}

// NewToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListSuggestionsToolTypeFromValue returns a pointer to a valid ToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListSuggestionsToolType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListSuggestionsToolTypeFromValue(v string) (*ToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListSuggestionsToolType, error) {
	ev := ToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListSuggestionsToolType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListSuggestionsToolType: valid values are %v", v, AllowedToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListSuggestionsToolTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListSuggestionsToolType) IsValid() bool {
	for _, existing := range AllowedToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListSuggestionsToolTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_diagnosis_suggestion_get_v2_data_suggestion_list_scene_list_suggestions_tool_type value
func (v ToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListSuggestionsToolType) Ptr() *ToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListSuggestionsToolType {
	return &v
}
