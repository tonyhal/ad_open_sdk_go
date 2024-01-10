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

// ToolsClueLiteContactGetV2DataListCallDirection
type ToolsClueLiteContactGetV2DataListCallDirection string

// List of tools_clue_lite_contact_get_v2_data_list_call_direction
const (
	CALL_OUT_ToolsClueLiteContactGetV2DataListCallDirection ToolsClueLiteContactGetV2DataListCallDirection = "CALL_OUT"
	CALL_IN_ToolsClueLiteContactGetV2DataListCallDirection  ToolsClueLiteContactGetV2DataListCallDirection = "CALL_IN"
)

// All allowed values of ToolsClueLiteContactGetV2DataListCallDirection enum
var AllowedToolsClueLiteContactGetV2DataListCallDirectionEnumValues = []ToolsClueLiteContactGetV2DataListCallDirection{
	"CALL_OUT",
	"CALL_IN",
}

// NewToolsClueLiteContactGetV2DataListCallDirectionFromValue returns a pointer to a valid ToolsClueLiteContactGetV2DataListCallDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsClueLiteContactGetV2DataListCallDirectionFromValue(v string) (*ToolsClueLiteContactGetV2DataListCallDirection, error) {
	ev := ToolsClueLiteContactGetV2DataListCallDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsClueLiteContactGetV2DataListCallDirection: valid values are %v", v, AllowedToolsClueLiteContactGetV2DataListCallDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsClueLiteContactGetV2DataListCallDirection) IsValid() bool {
	for _, existing := range AllowedToolsClueLiteContactGetV2DataListCallDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_clue_lite_contact_get_v2_data_list_call_direction value
func (v ToolsClueLiteContactGetV2DataListCallDirection) Ptr() *ToolsClueLiteContactGetV2DataListCallDirection {
	return &v
}
