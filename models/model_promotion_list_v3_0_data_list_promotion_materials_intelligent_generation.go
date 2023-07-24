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

// PromotionListV30DataListPromotionMaterialsIntelligentGeneration
type PromotionListV30DataListPromotionMaterialsIntelligentGeneration string

// List of promotion_list_v3.0_data_list_promotion_materials_intelligent_generation
const (
	OFF_PromotionListV30DataListPromotionMaterialsIntelligentGeneration PromotionListV30DataListPromotionMaterialsIntelligentGeneration = "OFF"
	ON_PromotionListV30DataListPromotionMaterialsIntelligentGeneration  PromotionListV30DataListPromotionMaterialsIntelligentGeneration = "ON"
)

// All allowed values of PromotionListV30DataListPromotionMaterialsIntelligentGeneration enum
var AllowedPromotionListV30DataListPromotionMaterialsIntelligentGenerationEnumValues = []PromotionListV30DataListPromotionMaterialsIntelligentGeneration{
	"OFF",
	"ON",
}

// NewPromotionListV30DataListPromotionMaterialsIntelligentGenerationFromValue returns a pointer to a valid PromotionListV30DataListPromotionMaterialsIntelligentGeneration
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListPromotionMaterialsIntelligentGenerationFromValue(v string) (*PromotionListV30DataListPromotionMaterialsIntelligentGeneration, error) {
	ev := PromotionListV30DataListPromotionMaterialsIntelligentGeneration(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListPromotionMaterialsIntelligentGeneration: valid values are %v", v, AllowedPromotionListV30DataListPromotionMaterialsIntelligentGenerationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListPromotionMaterialsIntelligentGeneration) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListPromotionMaterialsIntelligentGenerationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_promotion_materials_intelligent_generation value
func (v PromotionListV30DataListPromotionMaterialsIntelligentGeneration) Ptr() *PromotionListV30DataListPromotionMaterialsIntelligentGeneration {
	return &v
}
