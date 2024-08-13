/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeCustomCreativeUpdateV2AdDataParamsType
type CreativeCustomCreativeUpdateV2AdDataParamsType string

// List of creative_custom_creative_update_v2_ad_data_params_type
const (
	DPA_CreativeCustomCreativeUpdateV2AdDataParamsType    CreativeCustomCreativeUpdateV2AdDataParamsType = "DPA"
	CUSTOM_CreativeCustomCreativeUpdateV2AdDataParamsType CreativeCustomCreativeUpdateV2AdDataParamsType = "CUSTOM"
)

// All allowed values of CreativeCustomCreativeUpdateV2AdDataParamsType enum
var AllowedCreativeCustomCreativeUpdateV2AdDataParamsTypeEnumValues = []CreativeCustomCreativeUpdateV2AdDataParamsType{
	"DPA",
	"CUSTOM",
}

// NewCreativeCustomCreativeUpdateV2AdDataParamsTypeFromValue returns a pointer to a valid CreativeCustomCreativeUpdateV2AdDataParamsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeCustomCreativeUpdateV2AdDataParamsTypeFromValue(v string) (*CreativeCustomCreativeUpdateV2AdDataParamsType, error) {
	ev := CreativeCustomCreativeUpdateV2AdDataParamsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeCustomCreativeUpdateV2AdDataParamsType: valid values are %v", v, AllowedCreativeCustomCreativeUpdateV2AdDataParamsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeCustomCreativeUpdateV2AdDataParamsType) IsValid() bool {
	for _, existing := range AllowedCreativeCustomCreativeUpdateV2AdDataParamsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_custom_creative_update_v2_ad_data_params_type value
func (v CreativeCustomCreativeUpdateV2AdDataParamsType) Ptr() *CreativeCustomCreativeUpdateV2AdDataParamsType {
	return &v
}
