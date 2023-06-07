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

// ProjectUpdateV30AudienceFilterAwemeAbnormalActive
type ProjectUpdateV30AudienceFilterAwemeAbnormalActive string

// List of project_update_v3.0_audience_filter_aweme_abnormal_active
const (
	OFF_ProjectUpdateV30AudienceFilterAwemeAbnormalActive ProjectUpdateV30AudienceFilterAwemeAbnormalActive = "OFF"
	ON_ProjectUpdateV30AudienceFilterAwemeAbnormalActive  ProjectUpdateV30AudienceFilterAwemeAbnormalActive = "ON"
)

// All allowed values of ProjectUpdateV30AudienceFilterAwemeAbnormalActive enum
var AllowedProjectUpdateV30AudienceFilterAwemeAbnormalActiveEnumValues = []ProjectUpdateV30AudienceFilterAwemeAbnormalActive{
	"OFF",
	"ON",
}

// NewProjectUpdateV30AudienceFilterAwemeAbnormalActiveFromValue returns a pointer to a valid ProjectUpdateV30AudienceFilterAwemeAbnormalActive
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30AudienceFilterAwemeAbnormalActiveFromValue(v string) (*ProjectUpdateV30AudienceFilterAwemeAbnormalActive, error) {
	ev := ProjectUpdateV30AudienceFilterAwemeAbnormalActive(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30AudienceFilterAwemeAbnormalActive: valid values are %v", v, AllowedProjectUpdateV30AudienceFilterAwemeAbnormalActiveEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30AudienceFilterAwemeAbnormalActive) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30AudienceFilterAwemeAbnormalActiveEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_audience_filter_aweme_abnormal_active value
func (v ProjectUpdateV30AudienceFilterAwemeAbnormalActive) Ptr() *ProjectUpdateV30AudienceFilterAwemeAbnormalActive {
	return &v
}
