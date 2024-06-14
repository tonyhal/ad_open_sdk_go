/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsRtaGetInfoV2AccountType
type ToolsRtaGetInfoV2AccountType string

// List of tools_rta_get_info_v2_account_type
const (
	AD_ToolsRtaGetInfoV2AccountType   ToolsRtaGetInfoV2AccountType = "AD"
	STAR_ToolsRtaGetInfoV2AccountType ToolsRtaGetInfoV2AccountType = "STAR"
)

// All allowed values of ToolsRtaGetInfoV2AccountType enum
var AllowedToolsRtaGetInfoV2AccountTypeEnumValues = []ToolsRtaGetInfoV2AccountType{
	"AD",
	"STAR",
}

// NewToolsRtaGetInfoV2AccountTypeFromValue returns a pointer to a valid ToolsRtaGetInfoV2AccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsRtaGetInfoV2AccountTypeFromValue(v string) (*ToolsRtaGetInfoV2AccountType, error) {
	ev := ToolsRtaGetInfoV2AccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsRtaGetInfoV2AccountType: valid values are %v", v, AllowedToolsRtaGetInfoV2AccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsRtaGetInfoV2AccountType) IsValid() bool {
	for _, existing := range AllowedToolsRtaGetInfoV2AccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_rta_get_info_v2_account_type value
func (v ToolsRtaGetInfoV2AccountType) Ptr() *ToolsRtaGetInfoV2AccountType {
	return &v
}
