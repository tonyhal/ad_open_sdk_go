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

// ProjectUpdateV30AudienceDeviceType
type ProjectUpdateV30AudienceDeviceType string

// List of project_update_v3.0_audience_device_type
const (
	MOBILE_ProjectUpdateV30AudienceDeviceType ProjectUpdateV30AudienceDeviceType = "MOBILE"
	PAD_ProjectUpdateV30AudienceDeviceType    ProjectUpdateV30AudienceDeviceType = "PAD"
)

// All allowed values of ProjectUpdateV30AudienceDeviceType enum
var AllowedProjectUpdateV30AudienceDeviceTypeEnumValues = []ProjectUpdateV30AudienceDeviceType{
	"MOBILE",
	"PAD",
}

// NewProjectUpdateV30AudienceDeviceTypeFromValue returns a pointer to a valid ProjectUpdateV30AudienceDeviceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30AudienceDeviceTypeFromValue(v string) (*ProjectUpdateV30AudienceDeviceType, error) {
	ev := ProjectUpdateV30AudienceDeviceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30AudienceDeviceType: valid values are %v", v, AllowedProjectUpdateV30AudienceDeviceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30AudienceDeviceType) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30AudienceDeviceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_audience_device_type value
func (v ProjectUpdateV30AudienceDeviceType) Ptr() *ProjectUpdateV30AudienceDeviceType {
	return &v
}
