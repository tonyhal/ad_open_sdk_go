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

// ProjectListV30DataListOptStatus
type ProjectListV30DataListOptStatus string

// List of project_list_v3.0_data_list_opt_status
const (
	DISABLE_ProjectListV30DataListOptStatus ProjectListV30DataListOptStatus = "DISABLE"
	ENABLE_ProjectListV30DataListOptStatus  ProjectListV30DataListOptStatus = "ENABLE"
)

// All allowed values of ProjectListV30DataListOptStatus enum
var AllowedProjectListV30DataListOptStatusEnumValues = []ProjectListV30DataListOptStatus{
	"DISABLE",
	"ENABLE",
}

// NewProjectListV30DataListOptStatusFromValue returns a pointer to a valid ProjectListV30DataListOptStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListOptStatusFromValue(v string) (*ProjectListV30DataListOptStatus, error) {
	ev := ProjectListV30DataListOptStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListOptStatus: valid values are %v", v, AllowedProjectListV30DataListOptStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListOptStatus) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListOptStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_opt_status value
func (v ProjectListV30DataListOptStatus) Ptr() *ProjectListV30DataListOptStatus {
	return &v
}
