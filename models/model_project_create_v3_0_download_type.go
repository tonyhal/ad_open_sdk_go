/*
API Title

巨量引擎开放平台

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCreateV30DownloadType
type ProjectCreateV30DownloadType string

// List of project_create_v3.0_download_type
const (
	DOWNLOAD_URL_ProjectCreateV30DownloadType ProjectCreateV30DownloadType = "DOWNLOAD_URL"
	EXTERNAL_URL_ProjectCreateV30DownloadType ProjectCreateV30DownloadType = "EXTERNAL_URL"
)

// All allowed values of ProjectCreateV30DownloadType enum
var AllowedProjectCreateV30DownloadTypeEnumValues = []ProjectCreateV30DownloadType{
	"DOWNLOAD_URL",
	"EXTERNAL_URL",
}

// NewProjectCreateV30DownloadTypeFromValue returns a pointer to a valid ProjectCreateV30DownloadType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30DownloadTypeFromValue(v string) (*ProjectCreateV30DownloadType, error) {
	ev := ProjectCreateV30DownloadType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30DownloadType: valid values are %v", v, AllowedProjectCreateV30DownloadTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30DownloadType) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30DownloadTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_download_type value
func (v ProjectCreateV30DownloadType) Ptr() *ProjectCreateV30DownloadType {
	return &v
}
