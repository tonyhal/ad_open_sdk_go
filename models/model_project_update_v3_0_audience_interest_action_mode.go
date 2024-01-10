/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectUpdateV30AudienceInterestActionMode
type ProjectUpdateV30AudienceInterestActionMode string

// List of project_update_v3.0_audience_interest_action_mode
const (
	CUSTOM_ProjectUpdateV30AudienceInterestActionMode    ProjectUpdateV30AudienceInterestActionMode = "CUSTOM"
	RECOMMEND_ProjectUpdateV30AudienceInterestActionMode ProjectUpdateV30AudienceInterestActionMode = "RECOMMEND"
	UNLIMITED_ProjectUpdateV30AudienceInterestActionMode ProjectUpdateV30AudienceInterestActionMode = "UNLIMITED"
)

// All allowed values of ProjectUpdateV30AudienceInterestActionMode enum
var AllowedProjectUpdateV30AudienceInterestActionModeEnumValues = []ProjectUpdateV30AudienceInterestActionMode{
	"CUSTOM",
	"RECOMMEND",
	"UNLIMITED",
}

// NewProjectUpdateV30AudienceInterestActionModeFromValue returns a pointer to a valid ProjectUpdateV30AudienceInterestActionMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30AudienceInterestActionModeFromValue(v string) (*ProjectUpdateV30AudienceInterestActionMode, error) {
	ev := ProjectUpdateV30AudienceInterestActionMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30AudienceInterestActionMode: valid values are %v", v, AllowedProjectUpdateV30AudienceInterestActionModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30AudienceInterestActionMode) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30AudienceInterestActionModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_audience_interest_action_mode value
func (v ProjectUpdateV30AudienceInterestActionMode) Ptr() *ProjectUpdateV30AudienceInterestActionMode {
	return &v
}
