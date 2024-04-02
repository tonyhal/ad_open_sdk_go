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

// ToolsAwemeAuthV2AuthType
type ToolsAwemeAuthV2AuthType string

// List of tools_aweme_auth_v2_auth_type
const (
	AWEME_ACCOUNT_ToolsAwemeAuthV2AuthType ToolsAwemeAuthV2AuthType = "AWEME_ACCOUNT"
	LIVE_ACCOUNT_ToolsAwemeAuthV2AuthType  ToolsAwemeAuthV2AuthType = "LIVE_ACCOUNT"
	VIDEO_ITEM_ToolsAwemeAuthV2AuthType    ToolsAwemeAuthV2AuthType = "VIDEO_ITEM"
)

// All allowed values of ToolsAwemeAuthV2AuthType enum
var AllowedToolsAwemeAuthV2AuthTypeEnumValues = []ToolsAwemeAuthV2AuthType{
	"AWEME_ACCOUNT",
	"LIVE_ACCOUNT",
	"VIDEO_ITEM",
}

// NewToolsAwemeAuthV2AuthTypeFromValue returns a pointer to a valid ToolsAwemeAuthV2AuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAwemeAuthV2AuthTypeFromValue(v string) (*ToolsAwemeAuthV2AuthType, error) {
	ev := ToolsAwemeAuthV2AuthType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAwemeAuthV2AuthType: valid values are %v", v, AllowedToolsAwemeAuthV2AuthTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAwemeAuthV2AuthType) IsValid() bool {
	for _, existing := range AllowedToolsAwemeAuthV2AuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_aweme_auth_v2_auth_type value
func (v ToolsAwemeAuthV2AuthType) Ptr() *ToolsAwemeAuthV2AuthType {
	return &v
}
