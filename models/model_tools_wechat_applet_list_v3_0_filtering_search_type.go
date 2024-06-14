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

// ToolsWechatAppletListV30FilteringSearchType
type ToolsWechatAppletListV30FilteringSearchType string

// List of tools_wechat_applet_list_v3.0_filtering_search_type
const (
	CREATE_ONLY_ToolsWechatAppletListV30FilteringSearchType ToolsWechatAppletListV30FilteringSearchType = "CREATE_ONLY"
	SHARE_ONLY_ToolsWechatAppletListV30FilteringSearchType  ToolsWechatAppletListV30FilteringSearchType = "SHARE_ONLY"
)

// All allowed values of ToolsWechatAppletListV30FilteringSearchType enum
var AllowedToolsWechatAppletListV30FilteringSearchTypeEnumValues = []ToolsWechatAppletListV30FilteringSearchType{
	"CREATE_ONLY",
	"SHARE_ONLY",
}

// NewToolsWechatAppletListV30FilteringSearchTypeFromValue returns a pointer to a valid ToolsWechatAppletListV30FilteringSearchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsWechatAppletListV30FilteringSearchTypeFromValue(v string) (*ToolsWechatAppletListV30FilteringSearchType, error) {
	ev := ToolsWechatAppletListV30FilteringSearchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsWechatAppletListV30FilteringSearchType: valid values are %v", v, AllowedToolsWechatAppletListV30FilteringSearchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsWechatAppletListV30FilteringSearchType) IsValid() bool {
	for _, existing := range AllowedToolsWechatAppletListV30FilteringSearchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_wechat_applet_list_v3.0_filtering_search_type value
func (v ToolsWechatAppletListV30FilteringSearchType) Ptr() *ToolsWechatAppletListV30FilteringSearchType {
	return &v
}
