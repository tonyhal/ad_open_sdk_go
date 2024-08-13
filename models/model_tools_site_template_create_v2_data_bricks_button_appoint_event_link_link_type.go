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

// ToolsSiteTemplateCreateV2DataBricksButtonAppointEventLinkLinkType
type ToolsSiteTemplateCreateV2DataBricksButtonAppointEventLinkLinkType string

// List of tools_site_template_create_v2_data_bricks_button_appoint_event_link_link_type
const (
	QUICK_APP_ToolsSiteTemplateCreateV2DataBricksButtonAppointEventLinkLinkType ToolsSiteTemplateCreateV2DataBricksButtonAppointEventLinkLinkType = "QUICK_APP"
	SCHEME_ToolsSiteTemplateCreateV2DataBricksButtonAppointEventLinkLinkType    ToolsSiteTemplateCreateV2DataBricksButtonAppointEventLinkLinkType = "SCHEME"
	URL_ToolsSiteTemplateCreateV2DataBricksButtonAppointEventLinkLinkType       ToolsSiteTemplateCreateV2DataBricksButtonAppointEventLinkLinkType = "URL"
)

// All allowed values of ToolsSiteTemplateCreateV2DataBricksButtonAppointEventLinkLinkType enum
var AllowedToolsSiteTemplateCreateV2DataBricksButtonAppointEventLinkLinkTypeEnumValues = []ToolsSiteTemplateCreateV2DataBricksButtonAppointEventLinkLinkType{
	"QUICK_APP",
	"SCHEME",
	"URL",
}

// NewToolsSiteTemplateCreateV2DataBricksButtonAppointEventLinkLinkTypeFromValue returns a pointer to a valid ToolsSiteTemplateCreateV2DataBricksButtonAppointEventLinkLinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsSiteTemplateCreateV2DataBricksButtonAppointEventLinkLinkTypeFromValue(v string) (*ToolsSiteTemplateCreateV2DataBricksButtonAppointEventLinkLinkType, error) {
	ev := ToolsSiteTemplateCreateV2DataBricksButtonAppointEventLinkLinkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsSiteTemplateCreateV2DataBricksButtonAppointEventLinkLinkType: valid values are %v", v, AllowedToolsSiteTemplateCreateV2DataBricksButtonAppointEventLinkLinkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsSiteTemplateCreateV2DataBricksButtonAppointEventLinkLinkType) IsValid() bool {
	for _, existing := range AllowedToolsSiteTemplateCreateV2DataBricksButtonAppointEventLinkLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_site_template_create_v2_data_bricks_button_appoint_event_link_link_type value
func (v ToolsSiteTemplateCreateV2DataBricksButtonAppointEventLinkLinkType) Ptr() *ToolsSiteTemplateCreateV2DataBricksButtonAppointEventLinkLinkType {
	return &v
}
