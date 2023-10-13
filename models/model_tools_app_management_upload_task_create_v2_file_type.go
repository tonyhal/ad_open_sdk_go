/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAppManagementUploadTaskCreateV2FileType
type ToolsAppManagementUploadTaskCreateV2FileType string

// List of tools_app_management_upload_task_create_v2_file_type
const (
	APK_ToolsAppManagementUploadTaskCreateV2FileType   ToolsAppManagementUploadTaskCreateV2FileType = "APK"
	IMAGE_ToolsAppManagementUploadTaskCreateV2FileType ToolsAppManagementUploadTaskCreateV2FileType = "IMAGE"
	VIDEO_ToolsAppManagementUploadTaskCreateV2FileType ToolsAppManagementUploadTaskCreateV2FileType = "VIDEO"
)

// All allowed values of ToolsAppManagementUploadTaskCreateV2FileType enum
var AllowedToolsAppManagementUploadTaskCreateV2FileTypeEnumValues = []ToolsAppManagementUploadTaskCreateV2FileType{
	"APK",
	"IMAGE",
	"VIDEO",
}

// NewToolsAppManagementUploadTaskCreateV2FileTypeFromValue returns a pointer to a valid ToolsAppManagementUploadTaskCreateV2FileType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementUploadTaskCreateV2FileTypeFromValue(v string) (*ToolsAppManagementUploadTaskCreateV2FileType, error) {
	ev := ToolsAppManagementUploadTaskCreateV2FileType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementUploadTaskCreateV2FileType: valid values are %v", v, AllowedToolsAppManagementUploadTaskCreateV2FileTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementUploadTaskCreateV2FileType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementUploadTaskCreateV2FileTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_upload_task_create_v2_file_type value
func (v ToolsAppManagementUploadTaskCreateV2FileType) Ptr() *ToolsAppManagementUploadTaskCreateV2FileType {
	return &v
}
