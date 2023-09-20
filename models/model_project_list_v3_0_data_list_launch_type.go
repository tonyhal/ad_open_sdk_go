/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListLaunchType
type ProjectListV30DataListLaunchType string

// List of project_list_v3.0_data_list_launch_type
const (
	DIRECT_OPEN_ProjectListV30DataListLaunchType   ProjectListV30DataListLaunchType = "DIRECT_OPEN"
	EXTERNAL_OPEN_ProjectListV30DataListLaunchType ProjectListV30DataListLaunchType = "EXTERNAL_OPEN"
)

// All allowed values of ProjectListV30DataListLaunchType enum
var AllowedProjectListV30DataListLaunchTypeEnumValues = []ProjectListV30DataListLaunchType{
	"DIRECT_OPEN",
	"EXTERNAL_OPEN",
}

// NewProjectListV30DataListLaunchTypeFromValue returns a pointer to a valid ProjectListV30DataListLaunchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListLaunchTypeFromValue(v string) (*ProjectListV30DataListLaunchType, error) {
	ev := ProjectListV30DataListLaunchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListLaunchType: valid values are %v", v, AllowedProjectListV30DataListLaunchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListLaunchType) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListLaunchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_launch_type value
func (v ProjectListV30DataListLaunchType) Ptr() *ProjectListV30DataListLaunchType {
	return &v
}
