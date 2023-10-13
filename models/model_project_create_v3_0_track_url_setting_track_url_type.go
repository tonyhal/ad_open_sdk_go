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

// ProjectCreateV30TrackUrlSettingTrackUrlType
type ProjectCreateV30TrackUrlSettingTrackUrlType string

// List of project_create_v3.0_track_url_setting_track_url_type
const (
	CUSTOM_ProjectCreateV30TrackUrlSettingTrackUrlType   ProjectCreateV30TrackUrlSettingTrackUrlType = "CUSTOM"
	GROUP_ID_ProjectCreateV30TrackUrlSettingTrackUrlType ProjectCreateV30TrackUrlSettingTrackUrlType = "GROUP_ID"
)

// All allowed values of ProjectCreateV30TrackUrlSettingTrackUrlType enum
var AllowedProjectCreateV30TrackUrlSettingTrackUrlTypeEnumValues = []ProjectCreateV30TrackUrlSettingTrackUrlType{
	"CUSTOM",
	"GROUP_ID",
}

// NewProjectCreateV30TrackUrlSettingTrackUrlTypeFromValue returns a pointer to a valid ProjectCreateV30TrackUrlSettingTrackUrlType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30TrackUrlSettingTrackUrlTypeFromValue(v string) (*ProjectCreateV30TrackUrlSettingTrackUrlType, error) {
	ev := ProjectCreateV30TrackUrlSettingTrackUrlType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30TrackUrlSettingTrackUrlType: valid values are %v", v, AllowedProjectCreateV30TrackUrlSettingTrackUrlTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30TrackUrlSettingTrackUrlType) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30TrackUrlSettingTrackUrlTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_track_url_setting_track_url_type value
func (v ProjectCreateV30TrackUrlSettingTrackUrlType) Ptr() *ProjectCreateV30TrackUrlSettingTrackUrlType {
	return &v
}
