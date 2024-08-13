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

// ProjectUpdateV30AigcDynamicCreativeSwitch
type ProjectUpdateV30AigcDynamicCreativeSwitch string

// List of project_update_v3.0_aigc_dynamic_creative_switch
const (
	ON_ProjectUpdateV30AigcDynamicCreativeSwitch  ProjectUpdateV30AigcDynamicCreativeSwitch = "ON"
	OFF_ProjectUpdateV30AigcDynamicCreativeSwitch ProjectUpdateV30AigcDynamicCreativeSwitch = "OFF"
)

// All allowed values of ProjectUpdateV30AigcDynamicCreativeSwitch enum
var AllowedProjectUpdateV30AigcDynamicCreativeSwitchEnumValues = []ProjectUpdateV30AigcDynamicCreativeSwitch{
	"ON",
	"OFF",
}

// NewProjectUpdateV30AigcDynamicCreativeSwitchFromValue returns a pointer to a valid ProjectUpdateV30AigcDynamicCreativeSwitch
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30AigcDynamicCreativeSwitchFromValue(v string) (*ProjectUpdateV30AigcDynamicCreativeSwitch, error) {
	ev := ProjectUpdateV30AigcDynamicCreativeSwitch(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30AigcDynamicCreativeSwitch: valid values are %v", v, AllowedProjectUpdateV30AigcDynamicCreativeSwitchEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30AigcDynamicCreativeSwitch) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30AigcDynamicCreativeSwitchEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_aigc_dynamic_creative_switch value
func (v ProjectUpdateV30AigcDynamicCreativeSwitch) Ptr() *ProjectUpdateV30AigcDynamicCreativeSwitch {
	return &v
}
