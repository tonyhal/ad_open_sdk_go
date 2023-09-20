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

// ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventIosLinkLinkType
type ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventIosLinkLinkType string

// List of tools_site_template_get_v2_data_list_bricks_button_download_event_ios_link_link_type
const (
	QUICK_APP_ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventIosLinkLinkType ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventIosLinkLinkType = "QUICK_APP"
	SCHEME_ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventIosLinkLinkType    ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventIosLinkLinkType = "SCHEME"
	URL_ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventIosLinkLinkType       ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventIosLinkLinkType = "URL"
)

// All allowed values of ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventIosLinkLinkType enum
var AllowedToolsSiteTemplateGetV2DataListBricksButtonDownloadEventIosLinkLinkTypeEnumValues = []ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventIosLinkLinkType{
	"QUICK_APP",
	"SCHEME",
	"URL",
}

// NewToolsSiteTemplateGetV2DataListBricksButtonDownloadEventIosLinkLinkTypeFromValue returns a pointer to a valid ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventIosLinkLinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsSiteTemplateGetV2DataListBricksButtonDownloadEventIosLinkLinkTypeFromValue(v string) (*ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventIosLinkLinkType, error) {
	ev := ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventIosLinkLinkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventIosLinkLinkType: valid values are %v", v, AllowedToolsSiteTemplateGetV2DataListBricksButtonDownloadEventIosLinkLinkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventIosLinkLinkType) IsValid() bool {
	for _, existing := range AllowedToolsSiteTemplateGetV2DataListBricksButtonDownloadEventIosLinkLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_site_template_get_v2_data_list_bricks_button_download_event_ios_link_link_type value
func (v ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventIosLinkLinkType) Ptr() *ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventIosLinkLinkType {
	return &v
}