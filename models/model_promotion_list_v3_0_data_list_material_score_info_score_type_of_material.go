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

// PromotionListV30DataListMaterialScoreInfoScoreTypeOfMaterial
type PromotionListV30DataListMaterialScoreInfoScoreTypeOfMaterial string

// List of promotion_list_v3.0_data_list_material_score_info_score_type_of_material
const (
	LEVEL1_PromotionListV30DataListMaterialScoreInfoScoreTypeOfMaterial PromotionListV30DataListMaterialScoreInfoScoreTypeOfMaterial = "LEVEL1"
	LEVEL2_PromotionListV30DataListMaterialScoreInfoScoreTypeOfMaterial PromotionListV30DataListMaterialScoreInfoScoreTypeOfMaterial = "LEVEL2"
	LEVEL3_PromotionListV30DataListMaterialScoreInfoScoreTypeOfMaterial PromotionListV30DataListMaterialScoreInfoScoreTypeOfMaterial = "LEVEL3"
	LEVEL4_PromotionListV30DataListMaterialScoreInfoScoreTypeOfMaterial PromotionListV30DataListMaterialScoreInfoScoreTypeOfMaterial = "LEVEL4"
	LEVEL5_PromotionListV30DataListMaterialScoreInfoScoreTypeOfMaterial PromotionListV30DataListMaterialScoreInfoScoreTypeOfMaterial = "LEVEL5"
)

// All allowed values of PromotionListV30DataListMaterialScoreInfoScoreTypeOfMaterial enum
var AllowedPromotionListV30DataListMaterialScoreInfoScoreTypeOfMaterialEnumValues = []PromotionListV30DataListMaterialScoreInfoScoreTypeOfMaterial{
	"LEVEL1",
	"LEVEL2",
	"LEVEL3",
	"LEVEL4",
	"LEVEL5",
}

// NewPromotionListV30DataListMaterialScoreInfoScoreTypeOfMaterialFromValue returns a pointer to a valid PromotionListV30DataListMaterialScoreInfoScoreTypeOfMaterial
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListMaterialScoreInfoScoreTypeOfMaterialFromValue(v string) (*PromotionListV30DataListMaterialScoreInfoScoreTypeOfMaterial, error) {
	ev := PromotionListV30DataListMaterialScoreInfoScoreTypeOfMaterial(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListMaterialScoreInfoScoreTypeOfMaterial: valid values are %v", v, AllowedPromotionListV30DataListMaterialScoreInfoScoreTypeOfMaterialEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListMaterialScoreInfoScoreTypeOfMaterial) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListMaterialScoreInfoScoreTypeOfMaterialEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_material_score_info_score_type_of_material value
func (v PromotionListV30DataListMaterialScoreInfoScoreTypeOfMaterial) Ptr() *PromotionListV30DataListMaterialScoreInfoScoreTypeOfMaterial {
	return &v
}
