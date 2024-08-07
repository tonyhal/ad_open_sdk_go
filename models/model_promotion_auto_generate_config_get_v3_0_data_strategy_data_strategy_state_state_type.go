/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionAutoGenerateConfigGetV30DataStrategyDataStrategyStateStateType
type PromotionAutoGenerateConfigGetV30DataStrategyDataStrategyStateStateType string

// List of promotion_auto_generate_config_get_v3.0_data_strategy_data_strategy_state_state_type
const (
	IMAGE_PromotionAutoGenerateConfigGetV30DataStrategyDataStrategyStateStateType PromotionAutoGenerateConfigGetV30DataStrategyDataStrategyStateStateType = "Image"
	TEXT_PromotionAutoGenerateConfigGetV30DataStrategyDataStrategyStateStateType  PromotionAutoGenerateConfigGetV30DataStrategyDataStrategyStateStateType = "Text"
)

// All allowed values of PromotionAutoGenerateConfigGetV30DataStrategyDataStrategyStateStateType enum
var AllowedPromotionAutoGenerateConfigGetV30DataStrategyDataStrategyStateStateTypeEnumValues = []PromotionAutoGenerateConfigGetV30DataStrategyDataStrategyStateStateType{
	"Image",
	"Text",
}

// NewPromotionAutoGenerateConfigGetV30DataStrategyDataStrategyStateStateTypeFromValue returns a pointer to a valid PromotionAutoGenerateConfigGetV30DataStrategyDataStrategyStateStateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionAutoGenerateConfigGetV30DataStrategyDataStrategyStateStateTypeFromValue(v string) (*PromotionAutoGenerateConfigGetV30DataStrategyDataStrategyStateStateType, error) {
	ev := PromotionAutoGenerateConfigGetV30DataStrategyDataStrategyStateStateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionAutoGenerateConfigGetV30DataStrategyDataStrategyStateStateType: valid values are %v", v, AllowedPromotionAutoGenerateConfigGetV30DataStrategyDataStrategyStateStateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionAutoGenerateConfigGetV30DataStrategyDataStrategyStateStateType) IsValid() bool {
	for _, existing := range AllowedPromotionAutoGenerateConfigGetV30DataStrategyDataStrategyStateStateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_auto_generate_config_get_v3.0_data_strategy_data_strategy_state_state_type value
func (v PromotionAutoGenerateConfigGetV30DataStrategyDataStrategyStateStateType) Ptr() *PromotionAutoGenerateConfigGetV30DataStrategyDataStrategyStateStateType {
	return &v
}
