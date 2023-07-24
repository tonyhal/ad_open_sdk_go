/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectUpdateV30AudiencePlatform
type ProjectUpdateV30AudiencePlatform string

// List of project_update_v3.0_audience_platform
const (
	ANDROID_ProjectUpdateV30AudiencePlatform ProjectUpdateV30AudiencePlatform = "ANDROID"
	IOS_ProjectUpdateV30AudiencePlatform     ProjectUpdateV30AudiencePlatform = "IOS"
)

// All allowed values of ProjectUpdateV30AudiencePlatform enum
var AllowedProjectUpdateV30AudiencePlatformEnumValues = []ProjectUpdateV30AudiencePlatform{
	"ANDROID",
	"IOS",
}

// NewProjectUpdateV30AudiencePlatformFromValue returns a pointer to a valid ProjectUpdateV30AudiencePlatform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30AudiencePlatformFromValue(v string) (*ProjectUpdateV30AudiencePlatform, error) {
	ev := ProjectUpdateV30AudiencePlatform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30AudiencePlatform: valid values are %v", v, AllowedProjectUpdateV30AudiencePlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30AudiencePlatform) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30AudiencePlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_audience_platform value
func (v ProjectUpdateV30AudiencePlatform) Ptr() *ProjectUpdateV30AudiencePlatform {
	return &v
}
