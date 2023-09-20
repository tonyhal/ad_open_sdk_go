/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectUpdateV30AudienceLocationType
type ProjectUpdateV30AudienceLocationType string

// List of project_update_v3.0_audience_location_type
const (
	ALL_ProjectUpdateV30AudienceLocationType     ProjectUpdateV30AudienceLocationType = "ALL"
	CURRENT_ProjectUpdateV30AudienceLocationType ProjectUpdateV30AudienceLocationType = "CURRENT"
	HOME_ProjectUpdateV30AudienceLocationType    ProjectUpdateV30AudienceLocationType = "HOME"
	TRAVEL_ProjectUpdateV30AudienceLocationType  ProjectUpdateV30AudienceLocationType = "TRAVEL"
)

// All allowed values of ProjectUpdateV30AudienceLocationType enum
var AllowedProjectUpdateV30AudienceLocationTypeEnumValues = []ProjectUpdateV30AudienceLocationType{
	"ALL",
	"CURRENT",
	"HOME",
	"TRAVEL",
}

// NewProjectUpdateV30AudienceLocationTypeFromValue returns a pointer to a valid ProjectUpdateV30AudienceLocationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30AudienceLocationTypeFromValue(v string) (*ProjectUpdateV30AudienceLocationType, error) {
	ev := ProjectUpdateV30AudienceLocationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30AudienceLocationType: valid values are %v", v, AllowedProjectUpdateV30AudienceLocationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30AudienceLocationType) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30AudienceLocationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_audience_location_type value
func (v ProjectUpdateV30AudienceLocationType) Ptr() *ProjectUpdateV30AudienceLocationType {
	return &v
}
