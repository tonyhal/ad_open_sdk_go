/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeProceduralCreativeCreateV2AdDataCreativeAutoGenerateSwitch
type CreativeProceduralCreativeCreateV2AdDataCreativeAutoGenerateSwitch int64

// List of creative_procedural_creative_create_v2_ad_data_creative_auto_generate_switch
const (
	Enum_0_CreativeProceduralCreativeCreateV2AdDataCreativeAutoGenerateSwitch CreativeProceduralCreativeCreateV2AdDataCreativeAutoGenerateSwitch = 0
	Enum_1_CreativeProceduralCreativeCreateV2AdDataCreativeAutoGenerateSwitch CreativeProceduralCreativeCreateV2AdDataCreativeAutoGenerateSwitch = 1
)

// All allowed values of CreativeProceduralCreativeCreateV2AdDataCreativeAutoGenerateSwitch enum
var AllowedCreativeProceduralCreativeCreateV2AdDataCreativeAutoGenerateSwitchEnumValues = []CreativeProceduralCreativeCreateV2AdDataCreativeAutoGenerateSwitch{
	0,
	1,
}

// NewCreativeProceduralCreativeCreateV2AdDataCreativeAutoGenerateSwitchFromValue returns a pointer to a valid CreativeProceduralCreativeCreateV2AdDataCreativeAutoGenerateSwitch
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeProceduralCreativeCreateV2AdDataCreativeAutoGenerateSwitchFromValue(v int64) (*CreativeProceduralCreativeCreateV2AdDataCreativeAutoGenerateSwitch, error) {
	ev := CreativeProceduralCreativeCreateV2AdDataCreativeAutoGenerateSwitch(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeProceduralCreativeCreateV2AdDataCreativeAutoGenerateSwitch: valid values are %v", v, AllowedCreativeProceduralCreativeCreateV2AdDataCreativeAutoGenerateSwitchEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeProceduralCreativeCreateV2AdDataCreativeAutoGenerateSwitch) IsValid() bool {
	for _, existing := range AllowedCreativeProceduralCreativeCreateV2AdDataCreativeAutoGenerateSwitchEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_procedural_creative_create_v2_ad_data_creative_auto_generate_switch value
func (v CreativeProceduralCreativeCreateV2AdDataCreativeAutoGenerateSwitch) Ptr() *CreativeProceduralCreativeCreateV2AdDataCreativeAutoGenerateSwitch {
	return &v
}
