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

// PromotionListV30DataListPromotionMaterialsDecorationMaterialImageMode
type PromotionListV30DataListPromotionMaterialsDecorationMaterialImageMode string

// List of promotion_list_v3.0_data_list_promotion_materials_decoration_material_image_mode
const (
	CREATIVE_IMAGE_MODE_DECORATION_COUPON_PromotionListV30DataListPromotionMaterialsDecorationMaterialImageMode PromotionListV30DataListPromotionMaterialsDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_DECORATION_COUPON"
)

// All allowed values of PromotionListV30DataListPromotionMaterialsDecorationMaterialImageMode enum
var AllowedPromotionListV30DataListPromotionMaterialsDecorationMaterialImageModeEnumValues = []PromotionListV30DataListPromotionMaterialsDecorationMaterialImageMode{
	"CREATIVE_IMAGE_MODE_DECORATION_COUPON",
}

// NewPromotionListV30DataListPromotionMaterialsDecorationMaterialImageModeFromValue returns a pointer to a valid PromotionListV30DataListPromotionMaterialsDecorationMaterialImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListPromotionMaterialsDecorationMaterialImageModeFromValue(v string) (*PromotionListV30DataListPromotionMaterialsDecorationMaterialImageMode, error) {
	ev := PromotionListV30DataListPromotionMaterialsDecorationMaterialImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListPromotionMaterialsDecorationMaterialImageMode: valid values are %v", v, AllowedPromotionListV30DataListPromotionMaterialsDecorationMaterialImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListPromotionMaterialsDecorationMaterialImageMode) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListPromotionMaterialsDecorationMaterialImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_promotion_materials_decoration_material_image_mode value
func (v PromotionListV30DataListPromotionMaterialsDecorationMaterialImageMode) Ptr() *PromotionListV30DataListPromotionMaterialsDecorationMaterialImageMode {
	return &v
}
