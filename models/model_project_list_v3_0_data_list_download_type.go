/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListDownloadType
type ProjectListV30DataListDownloadType string

// List of project_list_v3.0_data_list_download_type
const (
	DOWNLOAD_URL_ProjectListV30DataListDownloadType  ProjectListV30DataListDownloadType = "DOWNLOAD_URL"
	EXTERNAL_URL_ProjectListV30DataListDownloadType  ProjectListV30DataListDownloadType = "EXTERNAL_URL"
	QUICK_APP_URL_ProjectListV30DataListDownloadType ProjectListV30DataListDownloadType = "QUICK_APP_URL"
)

// All allowed values of ProjectListV30DataListDownloadType enum
var AllowedProjectListV30DataListDownloadTypeEnumValues = []ProjectListV30DataListDownloadType{
	"DOWNLOAD_URL",
	"EXTERNAL_URL",
	"QUICK_APP_URL",
}

// NewProjectListV30DataListDownloadTypeFromValue returns a pointer to a valid ProjectListV30DataListDownloadType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListDownloadTypeFromValue(v string) (*ProjectListV30DataListDownloadType, error) {
	ev := ProjectListV30DataListDownloadType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListDownloadType: valid values are %v", v, AllowedProjectListV30DataListDownloadTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListDownloadType) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListDownloadTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_download_type value
func (v ProjectListV30DataListDownloadType) Ptr() *ProjectListV30DataListDownloadType {
	return &v
}
