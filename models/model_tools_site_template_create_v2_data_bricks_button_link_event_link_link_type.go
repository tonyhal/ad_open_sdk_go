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

// ToolsSiteTemplateCreateV2DataBricksButtonLinkEventLinkLinkType
type ToolsSiteTemplateCreateV2DataBricksButtonLinkEventLinkLinkType string

// List of tools_site_template_create_v2_data_bricks_button_link_event_link_link_type
const (
	QUICK_APP_ToolsSiteTemplateCreateV2DataBricksButtonLinkEventLinkLinkType ToolsSiteTemplateCreateV2DataBricksButtonLinkEventLinkLinkType = "QUICK_APP"
	SCHEME_ToolsSiteTemplateCreateV2DataBricksButtonLinkEventLinkLinkType    ToolsSiteTemplateCreateV2DataBricksButtonLinkEventLinkLinkType = "SCHEME"
	URL_ToolsSiteTemplateCreateV2DataBricksButtonLinkEventLinkLinkType       ToolsSiteTemplateCreateV2DataBricksButtonLinkEventLinkLinkType = "URL"
)

// All allowed values of ToolsSiteTemplateCreateV2DataBricksButtonLinkEventLinkLinkType enum
var AllowedToolsSiteTemplateCreateV2DataBricksButtonLinkEventLinkLinkTypeEnumValues = []ToolsSiteTemplateCreateV2DataBricksButtonLinkEventLinkLinkType{
	"QUICK_APP",
	"SCHEME",
	"URL",
}

// NewToolsSiteTemplateCreateV2DataBricksButtonLinkEventLinkLinkTypeFromValue returns a pointer to a valid ToolsSiteTemplateCreateV2DataBricksButtonLinkEventLinkLinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsSiteTemplateCreateV2DataBricksButtonLinkEventLinkLinkTypeFromValue(v string) (*ToolsSiteTemplateCreateV2DataBricksButtonLinkEventLinkLinkType, error) {
	ev := ToolsSiteTemplateCreateV2DataBricksButtonLinkEventLinkLinkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsSiteTemplateCreateV2DataBricksButtonLinkEventLinkLinkType: valid values are %v", v, AllowedToolsSiteTemplateCreateV2DataBricksButtonLinkEventLinkLinkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsSiteTemplateCreateV2DataBricksButtonLinkEventLinkLinkType) IsValid() bool {
	for _, existing := range AllowedToolsSiteTemplateCreateV2DataBricksButtonLinkEventLinkLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_site_template_create_v2_data_bricks_button_link_event_link_link_type value
func (v ToolsSiteTemplateCreateV2DataBricksButtonLinkEventLinkLinkType) Ptr() *ToolsSiteTemplateCreateV2DataBricksButtonLinkEventLinkLinkType {
	return &v
}
