/*
API Title

巨量引擎开放平台

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionCreateV30PromotionMaterialsProductInfoProductSellingPointType
type PromotionCreateV30PromotionMaterialsProductInfoProductSellingPointType string

// List of promotion_create_v3.0_promotion_materials_product_info_product_selling_point_type
const (
	CUSTOM_PromotionCreateV30PromotionMaterialsProductInfoProductSellingPointType PromotionCreateV30PromotionMaterialsProductInfoProductSellingPointType = "CUSTOM"
	DPA_PromotionCreateV30PromotionMaterialsProductInfoProductSellingPointType    PromotionCreateV30PromotionMaterialsProductInfoProductSellingPointType = "DPA"
)

// All allowed values of PromotionCreateV30PromotionMaterialsProductInfoProductSellingPointType enum
var AllowedPromotionCreateV30PromotionMaterialsProductInfoProductSellingPointTypeEnumValues = []PromotionCreateV30PromotionMaterialsProductInfoProductSellingPointType{
	"CUSTOM",
	"DPA",
}

// NewPromotionCreateV30PromotionMaterialsProductInfoProductSellingPointTypeFromValue returns a pointer to a valid PromotionCreateV30PromotionMaterialsProductInfoProductSellingPointType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionCreateV30PromotionMaterialsProductInfoProductSellingPointTypeFromValue(v string) (*PromotionCreateV30PromotionMaterialsProductInfoProductSellingPointType, error) {
	ev := PromotionCreateV30PromotionMaterialsProductInfoProductSellingPointType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionCreateV30PromotionMaterialsProductInfoProductSellingPointType: valid values are %v", v, AllowedPromotionCreateV30PromotionMaterialsProductInfoProductSellingPointTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionCreateV30PromotionMaterialsProductInfoProductSellingPointType) IsValid() bool {
	for _, existing := range AllowedPromotionCreateV30PromotionMaterialsProductInfoProductSellingPointTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_create_v3.0_promotion_materials_product_info_product_selling_point_type value
func (v PromotionCreateV30PromotionMaterialsProductInfoProductSellingPointType) Ptr() *PromotionCreateV30PromotionMaterialsProductInfoProductSellingPointType {
	return &v
}
