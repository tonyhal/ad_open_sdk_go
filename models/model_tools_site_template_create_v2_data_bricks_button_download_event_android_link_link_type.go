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

// ToolsSiteTemplateCreateV2DataBricksButtonDownloadEventAndroidLinkLinkType
type ToolsSiteTemplateCreateV2DataBricksButtonDownloadEventAndroidLinkLinkType string

// List of tools_site_template_create_v2_data_bricks_button_download_event_android_link_link_type
const (
	QUICK_APP_ToolsSiteTemplateCreateV2DataBricksButtonDownloadEventAndroidLinkLinkType ToolsSiteTemplateCreateV2DataBricksButtonDownloadEventAndroidLinkLinkType = "QUICK_APP"
	SCHEME_ToolsSiteTemplateCreateV2DataBricksButtonDownloadEventAndroidLinkLinkType    ToolsSiteTemplateCreateV2DataBricksButtonDownloadEventAndroidLinkLinkType = "SCHEME"
	URL_ToolsSiteTemplateCreateV2DataBricksButtonDownloadEventAndroidLinkLinkType       ToolsSiteTemplateCreateV2DataBricksButtonDownloadEventAndroidLinkLinkType = "URL"
)

// All allowed values of ToolsSiteTemplateCreateV2DataBricksButtonDownloadEventAndroidLinkLinkType enum
var AllowedToolsSiteTemplateCreateV2DataBricksButtonDownloadEventAndroidLinkLinkTypeEnumValues = []ToolsSiteTemplateCreateV2DataBricksButtonDownloadEventAndroidLinkLinkType{
	"QUICK_APP",
	"SCHEME",
	"URL",
}

// NewToolsSiteTemplateCreateV2DataBricksButtonDownloadEventAndroidLinkLinkTypeFromValue returns a pointer to a valid ToolsSiteTemplateCreateV2DataBricksButtonDownloadEventAndroidLinkLinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsSiteTemplateCreateV2DataBricksButtonDownloadEventAndroidLinkLinkTypeFromValue(v string) (*ToolsSiteTemplateCreateV2DataBricksButtonDownloadEventAndroidLinkLinkType, error) {
	ev := ToolsSiteTemplateCreateV2DataBricksButtonDownloadEventAndroidLinkLinkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsSiteTemplateCreateV2DataBricksButtonDownloadEventAndroidLinkLinkType: valid values are %v", v, AllowedToolsSiteTemplateCreateV2DataBricksButtonDownloadEventAndroidLinkLinkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsSiteTemplateCreateV2DataBricksButtonDownloadEventAndroidLinkLinkType) IsValid() bool {
	for _, existing := range AllowedToolsSiteTemplateCreateV2DataBricksButtonDownloadEventAndroidLinkLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_site_template_create_v2_data_bricks_button_download_event_android_link_link_type value
func (v ToolsSiteTemplateCreateV2DataBricksButtonDownloadEventAndroidLinkLinkType) Ptr() *ToolsSiteTemplateCreateV2DataBricksButtonDownloadEventAndroidLinkLinkType {
	return &v
}
