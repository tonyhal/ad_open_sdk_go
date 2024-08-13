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

// ToolsSiteTemplateGetV2DataListBricksPictureLinkDtoLinkType
type ToolsSiteTemplateGetV2DataListBricksPictureLinkDtoLinkType string

// List of tools_site_template_get_v2_data_list_bricks_picture_link_dto_link_type
const (
	QUICK_APP_ToolsSiteTemplateGetV2DataListBricksPictureLinkDtoLinkType ToolsSiteTemplateGetV2DataListBricksPictureLinkDtoLinkType = "QUICK_APP"
	SCHEME_ToolsSiteTemplateGetV2DataListBricksPictureLinkDtoLinkType    ToolsSiteTemplateGetV2DataListBricksPictureLinkDtoLinkType = "SCHEME"
	URL_ToolsSiteTemplateGetV2DataListBricksPictureLinkDtoLinkType       ToolsSiteTemplateGetV2DataListBricksPictureLinkDtoLinkType = "URL"
)

// All allowed values of ToolsSiteTemplateGetV2DataListBricksPictureLinkDtoLinkType enum
var AllowedToolsSiteTemplateGetV2DataListBricksPictureLinkDtoLinkTypeEnumValues = []ToolsSiteTemplateGetV2DataListBricksPictureLinkDtoLinkType{
	"QUICK_APP",
	"SCHEME",
	"URL",
}

// NewToolsSiteTemplateGetV2DataListBricksPictureLinkDtoLinkTypeFromValue returns a pointer to a valid ToolsSiteTemplateGetV2DataListBricksPictureLinkDtoLinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsSiteTemplateGetV2DataListBricksPictureLinkDtoLinkTypeFromValue(v string) (*ToolsSiteTemplateGetV2DataListBricksPictureLinkDtoLinkType, error) {
	ev := ToolsSiteTemplateGetV2DataListBricksPictureLinkDtoLinkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsSiteTemplateGetV2DataListBricksPictureLinkDtoLinkType: valid values are %v", v, AllowedToolsSiteTemplateGetV2DataListBricksPictureLinkDtoLinkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsSiteTemplateGetV2DataListBricksPictureLinkDtoLinkType) IsValid() bool {
	for _, existing := range AllowedToolsSiteTemplateGetV2DataListBricksPictureLinkDtoLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_site_template_get_v2_data_list_bricks_picture_link_dto_link_type value
func (v ToolsSiteTemplateGetV2DataListBricksPictureLinkDtoLinkType) Ptr() *ToolsSiteTemplateGetV2DataListBricksPictureLinkDtoLinkType {
	return &v
}
