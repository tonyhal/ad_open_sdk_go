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

// ProjectCreateV30AudienceHideIfExists
type ProjectCreateV30AudienceHideIfExists string

// List of project_create_v3.0_audience_hide_if_exists
const (
	FILTER_ProjectCreateV30AudienceHideIfExists    ProjectCreateV30AudienceHideIfExists = "FILTER"
	TARGETING_ProjectCreateV30AudienceHideIfExists ProjectCreateV30AudienceHideIfExists = "TARGETING"
	UNLIMITED_ProjectCreateV30AudienceHideIfExists ProjectCreateV30AudienceHideIfExists = "UNLIMITED"
)

// All allowed values of ProjectCreateV30AudienceHideIfExists enum
var AllowedProjectCreateV30AudienceHideIfExistsEnumValues = []ProjectCreateV30AudienceHideIfExists{
	"FILTER",
	"TARGETING",
	"UNLIMITED",
}

// NewProjectCreateV30AudienceHideIfExistsFromValue returns a pointer to a valid ProjectCreateV30AudienceHideIfExists
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30AudienceHideIfExistsFromValue(v string) (*ProjectCreateV30AudienceHideIfExists, error) {
	ev := ProjectCreateV30AudienceHideIfExists(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30AudienceHideIfExists: valid values are %v", v, AllowedProjectCreateV30AudienceHideIfExistsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30AudienceHideIfExists) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30AudienceHideIfExistsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_audience_hide_if_exists value
func (v ProjectCreateV30AudienceHideIfExists) Ptr() *ProjectCreateV30AudienceHideIfExists {
	return &v
}
