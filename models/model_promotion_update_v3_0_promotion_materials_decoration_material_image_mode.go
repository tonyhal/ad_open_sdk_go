/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionUpdateV30PromotionMaterialsDecorationMaterialImageMode
type PromotionUpdateV30PromotionMaterialsDecorationMaterialImageMode string

// List of promotion_update_v3.0_promotion_materials_decoration_material_image_mode
const (
	CREATIVE_IMAGE_MODE_DECORATION_COUPON_PromotionUpdateV30PromotionMaterialsDecorationMaterialImageMode PromotionUpdateV30PromotionMaterialsDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_DECORATION_COUPON"
)

// All allowed values of PromotionUpdateV30PromotionMaterialsDecorationMaterialImageMode enum
var AllowedPromotionUpdateV30PromotionMaterialsDecorationMaterialImageModeEnumValues = []PromotionUpdateV30PromotionMaterialsDecorationMaterialImageMode{
	"CREATIVE_IMAGE_MODE_DECORATION_COUPON",
}

// NewPromotionUpdateV30PromotionMaterialsDecorationMaterialImageModeFromValue returns a pointer to a valid PromotionUpdateV30PromotionMaterialsDecorationMaterialImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionUpdateV30PromotionMaterialsDecorationMaterialImageModeFromValue(v string) (*PromotionUpdateV30PromotionMaterialsDecorationMaterialImageMode, error) {
	ev := PromotionUpdateV30PromotionMaterialsDecorationMaterialImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionUpdateV30PromotionMaterialsDecorationMaterialImageMode: valid values are %v", v, AllowedPromotionUpdateV30PromotionMaterialsDecorationMaterialImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionUpdateV30PromotionMaterialsDecorationMaterialImageMode) IsValid() bool {
	for _, existing := range AllowedPromotionUpdateV30PromotionMaterialsDecorationMaterialImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_update_v3.0_promotion_materials_decoration_material_image_mode value
func (v PromotionUpdateV30PromotionMaterialsDecorationMaterialImageMode) Ptr() *PromotionUpdateV30PromotionMaterialsDecorationMaterialImageMode {
	return &v
}
