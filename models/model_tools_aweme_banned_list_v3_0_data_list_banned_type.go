/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAwemeBannedListV30DataListBannedType
type ToolsAwemeBannedListV30DataListBannedType string

// List of tools_aweme_banned_list_v3.0_data_list_banned_type
const (
	AWEME_TYPE_ToolsAwemeBannedListV30DataListBannedType  ToolsAwemeBannedListV30DataListBannedType = "AWEME_TYPE"
	CUSTOM_TYPE_ToolsAwemeBannedListV30DataListBannedType ToolsAwemeBannedListV30DataListBannedType = "CUSTOM_TYPE"
)

// All allowed values of ToolsAwemeBannedListV30DataListBannedType enum
var AllowedToolsAwemeBannedListV30DataListBannedTypeEnumValues = []ToolsAwemeBannedListV30DataListBannedType{
	"AWEME_TYPE",
	"CUSTOM_TYPE",
}

// NewToolsAwemeBannedListV30DataListBannedTypeFromValue returns a pointer to a valid ToolsAwemeBannedListV30DataListBannedType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAwemeBannedListV30DataListBannedTypeFromValue(v string) (*ToolsAwemeBannedListV30DataListBannedType, error) {
	ev := ToolsAwemeBannedListV30DataListBannedType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAwemeBannedListV30DataListBannedType: valid values are %v", v, AllowedToolsAwemeBannedListV30DataListBannedTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAwemeBannedListV30DataListBannedType) IsValid() bool {
	for _, existing := range AllowedToolsAwemeBannedListV30DataListBannedTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_aweme_banned_list_v3.0_data_list_banned_type value
func (v ToolsAwemeBannedListV30DataListBannedType) Ptr() *ToolsAwemeBannedListV30DataListBannedType {
	return &v
}
