/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeProceduralCreativeUpdateV2AdDataDynamicCreativeSwitch
type CreativeProceduralCreativeUpdateV2AdDataDynamicCreativeSwitch string

// List of creative_procedural_creative_update_v2_ad_data_dynamic_creative_switch
const (
	DYNAMIC_CREATIVE_SUBLINK_CreativeProceduralCreativeUpdateV2AdDataDynamicCreativeSwitch  CreativeProceduralCreativeUpdateV2AdDataDynamicCreativeSwitch = "DYNAMIC_CREATIVE_SUBLINK"
	DYNAMIC_CREATIVE_ON_CreativeProceduralCreativeUpdateV2AdDataDynamicCreativeSwitch       CreativeProceduralCreativeUpdateV2AdDataDynamicCreativeSwitch = "DYNAMIC_CREATIVE_ON"
	DYNAMIC_CREATIVE_ABSTRACT_CreativeProceduralCreativeUpdateV2AdDataDynamicCreativeSwitch CreativeProceduralCreativeUpdateV2AdDataDynamicCreativeSwitch = "DYNAMIC_CREATIVE_ABSTRACT"
	DYNAMIC_CREATIVE_TITLE_CreativeProceduralCreativeUpdateV2AdDataDynamicCreativeSwitch    CreativeProceduralCreativeUpdateV2AdDataDynamicCreativeSwitch = "DYNAMIC_CREATIVE_TITLE"
)

// All allowed values of CreativeProceduralCreativeUpdateV2AdDataDynamicCreativeSwitch enum
var AllowedCreativeProceduralCreativeUpdateV2AdDataDynamicCreativeSwitchEnumValues = []CreativeProceduralCreativeUpdateV2AdDataDynamicCreativeSwitch{
	"DYNAMIC_CREATIVE_SUBLINK",
	"DYNAMIC_CREATIVE_ON",
	"DYNAMIC_CREATIVE_ABSTRACT",
	"DYNAMIC_CREATIVE_TITLE",
}

// NewCreativeProceduralCreativeUpdateV2AdDataDynamicCreativeSwitchFromValue returns a pointer to a valid CreativeProceduralCreativeUpdateV2AdDataDynamicCreativeSwitch
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeProceduralCreativeUpdateV2AdDataDynamicCreativeSwitchFromValue(v string) (*CreativeProceduralCreativeUpdateV2AdDataDynamicCreativeSwitch, error) {
	ev := CreativeProceduralCreativeUpdateV2AdDataDynamicCreativeSwitch(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeProceduralCreativeUpdateV2AdDataDynamicCreativeSwitch: valid values are %v", v, AllowedCreativeProceduralCreativeUpdateV2AdDataDynamicCreativeSwitchEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeProceduralCreativeUpdateV2AdDataDynamicCreativeSwitch) IsValid() bool {
	for _, existing := range AllowedCreativeProceduralCreativeUpdateV2AdDataDynamicCreativeSwitchEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_procedural_creative_update_v2_ad_data_dynamic_creative_switch value
func (v CreativeProceduralCreativeUpdateV2AdDataDynamicCreativeSwitch) Ptr() *CreativeProceduralCreativeUpdateV2AdDataDynamicCreativeSwitch {
	return &v
}
