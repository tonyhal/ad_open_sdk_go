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

// PromotionListV30DataListPromotionMaterialsOpenUrlType
type PromotionListV30DataListPromotionMaterialsOpenUrlType string

// List of promotion_list_v3.0_data_list_promotion_materials_open_url_type
const (
	CUSTOM_PromotionListV30DataListPromotionMaterialsOpenUrlType PromotionListV30DataListPromotionMaterialsOpenUrlType = "CUSTOM"
	DPA_PromotionListV30DataListPromotionMaterialsOpenUrlType    PromotionListV30DataListPromotionMaterialsOpenUrlType = "DPA"
	NONE_PromotionListV30DataListPromotionMaterialsOpenUrlType   PromotionListV30DataListPromotionMaterialsOpenUrlType = "NONE"
)

// All allowed values of PromotionListV30DataListPromotionMaterialsOpenUrlType enum
var AllowedPromotionListV30DataListPromotionMaterialsOpenUrlTypeEnumValues = []PromotionListV30DataListPromotionMaterialsOpenUrlType{
	"CUSTOM",
	"DPA",
	"NONE",
}

// NewPromotionListV30DataListPromotionMaterialsOpenUrlTypeFromValue returns a pointer to a valid PromotionListV30DataListPromotionMaterialsOpenUrlType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListPromotionMaterialsOpenUrlTypeFromValue(v string) (*PromotionListV30DataListPromotionMaterialsOpenUrlType, error) {
	ev := PromotionListV30DataListPromotionMaterialsOpenUrlType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListPromotionMaterialsOpenUrlType: valid values are %v", v, AllowedPromotionListV30DataListPromotionMaterialsOpenUrlTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListPromotionMaterialsOpenUrlType) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListPromotionMaterialsOpenUrlTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_promotion_materials_open_url_type value
func (v PromotionListV30DataListPromotionMaterialsOpenUrlType) Ptr() *PromotionListV30DataListPromotionMaterialsOpenUrlType {
	return &v
}
