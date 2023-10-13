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

// PromotionListV30DataListPromotionMaterialsProductInfoProductNameType
type PromotionListV30DataListPromotionMaterialsProductInfoProductNameType string

// List of promotion_list_v3.0_data_list_promotion_materials_product_info_product_name_type
const (
	CUSTOM_PromotionListV30DataListPromotionMaterialsProductInfoProductNameType PromotionListV30DataListPromotionMaterialsProductInfoProductNameType = "CUSTOM"
	DPA_PromotionListV30DataListPromotionMaterialsProductInfoProductNameType    PromotionListV30DataListPromotionMaterialsProductInfoProductNameType = "DPA"
)

// All allowed values of PromotionListV30DataListPromotionMaterialsProductInfoProductNameType enum
var AllowedPromotionListV30DataListPromotionMaterialsProductInfoProductNameTypeEnumValues = []PromotionListV30DataListPromotionMaterialsProductInfoProductNameType{
	"CUSTOM",
	"DPA",
}

// NewPromotionListV30DataListPromotionMaterialsProductInfoProductNameTypeFromValue returns a pointer to a valid PromotionListV30DataListPromotionMaterialsProductInfoProductNameType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListPromotionMaterialsProductInfoProductNameTypeFromValue(v string) (*PromotionListV30DataListPromotionMaterialsProductInfoProductNameType, error) {
	ev := PromotionListV30DataListPromotionMaterialsProductInfoProductNameType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListPromotionMaterialsProductInfoProductNameType: valid values are %v", v, AllowedPromotionListV30DataListPromotionMaterialsProductInfoProductNameTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListPromotionMaterialsProductInfoProductNameType) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListPromotionMaterialsProductInfoProductNameTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_promotion_materials_product_info_product_name_type value
func (v PromotionListV30DataListPromotionMaterialsProductInfoProductNameType) Ptr() *PromotionListV30DataListPromotionMaterialsProductInfoProductNameType {
	return &v
}
