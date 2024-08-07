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

// ToolsAwemeAuthListV2DataListShareType
type ToolsAwemeAuthListV2DataListShareType string

// List of tools_aweme_auth_list_v2_data_list_share_type
const (
	SHARE_BY_ONESELF_ToolsAwemeAuthListV2DataListShareType     ToolsAwemeAuthListV2DataListShareType = "SHARE_BY_ONESELF"
	SHARE_BY_SAME_ENTITY_ToolsAwemeAuthListV2DataListShareType ToolsAwemeAuthListV2DataListShareType = "SHARE_BY_SAME_ENTITY"
	SHARE_FROM_BP_ToolsAwemeAuthListV2DataListShareType        ToolsAwemeAuthListV2DataListShareType = "SHARE_FROM_BP"
)

// All allowed values of ToolsAwemeAuthListV2DataListShareType enum
var AllowedToolsAwemeAuthListV2DataListShareTypeEnumValues = []ToolsAwemeAuthListV2DataListShareType{
	"SHARE_BY_ONESELF",
	"SHARE_BY_SAME_ENTITY",
	"SHARE_FROM_BP",
}

// NewToolsAwemeAuthListV2DataListShareTypeFromValue returns a pointer to a valid ToolsAwemeAuthListV2DataListShareType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAwemeAuthListV2DataListShareTypeFromValue(v string) (*ToolsAwemeAuthListV2DataListShareType, error) {
	ev := ToolsAwemeAuthListV2DataListShareType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAwemeAuthListV2DataListShareType: valid values are %v", v, AllowedToolsAwemeAuthListV2DataListShareTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAwemeAuthListV2DataListShareType) IsValid() bool {
	for _, existing := range AllowedToolsAwemeAuthListV2DataListShareTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_aweme_auth_list_v2_data_list_share_type value
func (v ToolsAwemeAuthListV2DataListShareType) Ptr() *ToolsAwemeAuthListV2DataListShareType {
	return &v
}
