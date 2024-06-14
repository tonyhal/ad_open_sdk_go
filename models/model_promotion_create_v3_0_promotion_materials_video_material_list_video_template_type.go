/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionCreateV30PromotionMaterialsVideoMaterialListVideoTemplateType
type PromotionCreateV30PromotionMaterialsVideoMaterialListVideoTemplateType string

// List of promotion_create_v3.0_promotion_materials_video_material_list_video_template_type
const (
	DPA_VIDEO_TEMPLATE_CUSTOM_PromotionCreateV30PromotionMaterialsVideoMaterialListVideoTemplateType PromotionCreateV30PromotionMaterialsVideoMaterialListVideoTemplateType = "DPA_VIDEO_TEMPLATE_CUSTOM"
	DPA_VIDEO_TEMPLATE_SMART_PromotionCreateV30PromotionMaterialsVideoMaterialListVideoTemplateType  PromotionCreateV30PromotionMaterialsVideoMaterialListVideoTemplateType = "DPA_VIDEO_TEMPLATE_SMART"
)

// All allowed values of PromotionCreateV30PromotionMaterialsVideoMaterialListVideoTemplateType enum
var AllowedPromotionCreateV30PromotionMaterialsVideoMaterialListVideoTemplateTypeEnumValues = []PromotionCreateV30PromotionMaterialsVideoMaterialListVideoTemplateType{
	"DPA_VIDEO_TEMPLATE_CUSTOM",
	"DPA_VIDEO_TEMPLATE_SMART",
}

// NewPromotionCreateV30PromotionMaterialsVideoMaterialListVideoTemplateTypeFromValue returns a pointer to a valid PromotionCreateV30PromotionMaterialsVideoMaterialListVideoTemplateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionCreateV30PromotionMaterialsVideoMaterialListVideoTemplateTypeFromValue(v string) (*PromotionCreateV30PromotionMaterialsVideoMaterialListVideoTemplateType, error) {
	ev := PromotionCreateV30PromotionMaterialsVideoMaterialListVideoTemplateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionCreateV30PromotionMaterialsVideoMaterialListVideoTemplateType: valid values are %v", v, AllowedPromotionCreateV30PromotionMaterialsVideoMaterialListVideoTemplateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionCreateV30PromotionMaterialsVideoMaterialListVideoTemplateType) IsValid() bool {
	for _, existing := range AllowedPromotionCreateV30PromotionMaterialsVideoMaterialListVideoTemplateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_create_v3.0_promotion_materials_video_material_list_video_template_type value
func (v PromotionCreateV30PromotionMaterialsVideoMaterialListVideoTemplateType) Ptr() *PromotionCreateV30PromotionMaterialsVideoMaterialListVideoTemplateType {
	return &v
}
