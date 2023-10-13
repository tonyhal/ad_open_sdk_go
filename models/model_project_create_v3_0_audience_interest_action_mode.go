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

// ProjectCreateV30AudienceInterestActionMode
type ProjectCreateV30AudienceInterestActionMode string

// List of project_create_v3.0_audience_interest_action_mode
const (
	CUSTOM_ProjectCreateV30AudienceInterestActionMode    ProjectCreateV30AudienceInterestActionMode = "CUSTOM"
	RECOMMEND_ProjectCreateV30AudienceInterestActionMode ProjectCreateV30AudienceInterestActionMode = "RECOMMEND"
	UNLIMITED_ProjectCreateV30AudienceInterestActionMode ProjectCreateV30AudienceInterestActionMode = "UNLIMITED"
)

// All allowed values of ProjectCreateV30AudienceInterestActionMode enum
var AllowedProjectCreateV30AudienceInterestActionModeEnumValues = []ProjectCreateV30AudienceInterestActionMode{
	"CUSTOM",
	"RECOMMEND",
	"UNLIMITED",
}

// NewProjectCreateV30AudienceInterestActionModeFromValue returns a pointer to a valid ProjectCreateV30AudienceInterestActionMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30AudienceInterestActionModeFromValue(v string) (*ProjectCreateV30AudienceInterestActionMode, error) {
	ev := ProjectCreateV30AudienceInterestActionMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30AudienceInterestActionMode: valid values are %v", v, AllowedProjectCreateV30AudienceInterestActionModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30AudienceInterestActionMode) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30AudienceInterestActionModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_audience_interest_action_mode value
func (v ProjectCreateV30AudienceInterestActionMode) Ptr() *ProjectCreateV30AudienceInterestActionMode {
	return &v
}
