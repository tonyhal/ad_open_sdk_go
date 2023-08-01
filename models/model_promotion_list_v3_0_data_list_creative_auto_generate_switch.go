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

// PromotionListV30DataListCreativeAutoGenerateSwitch
type PromotionListV30DataListCreativeAutoGenerateSwitch string

// List of promotion_list_v3.0_data_list_creative_auto_generate_switch
const (
	OFF_PromotionListV30DataListCreativeAutoGenerateSwitch PromotionListV30DataListCreativeAutoGenerateSwitch = "OFF"
	ON_PromotionListV30DataListCreativeAutoGenerateSwitch  PromotionListV30DataListCreativeAutoGenerateSwitch = "ON"
)

// All allowed values of PromotionListV30DataListCreativeAutoGenerateSwitch enum
var AllowedPromotionListV30DataListCreativeAutoGenerateSwitchEnumValues = []PromotionListV30DataListCreativeAutoGenerateSwitch{
	"OFF",
	"ON",
}

// NewPromotionListV30DataListCreativeAutoGenerateSwitchFromValue returns a pointer to a valid PromotionListV30DataListCreativeAutoGenerateSwitch
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListCreativeAutoGenerateSwitchFromValue(v string) (*PromotionListV30DataListCreativeAutoGenerateSwitch, error) {
	ev := PromotionListV30DataListCreativeAutoGenerateSwitch(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListCreativeAutoGenerateSwitch: valid values are %v", v, AllowedPromotionListV30DataListCreativeAutoGenerateSwitchEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListCreativeAutoGenerateSwitch) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListCreativeAutoGenerateSwitchEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_creative_auto_generate_switch value
func (v PromotionListV30DataListCreativeAutoGenerateSwitch) Ptr() *PromotionListV30DataListCreativeAutoGenerateSwitch {
	return &v
}