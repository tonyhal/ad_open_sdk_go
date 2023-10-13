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

// ProjectListV30DataListAudienceDpaLbs
type ProjectListV30DataListAudienceDpaLbs string

// List of project_list_v3.0_data_list_audience_dpa_lbs
const (
	OFF_ProjectListV30DataListAudienceDpaLbs ProjectListV30DataListAudienceDpaLbs = "OFF"
	ON_ProjectListV30DataListAudienceDpaLbs  ProjectListV30DataListAudienceDpaLbs = "ON"
)

// All allowed values of ProjectListV30DataListAudienceDpaLbs enum
var AllowedProjectListV30DataListAudienceDpaLbsEnumValues = []ProjectListV30DataListAudienceDpaLbs{
	"OFF",
	"ON",
}

// NewProjectListV30DataListAudienceDpaLbsFromValue returns a pointer to a valid ProjectListV30DataListAudienceDpaLbs
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListAudienceDpaLbsFromValue(v string) (*ProjectListV30DataListAudienceDpaLbs, error) {
	ev := ProjectListV30DataListAudienceDpaLbs(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListAudienceDpaLbs: valid values are %v", v, AllowedProjectListV30DataListAudienceDpaLbsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListAudienceDpaLbs) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListAudienceDpaLbsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_audience_dpa_lbs value
func (v ProjectListV30DataListAudienceDpaLbs) Ptr() *ProjectListV30DataListAudienceDpaLbs {
	return &v
}
