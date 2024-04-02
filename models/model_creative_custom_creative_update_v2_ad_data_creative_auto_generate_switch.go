/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeCustomCreativeUpdateV2AdDataCreativeAutoGenerateSwitch
type CreativeCustomCreativeUpdateV2AdDataCreativeAutoGenerateSwitch int64

// List of creative_custom_creative_update_v2_ad_data_creative_auto_generate_switch
const (
	Enum_0_CreativeCustomCreativeUpdateV2AdDataCreativeAutoGenerateSwitch CreativeCustomCreativeUpdateV2AdDataCreativeAutoGenerateSwitch = 0
	Enum_1_CreativeCustomCreativeUpdateV2AdDataCreativeAutoGenerateSwitch CreativeCustomCreativeUpdateV2AdDataCreativeAutoGenerateSwitch = 1
)

// All allowed values of CreativeCustomCreativeUpdateV2AdDataCreativeAutoGenerateSwitch enum
var AllowedCreativeCustomCreativeUpdateV2AdDataCreativeAutoGenerateSwitchEnumValues = []CreativeCustomCreativeUpdateV2AdDataCreativeAutoGenerateSwitch{
	0,
	1,
}

// NewCreativeCustomCreativeUpdateV2AdDataCreativeAutoGenerateSwitchFromValue returns a pointer to a valid CreativeCustomCreativeUpdateV2AdDataCreativeAutoGenerateSwitch
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeCustomCreativeUpdateV2AdDataCreativeAutoGenerateSwitchFromValue(v int64) (*CreativeCustomCreativeUpdateV2AdDataCreativeAutoGenerateSwitch, error) {
	ev := CreativeCustomCreativeUpdateV2AdDataCreativeAutoGenerateSwitch(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeCustomCreativeUpdateV2AdDataCreativeAutoGenerateSwitch: valid values are %v", v, AllowedCreativeCustomCreativeUpdateV2AdDataCreativeAutoGenerateSwitchEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeCustomCreativeUpdateV2AdDataCreativeAutoGenerateSwitch) IsValid() bool {
	for _, existing := range AllowedCreativeCustomCreativeUpdateV2AdDataCreativeAutoGenerateSwitchEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_custom_creative_update_v2_ad_data_creative_auto_generate_switch value
func (v CreativeCustomCreativeUpdateV2AdDataCreativeAutoGenerateSwitch) Ptr() *CreativeCustomCreativeUpdateV2AdDataCreativeAutoGenerateSwitch {
	return &v
}
