/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsSiteTemplateGetV2DataListBricksButtonLinkEventLinkLinkType
type ToolsSiteTemplateGetV2DataListBricksButtonLinkEventLinkLinkType string

// List of tools_site_template_get_v2_data_list_bricks_button_link_event_link_link_type
const (
	QUICK_APP_ToolsSiteTemplateGetV2DataListBricksButtonLinkEventLinkLinkType ToolsSiteTemplateGetV2DataListBricksButtonLinkEventLinkLinkType = "QUICK_APP"
	SCHEME_ToolsSiteTemplateGetV2DataListBricksButtonLinkEventLinkLinkType    ToolsSiteTemplateGetV2DataListBricksButtonLinkEventLinkLinkType = "SCHEME"
	URL_ToolsSiteTemplateGetV2DataListBricksButtonLinkEventLinkLinkType       ToolsSiteTemplateGetV2DataListBricksButtonLinkEventLinkLinkType = "URL"
)

// All allowed values of ToolsSiteTemplateGetV2DataListBricksButtonLinkEventLinkLinkType enum
var AllowedToolsSiteTemplateGetV2DataListBricksButtonLinkEventLinkLinkTypeEnumValues = []ToolsSiteTemplateGetV2DataListBricksButtonLinkEventLinkLinkType{
	"QUICK_APP",
	"SCHEME",
	"URL",
}

// NewToolsSiteTemplateGetV2DataListBricksButtonLinkEventLinkLinkTypeFromValue returns a pointer to a valid ToolsSiteTemplateGetV2DataListBricksButtonLinkEventLinkLinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsSiteTemplateGetV2DataListBricksButtonLinkEventLinkLinkTypeFromValue(v string) (*ToolsSiteTemplateGetV2DataListBricksButtonLinkEventLinkLinkType, error) {
	ev := ToolsSiteTemplateGetV2DataListBricksButtonLinkEventLinkLinkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsSiteTemplateGetV2DataListBricksButtonLinkEventLinkLinkType: valid values are %v", v, AllowedToolsSiteTemplateGetV2DataListBricksButtonLinkEventLinkLinkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsSiteTemplateGetV2DataListBricksButtonLinkEventLinkLinkType) IsValid() bool {
	for _, existing := range AllowedToolsSiteTemplateGetV2DataListBricksButtonLinkEventLinkLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_site_template_get_v2_data_list_bricks_button_link_event_link_link_type value
func (v ToolsSiteTemplateGetV2DataListBricksButtonLinkEventLinkLinkType) Ptr() *ToolsSiteTemplateGetV2DataListBricksButtonLinkEventLinkLinkType {
	return &v
}
