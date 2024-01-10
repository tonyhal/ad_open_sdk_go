/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionAutoGenerateConfigCreateV30StrategyDataStrategyStateStateType
type PromotionAutoGenerateConfigCreateV30StrategyDataStrategyStateStateType string

// List of promotion_auto_generate_config_create_v3.0_strategy_data_strategy_state_state_type
const (
	IMAGE_PromotionAutoGenerateConfigCreateV30StrategyDataStrategyStateStateType PromotionAutoGenerateConfigCreateV30StrategyDataStrategyStateStateType = "Image"
	TEXT_PromotionAutoGenerateConfigCreateV30StrategyDataStrategyStateStateType  PromotionAutoGenerateConfigCreateV30StrategyDataStrategyStateStateType = "Text"
)

// All allowed values of PromotionAutoGenerateConfigCreateV30StrategyDataStrategyStateStateType enum
var AllowedPromotionAutoGenerateConfigCreateV30StrategyDataStrategyStateStateTypeEnumValues = []PromotionAutoGenerateConfigCreateV30StrategyDataStrategyStateStateType{
	"Image",
	"Text",
}

// NewPromotionAutoGenerateConfigCreateV30StrategyDataStrategyStateStateTypeFromValue returns a pointer to a valid PromotionAutoGenerateConfigCreateV30StrategyDataStrategyStateStateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionAutoGenerateConfigCreateV30StrategyDataStrategyStateStateTypeFromValue(v string) (*PromotionAutoGenerateConfigCreateV30StrategyDataStrategyStateStateType, error) {
	ev := PromotionAutoGenerateConfigCreateV30StrategyDataStrategyStateStateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionAutoGenerateConfigCreateV30StrategyDataStrategyStateStateType: valid values are %v", v, AllowedPromotionAutoGenerateConfigCreateV30StrategyDataStrategyStateStateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionAutoGenerateConfigCreateV30StrategyDataStrategyStateStateType) IsValid() bool {
	for _, existing := range AllowedPromotionAutoGenerateConfigCreateV30StrategyDataStrategyStateStateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_auto_generate_config_create_v3.0_strategy_data_strategy_state_state_type value
func (v PromotionAutoGenerateConfigCreateV30StrategyDataStrategyStateStateType) Ptr() *PromotionAutoGenerateConfigCreateV30StrategyDataStrategyStateStateType {
	return &v
}
