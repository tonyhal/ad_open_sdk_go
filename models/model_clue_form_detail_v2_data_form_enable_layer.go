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

// ClueFormDetailV2DataFormEnableLayer
type ClueFormDetailV2DataFormEnableLayer string

// List of clue_form_detail_v2_data_form_enable_layer
const (
	Enum_0_ClueFormDetailV2DataFormEnableLayer ClueFormDetailV2DataFormEnableLayer = "0"
	Enum_1_ClueFormDetailV2DataFormEnableLayer ClueFormDetailV2DataFormEnableLayer = "1"
)

// All allowed values of ClueFormDetailV2DataFormEnableLayer enum
var AllowedClueFormDetailV2DataFormEnableLayerEnumValues = []ClueFormDetailV2DataFormEnableLayer{
	"0",
	"1",
}

// NewClueFormDetailV2DataFormEnableLayerFromValue returns a pointer to a valid ClueFormDetailV2DataFormEnableLayer
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueFormDetailV2DataFormEnableLayerFromValue(v string) (*ClueFormDetailV2DataFormEnableLayer, error) {
	ev := ClueFormDetailV2DataFormEnableLayer(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueFormDetailV2DataFormEnableLayer: valid values are %v", v, AllowedClueFormDetailV2DataFormEnableLayerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueFormDetailV2DataFormEnableLayer) IsValid() bool {
	for _, existing := range AllowedClueFormDetailV2DataFormEnableLayerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_form_detail_v2_data_form_enable_layer value
func (v ClueFormDetailV2DataFormEnableLayer) Ptr() *ClueFormDetailV2DataFormEnableLayer {
	return &v
}
