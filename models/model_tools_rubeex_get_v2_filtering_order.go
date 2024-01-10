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

// ToolsRubeexGetV2FilteringOrder
type ToolsRubeexGetV2FilteringOrder string

// List of tools_rubeex_get_v2_filtering_order
const (
	DESC_ToolsRubeexGetV2FilteringOrder ToolsRubeexGetV2FilteringOrder = "DESC"
	ASC_ToolsRubeexGetV2FilteringOrder  ToolsRubeexGetV2FilteringOrder = "ASC"
)

// All allowed values of ToolsRubeexGetV2FilteringOrder enum
var AllowedToolsRubeexGetV2FilteringOrderEnumValues = []ToolsRubeexGetV2FilteringOrder{
	"DESC",
	"ASC",
}

// NewToolsRubeexGetV2FilteringOrderFromValue returns a pointer to a valid ToolsRubeexGetV2FilteringOrder
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsRubeexGetV2FilteringOrderFromValue(v string) (*ToolsRubeexGetV2FilteringOrder, error) {
	ev := ToolsRubeexGetV2FilteringOrder(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsRubeexGetV2FilteringOrder: valid values are %v", v, AllowedToolsRubeexGetV2FilteringOrderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsRubeexGetV2FilteringOrder) IsValid() bool {
	for _, existing := range AllowedToolsRubeexGetV2FilteringOrderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_rubeex_get_v2_filtering_order value
func (v ToolsRubeexGetV2FilteringOrder) Ptr() *ToolsRubeexGetV2FilteringOrder {
	return &v
}
