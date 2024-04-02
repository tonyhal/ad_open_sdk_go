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

// ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventAndroidLinkLinkType
type ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventAndroidLinkLinkType string

// List of tools_site_template_get_v2_data_list_bricks_button_download_event_android_link_link_type
const (
	QUICK_APP_ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventAndroidLinkLinkType ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventAndroidLinkLinkType = "QUICK_APP"
	SCHEME_ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventAndroidLinkLinkType    ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventAndroidLinkLinkType = "SCHEME"
	URL_ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventAndroidLinkLinkType       ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventAndroidLinkLinkType = "URL"
)

// All allowed values of ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventAndroidLinkLinkType enum
var AllowedToolsSiteTemplateGetV2DataListBricksButtonDownloadEventAndroidLinkLinkTypeEnumValues = []ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventAndroidLinkLinkType{
	"QUICK_APP",
	"SCHEME",
	"URL",
}

// NewToolsSiteTemplateGetV2DataListBricksButtonDownloadEventAndroidLinkLinkTypeFromValue returns a pointer to a valid ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventAndroidLinkLinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsSiteTemplateGetV2DataListBricksButtonDownloadEventAndroidLinkLinkTypeFromValue(v string) (*ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventAndroidLinkLinkType, error) {
	ev := ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventAndroidLinkLinkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventAndroidLinkLinkType: valid values are %v", v, AllowedToolsSiteTemplateGetV2DataListBricksButtonDownloadEventAndroidLinkLinkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventAndroidLinkLinkType) IsValid() bool {
	for _, existing := range AllowedToolsSiteTemplateGetV2DataListBricksButtonDownloadEventAndroidLinkLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_site_template_get_v2_data_list_bricks_button_download_event_android_link_link_type value
func (v ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventAndroidLinkLinkType) Ptr() *ToolsSiteTemplateGetV2DataListBricksButtonDownloadEventAndroidLinkLinkType {
	return &v
}
