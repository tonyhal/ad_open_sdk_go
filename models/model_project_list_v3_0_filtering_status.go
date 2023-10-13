/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30FilteringStatus
type ProjectListV30FilteringStatus string

// List of project_list_v3.0_filtering_status
const (
	PROJECT_STATUS_ALL_ProjectListV30FilteringStatus                       ProjectListV30FilteringStatus = "PROJECT_STATUS_ALL"
	PROJECT_STATUS_BUDGET_EXCEED_ProjectListV30FilteringStatus             ProjectListV30FilteringStatus = "PROJECT_STATUS_BUDGET_EXCEED"
	PROJECT_STATUS_BUDGET_PRE_OFFLINE_BUDGET_ProjectListV30FilteringStatus ProjectListV30FilteringStatus = "PROJECT_STATUS_BUDGET_PRE_OFFLINE_BUDGET"
	PROJECT_STATUS_DELETE_ProjectListV30FilteringStatus                    ProjectListV30FilteringStatus = "PROJECT_STATUS_DELETE"
	PROJECT_STATUS_DISABLE_ProjectListV30FilteringStatus                   ProjectListV30FilteringStatus = "PROJECT_STATUS_DISABLE"
	PROJECT_STATUS_DONE_ProjectListV30FilteringStatus                      ProjectListV30FilteringStatus = "PROJECT_STATUS_DONE"
	PROJECT_STATUS_ENABLE_ProjectListV30FilteringStatus                    ProjectListV30FilteringStatus = "PROJECT_STATUS_ENABLE"
	PROJECT_STATUS_NOT_DELETE_ProjectListV30FilteringStatus                ProjectListV30FilteringStatus = "PROJECT_STATUS_NOT_DELETE"
	PROJECT_STATUS_NOT_START_ProjectListV30FilteringStatus                 ProjectListV30FilteringStatus = "PROJECT_STATUS_NOT_START"
	PROJECT_STATUS_NO_SCHEDULE_ProjectListV30FilteringStatus               ProjectListV30FilteringStatus = "PROJECT_STATUS_NO_SCHEDULE"
)

// All allowed values of ProjectListV30FilteringStatus enum
var AllowedProjectListV30FilteringStatusEnumValues = []ProjectListV30FilteringStatus{
	"PROJECT_STATUS_ALL",
	"PROJECT_STATUS_BUDGET_EXCEED",
	"PROJECT_STATUS_BUDGET_PRE_OFFLINE_BUDGET",
	"PROJECT_STATUS_DELETE",
	"PROJECT_STATUS_DISABLE",
	"PROJECT_STATUS_DONE",
	"PROJECT_STATUS_ENABLE",
	"PROJECT_STATUS_NOT_DELETE",
	"PROJECT_STATUS_NOT_START",
	"PROJECT_STATUS_NO_SCHEDULE",
}

// NewProjectListV30FilteringStatusFromValue returns a pointer to a valid ProjectListV30FilteringStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30FilteringStatusFromValue(v string) (*ProjectListV30FilteringStatus, error) {
	ev := ProjectListV30FilteringStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30FilteringStatus: valid values are %v", v, AllowedProjectListV30FilteringStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30FilteringStatus) IsValid() bool {
	for _, existing := range AllowedProjectListV30FilteringStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_filtering_status value
func (v ProjectListV30FilteringStatus) Ptr() *ProjectListV30FilteringStatus {
	return &v
}
