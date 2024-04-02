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

// PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoTemplateType
type PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoTemplateType string

// List of promotion_update_v3.0_promotion_materials_video_material_list_video_template_type
const (
	DPA_VIDEO_TEMPLATE_CUSTOM_PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoTemplateType     PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoTemplateType = "DPA_VIDEO_TEMPLATE_CUSTOM"
	DPA_VIDEO_TEMPLATE_DEPRECATED_PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoTemplateType PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoTemplateType = "DPA_VIDEO_TEMPLATE_DEPRECATED"
	DPA_VIDEO_TEMPLATE_SMART_PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoTemplateType      PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoTemplateType = "DPA_VIDEO_TEMPLATE_SMART"
)

// All allowed values of PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoTemplateType enum
var AllowedPromotionUpdateV30PromotionMaterialsVideoMaterialListVideoTemplateTypeEnumValues = []PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoTemplateType{
	"DPA_VIDEO_TEMPLATE_CUSTOM",
	"DPA_VIDEO_TEMPLATE_DEPRECATED",
	"DPA_VIDEO_TEMPLATE_SMART",
}

// NewPromotionUpdateV30PromotionMaterialsVideoMaterialListVideoTemplateTypeFromValue returns a pointer to a valid PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoTemplateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionUpdateV30PromotionMaterialsVideoMaterialListVideoTemplateTypeFromValue(v string) (*PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoTemplateType, error) {
	ev := PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoTemplateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoTemplateType: valid values are %v", v, AllowedPromotionUpdateV30PromotionMaterialsVideoMaterialListVideoTemplateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoTemplateType) IsValid() bool {
	for _, existing := range AllowedPromotionUpdateV30PromotionMaterialsVideoMaterialListVideoTemplateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_update_v3.0_promotion_materials_video_material_list_video_template_type value
func (v PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoTemplateType) Ptr() *PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoTemplateType {
	return &v
}
