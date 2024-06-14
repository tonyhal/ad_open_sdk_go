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

// PromotionUpdateV30PromotionMaterialsProductInfoProductImageType
type PromotionUpdateV30PromotionMaterialsProductInfoProductImageType string

// List of promotion_update_v3.0_promotion_materials_product_info_product_image_type
const (
	CUSTOM_PromotionUpdateV30PromotionMaterialsProductInfoProductImageType PromotionUpdateV30PromotionMaterialsProductInfoProductImageType = "CUSTOM"
	DPA_PromotionUpdateV30PromotionMaterialsProductInfoProductImageType    PromotionUpdateV30PromotionMaterialsProductInfoProductImageType = "DPA"
)

// All allowed values of PromotionUpdateV30PromotionMaterialsProductInfoProductImageType enum
var AllowedPromotionUpdateV30PromotionMaterialsProductInfoProductImageTypeEnumValues = []PromotionUpdateV30PromotionMaterialsProductInfoProductImageType{
	"CUSTOM",
	"DPA",
}

// NewPromotionUpdateV30PromotionMaterialsProductInfoProductImageTypeFromValue returns a pointer to a valid PromotionUpdateV30PromotionMaterialsProductInfoProductImageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionUpdateV30PromotionMaterialsProductInfoProductImageTypeFromValue(v string) (*PromotionUpdateV30PromotionMaterialsProductInfoProductImageType, error) {
	ev := PromotionUpdateV30PromotionMaterialsProductInfoProductImageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionUpdateV30PromotionMaterialsProductInfoProductImageType: valid values are %v", v, AllowedPromotionUpdateV30PromotionMaterialsProductInfoProductImageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionUpdateV30PromotionMaterialsProductInfoProductImageType) IsValid() bool {
	for _, existing := range AllowedPromotionUpdateV30PromotionMaterialsProductInfoProductImageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_update_v3.0_promotion_materials_product_info_product_image_type value
func (v PromotionUpdateV30PromotionMaterialsProductInfoProductImageType) Ptr() *PromotionUpdateV30PromotionMaterialsProductInfoProductImageType {
	return &v
}
