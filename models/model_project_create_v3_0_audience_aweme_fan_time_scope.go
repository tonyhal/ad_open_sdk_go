/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCreateV30AudienceAwemeFanTimeScope
type ProjectCreateV30AudienceAwemeFanTimeScope string

// List of project_create_v3.0_audience_aweme_fan_time_scope
const (
	FIFTEEN_DAYS_ProjectCreateV30AudienceAwemeFanTimeScope ProjectCreateV30AudienceAwemeFanTimeScope = "FIFTEEN_DAYS"
	SIXTY_DAYS_ProjectCreateV30AudienceAwemeFanTimeScope   ProjectCreateV30AudienceAwemeFanTimeScope = "SIXTY_DAYS"
	THIRTY_DAYS_ProjectCreateV30AudienceAwemeFanTimeScope  ProjectCreateV30AudienceAwemeFanTimeScope = "THIRTY_DAYS"
)

// All allowed values of ProjectCreateV30AudienceAwemeFanTimeScope enum
var AllowedProjectCreateV30AudienceAwemeFanTimeScopeEnumValues = []ProjectCreateV30AudienceAwemeFanTimeScope{
	"FIFTEEN_DAYS",
	"SIXTY_DAYS",
	"THIRTY_DAYS",
}

// NewProjectCreateV30AudienceAwemeFanTimeScopeFromValue returns a pointer to a valid ProjectCreateV30AudienceAwemeFanTimeScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30AudienceAwemeFanTimeScopeFromValue(v string) (*ProjectCreateV30AudienceAwemeFanTimeScope, error) {
	ev := ProjectCreateV30AudienceAwemeFanTimeScope(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30AudienceAwemeFanTimeScope: valid values are %v", v, AllowedProjectCreateV30AudienceAwemeFanTimeScopeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30AudienceAwemeFanTimeScope) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30AudienceAwemeFanTimeScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_audience_aweme_fan_time_scope value
func (v ProjectCreateV30AudienceAwemeFanTimeScope) Ptr() *ProjectCreateV30AudienceAwemeFanTimeScope {
	return &v
}
