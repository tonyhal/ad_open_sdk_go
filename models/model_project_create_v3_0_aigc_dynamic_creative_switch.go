/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCreateV30AigcDynamicCreativeSwitch
type ProjectCreateV30AigcDynamicCreativeSwitch string

// List of project_create_v3.0_aigc_dynamic_creative_switch
const (
	ON_ProjectCreateV30AigcDynamicCreativeSwitch  ProjectCreateV30AigcDynamicCreativeSwitch = "ON"
	OFF_ProjectCreateV30AigcDynamicCreativeSwitch ProjectCreateV30AigcDynamicCreativeSwitch = "OFF"
)

// All allowed values of ProjectCreateV30AigcDynamicCreativeSwitch enum
var AllowedProjectCreateV30AigcDynamicCreativeSwitchEnumValues = []ProjectCreateV30AigcDynamicCreativeSwitch{
	"ON",
	"OFF",
}

// NewProjectCreateV30AigcDynamicCreativeSwitchFromValue returns a pointer to a valid ProjectCreateV30AigcDynamicCreativeSwitch
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30AigcDynamicCreativeSwitchFromValue(v string) (*ProjectCreateV30AigcDynamicCreativeSwitch, error) {
	ev := ProjectCreateV30AigcDynamicCreativeSwitch(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30AigcDynamicCreativeSwitch: valid values are %v", v, AllowedProjectCreateV30AigcDynamicCreativeSwitchEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30AigcDynamicCreativeSwitch) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30AigcDynamicCreativeSwitchEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_aigc_dynamic_creative_switch value
func (v ProjectCreateV30AigcDynamicCreativeSwitch) Ptr() *ProjectCreateV30AigcDynamicCreativeSwitch {
	return &v
}
