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

// ProjectListV30DataListAudienceActivateType
type ProjectListV30DataListAudienceActivateType string

// List of project_list_v3.0_data_list_audience_activate_type
const (
	ONE_MONTH_2_THREE_MONTH_ProjectListV30DataListAudienceActivateType ProjectListV30DataListAudienceActivateType = "ONE_MONTH_2_THREE_MONTH"
	THREE_MONTH_EAILIER_ProjectListV30DataListAudienceActivateType     ProjectListV30DataListAudienceActivateType = "THREE_MONTH_EAILIER"
	UNLIMITED_ProjectListV30DataListAudienceActivateType               ProjectListV30DataListAudienceActivateType = "UNLIMITED"
	WITH_IN_A_MONTH_ProjectListV30DataListAudienceActivateType         ProjectListV30DataListAudienceActivateType = "WITH_IN_A_MONTH"
)

// All allowed values of ProjectListV30DataListAudienceActivateType enum
var AllowedProjectListV30DataListAudienceActivateTypeEnumValues = []ProjectListV30DataListAudienceActivateType{
	"ONE_MONTH_2_THREE_MONTH",
	"THREE_MONTH_EAILIER",
	"UNLIMITED",
	"WITH_IN_A_MONTH",
}

// NewProjectListV30DataListAudienceActivateTypeFromValue returns a pointer to a valid ProjectListV30DataListAudienceActivateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListAudienceActivateTypeFromValue(v string) (*ProjectListV30DataListAudienceActivateType, error) {
	ev := ProjectListV30DataListAudienceActivateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListAudienceActivateType: valid values are %v", v, AllowedProjectListV30DataListAudienceActivateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListAudienceActivateType) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListAudienceActivateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_audience_activate_type value
func (v ProjectListV30DataListAudienceActivateType) Ptr() *ProjectListV30DataListAudienceActivateType {
	return &v
}
