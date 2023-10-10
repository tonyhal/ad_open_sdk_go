/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsSiteTemplateGetV2DataListBricksType
type ToolsSiteTemplateGetV2DataListBricksType string

// List of tools_site_template_get_v2_data_list_bricks_type
const (
	BUTTON_ToolsSiteTemplateGetV2DataListBricksType        ToolsSiteTemplateGetV2DataListBricksType = "BUTTON"
	COUPON_ToolsSiteTemplateGetV2DataListBricksType        ToolsSiteTemplateGetV2DataListBricksType = "COUPON"
	FORM_ToolsSiteTemplateGetV2DataListBricksType          ToolsSiteTemplateGetV2DataListBricksType = "FORM"
	PICTURE_ToolsSiteTemplateGetV2DataListBricksType       ToolsSiteTemplateGetV2DataListBricksType = "PICTURE"
	PICTURE_GROUP_ToolsSiteTemplateGetV2DataListBricksType ToolsSiteTemplateGetV2DataListBricksType = "PICTURE_GROUP"
	RICH_TEXT_ToolsSiteTemplateGetV2DataListBricksType     ToolsSiteTemplateGetV2DataListBricksType = "RICH_TEXT"
	SIMPLE_TEXT_ToolsSiteTemplateGetV2DataListBricksType   ToolsSiteTemplateGetV2DataListBricksType = "SIMPLE_TEXT"
	VIDEO_ToolsSiteTemplateGetV2DataListBricksType         ToolsSiteTemplateGetV2DataListBricksType = "VIDEO"
	WECHAT_APPLET_ToolsSiteTemplateGetV2DataListBricksType ToolsSiteTemplateGetV2DataListBricksType = "WECHAT_APPLET"
	WECHAT_GAME_ToolsSiteTemplateGetV2DataListBricksType   ToolsSiteTemplateGetV2DataListBricksType = "WECHAT_GAME"
)

// All allowed values of ToolsSiteTemplateGetV2DataListBricksType enum
var AllowedToolsSiteTemplateGetV2DataListBricksTypeEnumValues = []ToolsSiteTemplateGetV2DataListBricksType{
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

// NewToolsSiteTemplateGetV2DataListBricksTypeFromValue returns a pointer to a valid ToolsSiteTemplateGetV2DataListBricksType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsSiteTemplateGetV2DataListBricksTypeFromValue(v string) (*ToolsSiteTemplateGetV2DataListBricksType, error) {
	ev := ToolsSiteTemplateGetV2DataListBricksType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsSiteTemplateGetV2DataListBricksType: valid values are %v", v, AllowedToolsSiteTemplateGetV2DataListBricksTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsSiteTemplateGetV2DataListBricksType) IsValid() bool {
	for _, existing := range AllowedToolsSiteTemplateGetV2DataListBricksTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_site_template_get_v2_data_list_bricks_type value
func (v ToolsSiteTemplateGetV2DataListBricksType) Ptr() *ToolsSiteTemplateGetV2DataListBricksType {
	return &v
}
