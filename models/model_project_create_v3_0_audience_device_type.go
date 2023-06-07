/*
API Title

巨量引擎开放平台

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCreateV30AudienceDeviceType
type ProjectCreateV30AudienceDeviceType string

// List of project_create_v3.0_audience_device_type
const (
	MOBILE_ProjectCreateV30AudienceDeviceType ProjectCreateV30AudienceDeviceType = "MOBILE"
	PAD_ProjectCreateV30AudienceDeviceType    ProjectCreateV30AudienceDeviceType = "PAD"
)

// All allowed values of ProjectCreateV30AudienceDeviceType enum
var AllowedProjectCreateV30AudienceDeviceTypeEnumValues = []ProjectCreateV30AudienceDeviceType{
	"MOBILE",
	"PAD",
}

// NewProjectCreateV30AudienceDeviceTypeFromValue returns a pointer to a valid ProjectCreateV30AudienceDeviceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30AudienceDeviceTypeFromValue(v string) (*ProjectCreateV30AudienceDeviceType, error) {
	ev := ProjectCreateV30AudienceDeviceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30AudienceDeviceType: valid values are %v", v, AllowedProjectCreateV30AudienceDeviceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30AudienceDeviceType) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30AudienceDeviceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_audience_device_type value
func (v ProjectCreateV30AudienceDeviceType) Ptr() *ProjectCreateV30AudienceDeviceType {
	return &v
}
