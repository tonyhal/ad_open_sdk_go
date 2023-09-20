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

// ToolsVideoCheckAvailableAnchorV2ExternalAction
type ToolsVideoCheckAvailableAnchorV2ExternalAction string

// List of tools_video_check_available_anchor_v2_external_action
const (
	AD_CONVERT_TYPE_ACTIVE_ToolsVideoCheckAvailableAnchorV2ExternalAction          ToolsVideoCheckAvailableAnchorV2ExternalAction = "AD_CONVERT_TYPE_ACTIVE"
	AD_CONVERT_TYPE_ACTIVE_REGISTER_ToolsVideoCheckAvailableAnchorV2ExternalAction ToolsVideoCheckAvailableAnchorV2ExternalAction = "AD_CONVERT_TYPE_ACTIVE_REGISTER"
	AD_CONVERT_TYPE_APP_ORDER_ToolsVideoCheckAvailableAnchorV2ExternalAction       ToolsVideoCheckAvailableAnchorV2ExternalAction = "AD_CONVERT_TYPE_APP_ORDER"
	AD_CONVERT_TYPE_APP_PAY_ToolsVideoCheckAvailableAnchorV2ExternalAction         ToolsVideoCheckAvailableAnchorV2ExternalAction = "AD_CONVERT_TYPE_APP_PAY"
	AD_CONVERT_TYPE_APP_UV_ToolsVideoCheckAvailableAnchorV2ExternalAction          ToolsVideoCheckAvailableAnchorV2ExternalAction = "AD_CONVERT_TYPE_APP_UV"
	AD_CONVERT_TYPE_PAY_ToolsVideoCheckAvailableAnchorV2ExternalAction             ToolsVideoCheckAvailableAnchorV2ExternalAction = "AD_CONVERT_TYPE_PAY"
)

// All allowed values of ToolsVideoCheckAvailableAnchorV2ExternalAction enum
var AllowedToolsVideoCheckAvailableAnchorV2ExternalActionEnumValues = []ToolsVideoCheckAvailableAnchorV2ExternalAction{
	"AD_CONVERT_TYPE_ACTIVE",
	"AD_CONVERT_TYPE_ACTIVE_REGISTER",
	"AD_CONVERT_TYPE_APP_ORDER",
	"AD_CONVERT_TYPE_APP_PAY",
	"AD_CONVERT_TYPE_APP_UV",
	"AD_CONVERT_TYPE_PAY",
}

// NewToolsVideoCheckAvailableAnchorV2ExternalActionFromValue returns a pointer to a valid ToolsVideoCheckAvailableAnchorV2ExternalAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsVideoCheckAvailableAnchorV2ExternalActionFromValue(v string) (*ToolsVideoCheckAvailableAnchorV2ExternalAction, error) {
	ev := ToolsVideoCheckAvailableAnchorV2ExternalAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsVideoCheckAvailableAnchorV2ExternalAction: valid values are %v", v, AllowedToolsVideoCheckAvailableAnchorV2ExternalActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsVideoCheckAvailableAnchorV2ExternalAction) IsValid() bool {
	for _, existing := range AllowedToolsVideoCheckAvailableAnchorV2ExternalActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_video_check_available_anchor_v2_external_action value
func (v ToolsVideoCheckAvailableAnchorV2ExternalAction) Ptr() *ToolsVideoCheckAvailableAnchorV2ExternalAction {
	return &v
}