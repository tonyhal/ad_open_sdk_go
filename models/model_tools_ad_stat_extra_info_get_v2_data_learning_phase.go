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

// ToolsAdStatExtraInfoGetV2DataLearningPhase
type ToolsAdStatExtraInfoGetV2DataLearningPhase string

// List of tools_ad_stat_extra_info_get_v2_data_learning_phase
const (
	LEARN_FAILED_ToolsAdStatExtraInfoGetV2DataLearningPhase ToolsAdStatExtraInfoGetV2DataLearningPhase = "LEARN_FAILED"
	LEARNED_ToolsAdStatExtraInfoGetV2DataLearningPhase      ToolsAdStatExtraInfoGetV2DataLearningPhase = "LEARNED"
	LEARNING_ToolsAdStatExtraInfoGetV2DataLearningPhase     ToolsAdStatExtraInfoGetV2DataLearningPhase = "LEARNING"
	DEFAULT_ToolsAdStatExtraInfoGetV2DataLearningPhase      ToolsAdStatExtraInfoGetV2DataLearningPhase = "DEFAULT"
)

// All allowed values of ToolsAdStatExtraInfoGetV2DataLearningPhase enum
var AllowedToolsAdStatExtraInfoGetV2DataLearningPhaseEnumValues = []ToolsAdStatExtraInfoGetV2DataLearningPhase{
	"LEARN_FAILED",
	"LEARNED",
	"LEARNING",
	"DEFAULT",
}

// NewToolsAdStatExtraInfoGetV2DataLearningPhaseFromValue returns a pointer to a valid ToolsAdStatExtraInfoGetV2DataLearningPhase
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdStatExtraInfoGetV2DataLearningPhaseFromValue(v string) (*ToolsAdStatExtraInfoGetV2DataLearningPhase, error) {
	ev := ToolsAdStatExtraInfoGetV2DataLearningPhase(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdStatExtraInfoGetV2DataLearningPhase: valid values are %v", v, AllowedToolsAdStatExtraInfoGetV2DataLearningPhaseEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdStatExtraInfoGetV2DataLearningPhase) IsValid() bool {
	for _, existing := range AllowedToolsAdStatExtraInfoGetV2DataLearningPhaseEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ad_stat_extra_info_get_v2_data_learning_phase value
func (v ToolsAdStatExtraInfoGetV2DataLearningPhase) Ptr() *ToolsAdStatExtraInfoGetV2DataLearningPhase {
	return &v
}
