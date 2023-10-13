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

// ProjectUpdateV30AudienceRegionRecommend
type ProjectUpdateV30AudienceRegionRecommend string

// List of project_update_v3.0_audience_region_recommend
const (
	OFF_ProjectUpdateV30AudienceRegionRecommend ProjectUpdateV30AudienceRegionRecommend = "OFF"
	ON_ProjectUpdateV30AudienceRegionRecommend  ProjectUpdateV30AudienceRegionRecommend = "ON"
)

// All allowed values of ProjectUpdateV30AudienceRegionRecommend enum
var AllowedProjectUpdateV30AudienceRegionRecommendEnumValues = []ProjectUpdateV30AudienceRegionRecommend{
	"OFF",
	"ON",
}

// NewProjectUpdateV30AudienceRegionRecommendFromValue returns a pointer to a valid ProjectUpdateV30AudienceRegionRecommend
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30AudienceRegionRecommendFromValue(v string) (*ProjectUpdateV30AudienceRegionRecommend, error) {
	ev := ProjectUpdateV30AudienceRegionRecommend(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30AudienceRegionRecommend: valid values are %v", v, AllowedProjectUpdateV30AudienceRegionRecommendEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30AudienceRegionRecommend) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30AudienceRegionRecommendEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_audience_region_recommend value
func (v ProjectUpdateV30AudienceRegionRecommend) Ptr() *ProjectUpdateV30AudienceRegionRecommend {
	return &v
}
