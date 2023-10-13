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

// ProjectUpdateV30AudienceFilterAwemeFansCount
type ProjectUpdateV30AudienceFilterAwemeFansCount int64

// List of project_update_v3.0_audience_filter_aweme_fans_count
const (
	Enum_0_ProjectUpdateV30AudienceFilterAwemeFansCount    ProjectUpdateV30AudienceFilterAwemeFansCount = 0
	Enum_1000_ProjectUpdateV30AudienceFilterAwemeFansCount ProjectUpdateV30AudienceFilterAwemeFansCount = 1000
	Enum_200_ProjectUpdateV30AudienceFilterAwemeFansCount  ProjectUpdateV30AudienceFilterAwemeFansCount = 200
	Enum_500_ProjectUpdateV30AudienceFilterAwemeFansCount  ProjectUpdateV30AudienceFilterAwemeFansCount = 500
)

// All allowed values of ProjectUpdateV30AudienceFilterAwemeFansCount enum
var AllowedProjectUpdateV30AudienceFilterAwemeFansCountEnumValues = []ProjectUpdateV30AudienceFilterAwemeFansCount{
	0,
	1000,
	200,
	500,
}

// NewProjectUpdateV30AudienceFilterAwemeFansCountFromValue returns a pointer to a valid ProjectUpdateV30AudienceFilterAwemeFansCount
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30AudienceFilterAwemeFansCountFromValue(v int64) (*ProjectUpdateV30AudienceFilterAwemeFansCount, error) {
	ev := ProjectUpdateV30AudienceFilterAwemeFansCount(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30AudienceFilterAwemeFansCount: valid values are %v", v, AllowedProjectUpdateV30AudienceFilterAwemeFansCountEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30AudienceFilterAwemeFansCount) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30AudienceFilterAwemeFansCountEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_audience_filter_aweme_fans_count value
func (v ProjectUpdateV30AudienceFilterAwemeFansCount) Ptr() *ProjectUpdateV30AudienceFilterAwemeFansCount {
	return &v
}
