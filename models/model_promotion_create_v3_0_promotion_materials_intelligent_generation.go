/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionCreateV30PromotionMaterialsIntelligentGeneration
type PromotionCreateV30PromotionMaterialsIntelligentGeneration string

// List of promotion_create_v3.0_promotion_materials_intelligent_generation
const (
	OFF_PromotionCreateV30PromotionMaterialsIntelligentGeneration PromotionCreateV30PromotionMaterialsIntelligentGeneration = "OFF"
	ON_PromotionCreateV30PromotionMaterialsIntelligentGeneration  PromotionCreateV30PromotionMaterialsIntelligentGeneration = "ON"
)

// All allowed values of PromotionCreateV30PromotionMaterialsIntelligentGeneration enum
var AllowedPromotionCreateV30PromotionMaterialsIntelligentGenerationEnumValues = []PromotionCreateV30PromotionMaterialsIntelligentGeneration{
	"OFF",
	"ON",
}

// NewPromotionCreateV30PromotionMaterialsIntelligentGenerationFromValue returns a pointer to a valid PromotionCreateV30PromotionMaterialsIntelligentGeneration
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionCreateV30PromotionMaterialsIntelligentGenerationFromValue(v string) (*PromotionCreateV30PromotionMaterialsIntelligentGeneration, error) {
	ev := PromotionCreateV30PromotionMaterialsIntelligentGeneration(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionCreateV30PromotionMaterialsIntelligentGeneration: valid values are %v", v, AllowedPromotionCreateV30PromotionMaterialsIntelligentGenerationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionCreateV30PromotionMaterialsIntelligentGeneration) IsValid() bool {
	for _, existing := range AllowedPromotionCreateV30PromotionMaterialsIntelligentGenerationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_create_v3.0_promotion_materials_intelligent_generation value
func (v PromotionCreateV30PromotionMaterialsIntelligentGeneration) Ptr() *PromotionCreateV30PromotionMaterialsIntelligentGeneration {
	return &v
}
