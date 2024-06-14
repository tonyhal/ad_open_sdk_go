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

// ToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListScene
type ToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListScene string

// List of tools_diagnosis_suggestion_get_v2_data_suggestion_list_scene_list_scene
const (
	CLEAN_ToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListScene     ToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListScene = "CLEAN"
	POTENTIAL_ToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListScene ToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListScene = "POTENTIAL"
	PROBLEM_ToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListScene   ToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListScene = "PROBLEM"
)

// All allowed values of ToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListScene enum
var AllowedToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListSceneEnumValues = []ToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListScene{
	"CLEAN",
	"POTENTIAL",
	"PROBLEM",
}

// NewToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListSceneFromValue returns a pointer to a valid ToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListScene
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListSceneFromValue(v string) (*ToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListScene, error) {
	ev := ToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListScene(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListScene: valid values are %v", v, AllowedToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListSceneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListScene) IsValid() bool {
	for _, existing := range AllowedToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListSceneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_diagnosis_suggestion_get_v2_data_suggestion_list_scene_list_scene value
func (v ToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListScene) Ptr() *ToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListScene {
	return &v
}
