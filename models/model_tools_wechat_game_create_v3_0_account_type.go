/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsWechatGameCreateV30AccountType
type ToolsWechatGameCreateV30AccountType string

// List of tools_wechat_game_create_v3.0_account_type
const (
	AD_ToolsWechatGameCreateV30AccountType ToolsWechatGameCreateV30AccountType = "AD"
	BP_ToolsWechatGameCreateV30AccountType ToolsWechatGameCreateV30AccountType = "BP"
)

// All allowed values of ToolsWechatGameCreateV30AccountType enum
var AllowedToolsWechatGameCreateV30AccountTypeEnumValues = []ToolsWechatGameCreateV30AccountType{
	"AD",
	"BP",
}

// NewToolsWechatGameCreateV30AccountTypeFromValue returns a pointer to a valid ToolsWechatGameCreateV30AccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsWechatGameCreateV30AccountTypeFromValue(v string) (*ToolsWechatGameCreateV30AccountType, error) {
	ev := ToolsWechatGameCreateV30AccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsWechatGameCreateV30AccountType: valid values are %v", v, AllowedToolsWechatGameCreateV30AccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsWechatGameCreateV30AccountType) IsValid() bool {
	for _, existing := range AllowedToolsWechatGameCreateV30AccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_wechat_game_create_v3.0_account_type value
func (v ToolsWechatGameCreateV30AccountType) Ptr() *ToolsWechatGameCreateV30AccountType {
	return &v
}
