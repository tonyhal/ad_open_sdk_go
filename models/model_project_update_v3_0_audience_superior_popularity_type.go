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

// ProjectUpdateV30AudienceSuperiorPopularityType
type ProjectUpdateV30AudienceSuperiorPopularityType string

// List of project_update_v3.0_audience_superior_popularity_type
const (
	GAME_ProjectUpdateV30AudienceSuperiorPopularityType ProjectUpdateV30AudienceSuperiorPopularityType = "GAME"
	NONE_ProjectUpdateV30AudienceSuperiorPopularityType ProjectUpdateV30AudienceSuperiorPopularityType = "NONE"
)

// All allowed values of ProjectUpdateV30AudienceSuperiorPopularityType enum
var AllowedProjectUpdateV30AudienceSuperiorPopularityTypeEnumValues = []ProjectUpdateV30AudienceSuperiorPopularityType{
	"GAME",
	"NONE",
}

// NewProjectUpdateV30AudienceSuperiorPopularityTypeFromValue returns a pointer to a valid ProjectUpdateV30AudienceSuperiorPopularityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30AudienceSuperiorPopularityTypeFromValue(v string) (*ProjectUpdateV30AudienceSuperiorPopularityType, error) {
	ev := ProjectUpdateV30AudienceSuperiorPopularityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30AudienceSuperiorPopularityType: valid values are %v", v, AllowedProjectUpdateV30AudienceSuperiorPopularityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30AudienceSuperiorPopularityType) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30AudienceSuperiorPopularityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_audience_superior_popularity_type value
func (v ProjectUpdateV30AudienceSuperiorPopularityType) Ptr() *ProjectUpdateV30AudienceSuperiorPopularityType {
	return &v
}
