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

// CreativeDetailGetV30DataAdDataParamsType
type CreativeDetailGetV30DataAdDataParamsType string

// List of creative_detail_get_v3.0_data_ad_data_params_type
const (
	DEFAULT_CreativeDetailGetV30DataAdDataParamsType CreativeDetailGetV30DataAdDataParamsType = "DEFAULT"
	CUSTOM_CreativeDetailGetV30DataAdDataParamsType  CreativeDetailGetV30DataAdDataParamsType = "CUSTOM"
	DPA_CreativeDetailGetV30DataAdDataParamsType     CreativeDetailGetV30DataAdDataParamsType = "DPA"
)

// All allowed values of CreativeDetailGetV30DataAdDataParamsType enum
var AllowedCreativeDetailGetV30DataAdDataParamsTypeEnumValues = []CreativeDetailGetV30DataAdDataParamsType{
	"DEFAULT",
	"CUSTOM",
	"DPA",
}

// NewCreativeDetailGetV30DataAdDataParamsTypeFromValue returns a pointer to a valid CreativeDetailGetV30DataAdDataParamsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeDetailGetV30DataAdDataParamsTypeFromValue(v string) (*CreativeDetailGetV30DataAdDataParamsType, error) {
	ev := CreativeDetailGetV30DataAdDataParamsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeDetailGetV30DataAdDataParamsType: valid values are %v", v, AllowedCreativeDetailGetV30DataAdDataParamsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeDetailGetV30DataAdDataParamsType) IsValid() bool {
	for _, existing := range AllowedCreativeDetailGetV30DataAdDataParamsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_detail_get_v3.0_data_ad_data_params_type value
func (v CreativeDetailGetV30DataAdDataParamsType) Ptr() *CreativeDetailGetV30DataAdDataParamsType {
	return &v
}
