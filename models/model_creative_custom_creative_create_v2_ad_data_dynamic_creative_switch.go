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

// CreativeCustomCreativeCreateV2AdDataDynamicCreativeSwitch
type CreativeCustomCreativeCreateV2AdDataDynamicCreativeSwitch string

// List of creative_custom_creative_create_v2_ad_data_dynamic_creative_switch
const (
	DYNAMIC_CREATIVE_ON_CreativeCustomCreativeCreateV2AdDataDynamicCreativeSwitch       CreativeCustomCreativeCreateV2AdDataDynamicCreativeSwitch = "DYNAMIC_CREATIVE_ON"
	DYNAMIC_CREATIVE_ABSTRACT_CreativeCustomCreativeCreateV2AdDataDynamicCreativeSwitch CreativeCustomCreativeCreateV2AdDataDynamicCreativeSwitch = "DYNAMIC_CREATIVE_ABSTRACT"
	DYNAMIC_CREATIVE_SUBLINK_CreativeCustomCreativeCreateV2AdDataDynamicCreativeSwitch  CreativeCustomCreativeCreateV2AdDataDynamicCreativeSwitch = "DYNAMIC_CREATIVE_SUBLINK"
	DYNAMIC_CREATIVE_TITLE_CreativeCustomCreativeCreateV2AdDataDynamicCreativeSwitch    CreativeCustomCreativeCreateV2AdDataDynamicCreativeSwitch = "DYNAMIC_CREATIVE_TITLE"
)

// All allowed values of CreativeCustomCreativeCreateV2AdDataDynamicCreativeSwitch enum
var AllowedCreativeCustomCreativeCreateV2AdDataDynamicCreativeSwitchEnumValues = []CreativeCustomCreativeCreateV2AdDataDynamicCreativeSwitch{
	"DYNAMIC_CREATIVE_ON",
	"DYNAMIC_CREATIVE_ABSTRACT",
	"DYNAMIC_CREATIVE_SUBLINK",
	"DYNAMIC_CREATIVE_TITLE",
}

// NewCreativeCustomCreativeCreateV2AdDataDynamicCreativeSwitchFromValue returns a pointer to a valid CreativeCustomCreativeCreateV2AdDataDynamicCreativeSwitch
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeCustomCreativeCreateV2AdDataDynamicCreativeSwitchFromValue(v string) (*CreativeCustomCreativeCreateV2AdDataDynamicCreativeSwitch, error) {
	ev := CreativeCustomCreativeCreateV2AdDataDynamicCreativeSwitch(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeCustomCreativeCreateV2AdDataDynamicCreativeSwitch: valid values are %v", v, AllowedCreativeCustomCreativeCreateV2AdDataDynamicCreativeSwitchEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeCustomCreativeCreateV2AdDataDynamicCreativeSwitch) IsValid() bool {
	for _, existing := range AllowedCreativeCustomCreativeCreateV2AdDataDynamicCreativeSwitchEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_custom_creative_create_v2_ad_data_dynamic_creative_switch value
func (v CreativeCustomCreativeCreateV2AdDataDynamicCreativeSwitch) Ptr() *CreativeCustomCreativeCreateV2AdDataDynamicCreativeSwitch {
	return &v
}
