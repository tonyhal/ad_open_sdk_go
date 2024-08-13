/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionUpdateV30PromotionMaterialsProductInfoProductNameType
type PromotionUpdateV30PromotionMaterialsProductInfoProductNameType string

// List of promotion_update_v3.0_promotion_materials_product_info_product_name_type
const (
	CUSTOM_PromotionUpdateV30PromotionMaterialsProductInfoProductNameType PromotionUpdateV30PromotionMaterialsProductInfoProductNameType = "CUSTOM"
	DPA_PromotionUpdateV30PromotionMaterialsProductInfoProductNameType    PromotionUpdateV30PromotionMaterialsProductInfoProductNameType = "DPA"
)

// All allowed values of PromotionUpdateV30PromotionMaterialsProductInfoProductNameType enum
var AllowedPromotionUpdateV30PromotionMaterialsProductInfoProductNameTypeEnumValues = []PromotionUpdateV30PromotionMaterialsProductInfoProductNameType{
	"CUSTOM",
	"DPA",
}

// NewPromotionUpdateV30PromotionMaterialsProductInfoProductNameTypeFromValue returns a pointer to a valid PromotionUpdateV30PromotionMaterialsProductInfoProductNameType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionUpdateV30PromotionMaterialsProductInfoProductNameTypeFromValue(v string) (*PromotionUpdateV30PromotionMaterialsProductInfoProductNameType, error) {
	ev := PromotionUpdateV30PromotionMaterialsProductInfoProductNameType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionUpdateV30PromotionMaterialsProductInfoProductNameType: valid values are %v", v, AllowedPromotionUpdateV30PromotionMaterialsProductInfoProductNameTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionUpdateV30PromotionMaterialsProductInfoProductNameType) IsValid() bool {
	for _, existing := range AllowedPromotionUpdateV30PromotionMaterialsProductInfoProductNameTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_update_v3.0_promotion_materials_product_info_product_name_type value
func (v PromotionUpdateV30PromotionMaterialsProductInfoProductNameType) Ptr() *PromotionUpdateV30PromotionMaterialsProductInfoProductNameType {
	return &v
}
