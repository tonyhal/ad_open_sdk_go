/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeTemplateListGetV2ImageMode
type CreativeTemplateListGetV2ImageMode string

// List of creative_template_list_get_v2_image_mode
const (
	VIDEO_HORIZONTAL_CreativeTemplateListGetV2ImageMode CreativeTemplateListGetV2ImageMode = "VIDEO_HORIZONTAL"
	VIDEO_VERTICAL_CreativeTemplateListGetV2ImageMode   CreativeTemplateListGetV2ImageMode = "VIDEO_VERTICAL"
)

// All allowed values of CreativeTemplateListGetV2ImageMode enum
var AllowedCreativeTemplateListGetV2ImageModeEnumValues = []CreativeTemplateListGetV2ImageMode{
	"VIDEO_HORIZONTAL",
	"VIDEO_VERTICAL",
}

// NewCreativeTemplateListGetV2ImageModeFromValue returns a pointer to a valid CreativeTemplateListGetV2ImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeTemplateListGetV2ImageModeFromValue(v string) (*CreativeTemplateListGetV2ImageMode, error) {
	ev := CreativeTemplateListGetV2ImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeTemplateListGetV2ImageMode: valid values are %v", v, AllowedCreativeTemplateListGetV2ImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeTemplateListGetV2ImageMode) IsValid() bool {
	for _, existing := range AllowedCreativeTemplateListGetV2ImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_template_list_get_v2_image_mode value
func (v CreativeTemplateListGetV2ImageMode) Ptr() *CreativeTemplateListGetV2ImageMode {
	return &v
}
