/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAdConvertQueryV2DeliveryRange
type ToolsAdConvertQueryV2DeliveryRange string

// List of tools_ad_convert_query_v2_delivery_range
const (
	DEFAULT_ToolsAdConvertQueryV2DeliveryRange   ToolsAdConvertQueryV2DeliveryRange = "DEFAULT"
	UNION_ToolsAdConvertQueryV2DeliveryRange     ToolsAdConvertQueryV2DeliveryRange = "UNION"
	UNIVERSAL_ToolsAdConvertQueryV2DeliveryRange ToolsAdConvertQueryV2DeliveryRange = "UNIVERSAL"
)

// All allowed values of ToolsAdConvertQueryV2DeliveryRange enum
var AllowedToolsAdConvertQueryV2DeliveryRangeEnumValues = []ToolsAdConvertQueryV2DeliveryRange{
	"DEFAULT",
	"UNION",
	"UNIVERSAL",
}

// NewToolsAdConvertQueryV2DeliveryRangeFromValue returns a pointer to a valid ToolsAdConvertQueryV2DeliveryRange
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdConvertQueryV2DeliveryRangeFromValue(v string) (*ToolsAdConvertQueryV2DeliveryRange, error) {
	ev := ToolsAdConvertQueryV2DeliveryRange(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdConvertQueryV2DeliveryRange: valid values are %v", v, AllowedToolsAdConvertQueryV2DeliveryRangeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdConvertQueryV2DeliveryRange) IsValid() bool {
	for _, existing := range AllowedToolsAdConvertQueryV2DeliveryRangeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ad_convert_query_v2_delivery_range value
func (v ToolsAdConvertQueryV2DeliveryRange) Ptr() *ToolsAdConvertQueryV2DeliveryRange {
	return &v
}