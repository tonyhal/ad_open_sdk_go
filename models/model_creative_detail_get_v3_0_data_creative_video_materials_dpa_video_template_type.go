/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeDetailGetV30DataCreativeVideoMaterialsDpaVideoTemplateType
type CreativeDetailGetV30DataCreativeVideoMaterialsDpaVideoTemplateType string

// List of creative_detail_get_v3.0_data_creative_video_materials_dpa_video_template_type
const (
	DPA_VIDEO_TEMPLATE_CUSTOM_CreativeDetailGetV30DataCreativeVideoMaterialsDpaVideoTemplateType     CreativeDetailGetV30DataCreativeVideoMaterialsDpaVideoTemplateType = "DPA_VIDEO_TEMPLATE_CUSTOM"
	DPA_VIDEO_TEMPLATE_DEPRECATED_CreativeDetailGetV30DataCreativeVideoMaterialsDpaVideoTemplateType CreativeDetailGetV30DataCreativeVideoMaterialsDpaVideoTemplateType = "DPA_VIDEO_TEMPLATE_DEPRECATED"
	DPA_VIDEO_TEMPLATE_SMART_CreativeDetailGetV30DataCreativeVideoMaterialsDpaVideoTemplateType      CreativeDetailGetV30DataCreativeVideoMaterialsDpaVideoTemplateType = "DPA_VIDEO_TEMPLATE_SMART"
)

// All allowed values of CreativeDetailGetV30DataCreativeVideoMaterialsDpaVideoTemplateType enum
var AllowedCreativeDetailGetV30DataCreativeVideoMaterialsDpaVideoTemplateTypeEnumValues = []CreativeDetailGetV30DataCreativeVideoMaterialsDpaVideoTemplateType{
	"DPA_VIDEO_TEMPLATE_CUSTOM",
	"DPA_VIDEO_TEMPLATE_DEPRECATED",
	"DPA_VIDEO_TEMPLATE_SMART",
}

// NewCreativeDetailGetV30DataCreativeVideoMaterialsDpaVideoTemplateTypeFromValue returns a pointer to a valid CreativeDetailGetV30DataCreativeVideoMaterialsDpaVideoTemplateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeDetailGetV30DataCreativeVideoMaterialsDpaVideoTemplateTypeFromValue(v string) (*CreativeDetailGetV30DataCreativeVideoMaterialsDpaVideoTemplateType, error) {
	ev := CreativeDetailGetV30DataCreativeVideoMaterialsDpaVideoTemplateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeDetailGetV30DataCreativeVideoMaterialsDpaVideoTemplateType: valid values are %v", v, AllowedCreativeDetailGetV30DataCreativeVideoMaterialsDpaVideoTemplateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeDetailGetV30DataCreativeVideoMaterialsDpaVideoTemplateType) IsValid() bool {
	for _, existing := range AllowedCreativeDetailGetV30DataCreativeVideoMaterialsDpaVideoTemplateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_detail_get_v3.0_data_creative_video_materials_dpa_video_template_type value
func (v CreativeDetailGetV30DataCreativeVideoMaterialsDpaVideoTemplateType) Ptr() *CreativeDetailGetV30DataCreativeVideoMaterialsDpaVideoTemplateType {
	return &v
}
