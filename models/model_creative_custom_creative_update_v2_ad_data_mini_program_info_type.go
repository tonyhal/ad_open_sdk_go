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

// CreativeCustomCreativeUpdateV2AdDataMiniProgramInfoType
type CreativeCustomCreativeUpdateV2AdDataMiniProgramInfoType string

// List of creative_custom_creative_update_v2_ad_data_mini_program_info_type
const (
	SHELL_APP_CreativeCustomCreativeUpdateV2AdDataMiniProgramInfoType    CreativeCustomCreativeUpdateV2AdDataMiniProgramInfoType = "SHELL_APP"
	BYTE_GAME_CreativeCustomCreativeUpdateV2AdDataMiniProgramInfoType    CreativeCustomCreativeUpdateV2AdDataMiniProgramInfoType = "BYTE_GAME"
	BYTE_APP_CreativeCustomCreativeUpdateV2AdDataMiniProgramInfoType     CreativeCustomCreativeUpdateV2AdDataMiniProgramInfoType = "BYTE_APP"
	TEMPLATE_APP_CreativeCustomCreativeUpdateV2AdDataMiniProgramInfoType CreativeCustomCreativeUpdateV2AdDataMiniProgramInfoType = "TEMPLATE_APP"
)

// All allowed values of CreativeCustomCreativeUpdateV2AdDataMiniProgramInfoType enum
var AllowedCreativeCustomCreativeUpdateV2AdDataMiniProgramInfoTypeEnumValues = []CreativeCustomCreativeUpdateV2AdDataMiniProgramInfoType{
	"SHELL_APP",
	"BYTE_GAME",
	"BYTE_APP",
	"TEMPLATE_APP",
}

// NewCreativeCustomCreativeUpdateV2AdDataMiniProgramInfoTypeFromValue returns a pointer to a valid CreativeCustomCreativeUpdateV2AdDataMiniProgramInfoType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeCustomCreativeUpdateV2AdDataMiniProgramInfoTypeFromValue(v string) (*CreativeCustomCreativeUpdateV2AdDataMiniProgramInfoType, error) {
	ev := CreativeCustomCreativeUpdateV2AdDataMiniProgramInfoType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeCustomCreativeUpdateV2AdDataMiniProgramInfoType: valid values are %v", v, AllowedCreativeCustomCreativeUpdateV2AdDataMiniProgramInfoTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeCustomCreativeUpdateV2AdDataMiniProgramInfoType) IsValid() bool {
	for _, existing := range AllowedCreativeCustomCreativeUpdateV2AdDataMiniProgramInfoTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_custom_creative_update_v2_ad_data_mini_program_info_type value
func (v CreativeCustomCreativeUpdateV2AdDataMiniProgramInfoType) Ptr() *CreativeCustomCreativeUpdateV2AdDataMiniProgramInfoType {
	return &v
}
