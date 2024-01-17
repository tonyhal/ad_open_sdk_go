/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsSource
type AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsSource string

// List of ad_convert_optimized_target_get_v3.0_data_list_converts_external_actions_source
const (
	SOURCE_CONFIG_AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsSource AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsSource = "SOURCE_CONFIG"
	SOURCE_METEOR_AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsSource AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsSource = "SOURCE_METEOR"
	SOURCE_TETRIS_AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsSource AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsSource = "SOURCE_TETRIS"
)

// All allowed values of AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsSource enum
var AllowedAdConvertOptimizedTargetGetV30DataListConvertsExternalActionsSourceEnumValues = []AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsSource{
	"SOURCE_CONFIG",
	"SOURCE_METEOR",
	"SOURCE_TETRIS",
}

// NewAdConvertOptimizedTargetGetV30DataListConvertsExternalActionsSourceFromValue returns a pointer to a valid AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdConvertOptimizedTargetGetV30DataListConvertsExternalActionsSourceFromValue(v string) (*AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsSource, error) {
	ev := AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsSource: valid values are %v", v, AllowedAdConvertOptimizedTargetGetV30DataListConvertsExternalActionsSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsSource) IsValid() bool {
	for _, existing := range AllowedAdConvertOptimizedTargetGetV30DataListConvertsExternalActionsSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_convert_optimized_target_get_v3.0_data_list_converts_external_actions_source value
func (v AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsSource) Ptr() *AdConvertOptimizedTargetGetV30DataListConvertsExternalActionsSource {
	return &v
}