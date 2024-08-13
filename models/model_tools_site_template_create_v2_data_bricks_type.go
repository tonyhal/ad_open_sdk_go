/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsSiteTemplateCreateV2DataBricksType
type ToolsSiteTemplateCreateV2DataBricksType string

// List of tools_site_template_create_v2_data_bricks_type
const (
	BUTTON_ToolsSiteTemplateCreateV2DataBricksType        ToolsSiteTemplateCreateV2DataBricksType = "BUTTON"
	COUPON_ToolsSiteTemplateCreateV2DataBricksType        ToolsSiteTemplateCreateV2DataBricksType = "COUPON"
	FORM_ToolsSiteTemplateCreateV2DataBricksType          ToolsSiteTemplateCreateV2DataBricksType = "FORM"
	PICTURE_ToolsSiteTemplateCreateV2DataBricksType       ToolsSiteTemplateCreateV2DataBricksType = "PICTURE"
	PICTURE_GROUP_ToolsSiteTemplateCreateV2DataBricksType ToolsSiteTemplateCreateV2DataBricksType = "PICTURE_GROUP"
	RICH_TEXT_ToolsSiteTemplateCreateV2DataBricksType     ToolsSiteTemplateCreateV2DataBricksType = "RICH_TEXT"
	SIMPLE_TEXT_ToolsSiteTemplateCreateV2DataBricksType   ToolsSiteTemplateCreateV2DataBricksType = "SIMPLE_TEXT"
	VIDEO_ToolsSiteTemplateCreateV2DataBricksType         ToolsSiteTemplateCreateV2DataBricksType = "VIDEO"
	WECHAT_APPLET_ToolsSiteTemplateCreateV2DataBricksType ToolsSiteTemplateCreateV2DataBricksType = "WECHAT_APPLET"
	WECHAT_GAME_ToolsSiteTemplateCreateV2DataBricksType   ToolsSiteTemplateCreateV2DataBricksType = "WECHAT_GAME"
)

// All allowed values of ToolsSiteTemplateCreateV2DataBricksType enum
var AllowedToolsSiteTemplateCreateV2DataBricksTypeEnumValues = []ToolsSiteTemplateCreateV2DataBricksType{
	"BUTTON",
	"COUPON",
	"FORM",
	"PICTURE",
	"PICTURE_GROUP",
	"RICH_TEXT",
	"SIMPLE_TEXT",
	"VIDEO",
	"WECHAT_APPLET",
	"WECHAT_GAME",
}

// NewToolsSiteTemplateCreateV2DataBricksTypeFromValue returns a pointer to a valid ToolsSiteTemplateCreateV2DataBricksType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsSiteTemplateCreateV2DataBricksTypeFromValue(v string) (*ToolsSiteTemplateCreateV2DataBricksType, error) {
	ev := ToolsSiteTemplateCreateV2DataBricksType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsSiteTemplateCreateV2DataBricksType: valid values are %v", v, AllowedToolsSiteTemplateCreateV2DataBricksTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsSiteTemplateCreateV2DataBricksType) IsValid() bool {
	for _, existing := range AllowedToolsSiteTemplateCreateV2DataBricksTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_site_template_create_v2_data_bricks_type value
func (v ToolsSiteTemplateCreateV2DataBricksType) Ptr() *ToolsSiteTemplateCreateV2DataBricksType {
	return &v
}
