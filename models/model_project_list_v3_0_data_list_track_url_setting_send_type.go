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

// ProjectListV30DataListTrackUrlSettingSendType
type ProjectListV30DataListTrackUrlSettingSendType string

// List of project_list_v3.0_data_list_track_url_setting_send_type
const (
	CLIENT_SEND_ProjectListV30DataListTrackUrlSettingSendType ProjectListV30DataListTrackUrlSettingSendType = "CLIENT_SEND"
	SERVER_SEND_ProjectListV30DataListTrackUrlSettingSendType ProjectListV30DataListTrackUrlSettingSendType = "SERVER_SEND"
)

// All allowed values of ProjectListV30DataListTrackUrlSettingSendType enum
var AllowedProjectListV30DataListTrackUrlSettingSendTypeEnumValues = []ProjectListV30DataListTrackUrlSettingSendType{
	"CLIENT_SEND",
	"SERVER_SEND",
}

// NewProjectListV30DataListTrackUrlSettingSendTypeFromValue returns a pointer to a valid ProjectListV30DataListTrackUrlSettingSendType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListTrackUrlSettingSendTypeFromValue(v string) (*ProjectListV30DataListTrackUrlSettingSendType, error) {
	ev := ProjectListV30DataListTrackUrlSettingSendType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListTrackUrlSettingSendType: valid values are %v", v, AllowedProjectListV30DataListTrackUrlSettingSendTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListTrackUrlSettingSendType) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListTrackUrlSettingSendTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_track_url_setting_send_type value
func (v ProjectListV30DataListTrackUrlSettingSendType) Ptr() *ProjectListV30DataListTrackUrlSettingSendType {
	return &v
}
