/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeStrategyListV2DataStrategyModelsType
type CreativeStrategyListV2DataStrategyModelsType string

// List of creative_strategy_list_v2_data_strategy_models_type
const (
	HORIZONTAL2_HORIZONTAL_CreativeStrategyListV2DataStrategyModelsType CreativeStrategyListV2DataStrategyModelsType = "Horizontal2Horizontal"
	HORIZONTAL2_VERTICAL_CreativeStrategyListV2DataStrategyModelsType   CreativeStrategyListV2DataStrategyModelsType = "Horizontal2Vertical"
	VERTICAL2_HORIZONTAL_CreativeStrategyListV2DataStrategyModelsType   CreativeStrategyListV2DataStrategyModelsType = "Vertical2Horizontal"
	VERTICAL2_VERTICAL_CreativeStrategyListV2DataStrategyModelsType     CreativeStrategyListV2DataStrategyModelsType = "Vertical2Vertical"
)

// All allowed values of CreativeStrategyListV2DataStrategyModelsType enum
var AllowedCreativeStrategyListV2DataStrategyModelsTypeEnumValues = []CreativeStrategyListV2DataStrategyModelsType{
	"Horizontal2Horizontal",
	"Horizontal2Vertical",
	"Vertical2Horizontal",
	"Vertical2Vertical",
}

// NewCreativeStrategyListV2DataStrategyModelsTypeFromValue returns a pointer to a valid CreativeStrategyListV2DataStrategyModelsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeStrategyListV2DataStrategyModelsTypeFromValue(v string) (*CreativeStrategyListV2DataStrategyModelsType, error) {
	ev := CreativeStrategyListV2DataStrategyModelsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeStrategyListV2DataStrategyModelsType: valid values are %v", v, AllowedCreativeStrategyListV2DataStrategyModelsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeStrategyListV2DataStrategyModelsType) IsValid() bool {
	for _, existing := range AllowedCreativeStrategyListV2DataStrategyModelsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_strategy_list_v2_data_strategy_models_type value
func (v CreativeStrategyListV2DataStrategyModelsType) Ptr() *CreativeStrategyListV2DataStrategyModelsType {
	return &v
}