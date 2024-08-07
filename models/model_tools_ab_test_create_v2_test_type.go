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

// ToolsAbTestCreateV2TestType
type ToolsAbTestCreateV2TestType string

// List of tools_ab_test_create_v2_test_type
const (
	CAMPAIGN_ToolsAbTestCreateV2TestType ToolsAbTestCreateV2TestType = "CAMPAIGN"
	AD_ToolsAbTestCreateV2TestType       ToolsAbTestCreateV2TestType = "AD"
)

// All allowed values of ToolsAbTestCreateV2TestType enum
var AllowedToolsAbTestCreateV2TestTypeEnumValues = []ToolsAbTestCreateV2TestType{
	"CAMPAIGN",
	"AD",
}

// NewToolsAbTestCreateV2TestTypeFromValue returns a pointer to a valid ToolsAbTestCreateV2TestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAbTestCreateV2TestTypeFromValue(v string) (*ToolsAbTestCreateV2TestType, error) {
	ev := ToolsAbTestCreateV2TestType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAbTestCreateV2TestType: valid values are %v", v, AllowedToolsAbTestCreateV2TestTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAbTestCreateV2TestType) IsValid() bool {
	for _, existing := range AllowedToolsAbTestCreateV2TestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ab_test_create_v2_test_type value
func (v ToolsAbTestCreateV2TestType) Ptr() *ToolsAbTestCreateV2TestType {
	return &v
}
