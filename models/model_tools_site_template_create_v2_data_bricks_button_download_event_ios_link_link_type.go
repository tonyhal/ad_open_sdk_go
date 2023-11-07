/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsSiteTemplateCreateV2DataBricksButtonDownloadEventIosLinkLinkType
type ToolsSiteTemplateCreateV2DataBricksButtonDownloadEventIosLinkLinkType string

// List of tools_site_template_create_v2_data_bricks_button_download_event_ios_link_link_type
const (
	QUICK_APP_ToolsSiteTemplateCreateV2DataBricksButtonDownloadEventIosLinkLinkType ToolsSiteTemplateCreateV2DataBricksButtonDownloadEventIosLinkLinkType = "QUICK_APP"
	SCHEME_ToolsSiteTemplateCreateV2DataBricksButtonDownloadEventIosLinkLinkType    ToolsSiteTemplateCreateV2DataBricksButtonDownloadEventIosLinkLinkType = "SCHEME"
	URL_ToolsSiteTemplateCreateV2DataBricksButtonDownloadEventIosLinkLinkType       ToolsSiteTemplateCreateV2DataBricksButtonDownloadEventIosLinkLinkType = "URL"
)

// All allowed values of ToolsSiteTemplateCreateV2DataBricksButtonDownloadEventIosLinkLinkType enum
var AllowedToolsSiteTemplateCreateV2DataBricksButtonDownloadEventIosLinkLinkTypeEnumValues = []ToolsSiteTemplateCreateV2DataBricksButtonDownloadEventIosLinkLinkType{
	"QUICK_APP",
	"SCHEME",
	"URL",
}

// NewToolsSiteTemplateCreateV2DataBricksButtonDownloadEventIosLinkLinkTypeFromValue returns a pointer to a valid ToolsSiteTemplateCreateV2DataBricksButtonDownloadEventIosLinkLinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsSiteTemplateCreateV2DataBricksButtonDownloadEventIosLinkLinkTypeFromValue(v string) (*ToolsSiteTemplateCreateV2DataBricksButtonDownloadEventIosLinkLinkType, error) {
	ev := ToolsSiteTemplateCreateV2DataBricksButtonDownloadEventIosLinkLinkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsSiteTemplateCreateV2DataBricksButtonDownloadEventIosLinkLinkType: valid values are %v", v, AllowedToolsSiteTemplateCreateV2DataBricksButtonDownloadEventIosLinkLinkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsSiteTemplateCreateV2DataBricksButtonDownloadEventIosLinkLinkType) IsValid() bool {
	for _, existing := range AllowedToolsSiteTemplateCreateV2DataBricksButtonDownloadEventIosLinkLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_site_template_create_v2_data_bricks_button_download_event_ios_link_link_type value
func (v ToolsSiteTemplateCreateV2DataBricksButtonDownloadEventIosLinkLinkType) Ptr() *ToolsSiteTemplateCreateV2DataBricksButtonDownloadEventIosLinkLinkType {
	return &v
}
