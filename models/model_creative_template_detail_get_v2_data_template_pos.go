/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeTemplateDetailGetV2DataTemplatePos
type CreativeTemplateDetailGetV2DataTemplatePos string

// List of creative_template_detail_get_v2_data_template_pos
const (
	CENTER_CreativeTemplateDetailGetV2DataTemplatePos CreativeTemplateDetailGetV2DataTemplatePos = "Center"
	CUSTOM_CreativeTemplateDetailGetV2DataTemplatePos CreativeTemplateDetailGetV2DataTemplatePos = "Custom"
)

// All allowed values of CreativeTemplateDetailGetV2DataTemplatePos enum
var AllowedCreativeTemplateDetailGetV2DataTemplatePosEnumValues = []CreativeTemplateDetailGetV2DataTemplatePos{
	"Center",
	"Custom",
}

// NewCreativeTemplateDetailGetV2DataTemplatePosFromValue returns a pointer to a valid CreativeTemplateDetailGetV2DataTemplatePos
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeTemplateDetailGetV2DataTemplatePosFromValue(v string) (*CreativeTemplateDetailGetV2DataTemplatePos, error) {
	ev := CreativeTemplateDetailGetV2DataTemplatePos(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeTemplateDetailGetV2DataTemplatePos: valid values are %v", v, AllowedCreativeTemplateDetailGetV2DataTemplatePosEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeTemplateDetailGetV2DataTemplatePos) IsValid() bool {
	for _, existing := range AllowedCreativeTemplateDetailGetV2DataTemplatePosEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_template_detail_get_v2_data_template_pos value
func (v CreativeTemplateDetailGetV2DataTemplatePos) Ptr() *CreativeTemplateDetailGetV2DataTemplatePos {
	return &v
}
