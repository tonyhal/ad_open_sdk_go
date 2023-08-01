/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionCreateV30CreativeAutoGenerateSwitch
type PromotionCreateV30CreativeAutoGenerateSwitch string

// List of promotion_create_v3.0_creative_auto_generate_switch
const (
	OFF_PromotionCreateV30CreativeAutoGenerateSwitch PromotionCreateV30CreativeAutoGenerateSwitch = "OFF"
	ON_PromotionCreateV30CreativeAutoGenerateSwitch  PromotionCreateV30CreativeAutoGenerateSwitch = "ON"
)

// All allowed values of PromotionCreateV30CreativeAutoGenerateSwitch enum
var AllowedPromotionCreateV30CreativeAutoGenerateSwitchEnumValues = []PromotionCreateV30CreativeAutoGenerateSwitch{
	"OFF",
	"ON",
}

// NewPromotionCreateV30CreativeAutoGenerateSwitchFromValue returns a pointer to a valid PromotionCreateV30CreativeAutoGenerateSwitch
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionCreateV30CreativeAutoGenerateSwitchFromValue(v string) (*PromotionCreateV30CreativeAutoGenerateSwitch, error) {
	ev := PromotionCreateV30CreativeAutoGenerateSwitch(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionCreateV30CreativeAutoGenerateSwitch: valid values are %v", v, AllowedPromotionCreateV30CreativeAutoGenerateSwitchEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionCreateV30CreativeAutoGenerateSwitch) IsValid() bool {
	for _, existing := range AllowedPromotionCreateV30CreativeAutoGenerateSwitchEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_create_v3.0_creative_auto_generate_switch value
func (v PromotionCreateV30CreativeAutoGenerateSwitch) Ptr() *PromotionCreateV30CreativeAutoGenerateSwitch {
	return &v
}