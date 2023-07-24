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

// ProjectUpdateV30DataErrorListObjectType
type ProjectUpdateV30DataErrorListObjectType string

// List of project_update_v3.0_data_error_list_object_type
const (
	BUDGET_ProjectUpdateV30DataErrorListObjectType             ProjectUpdateV30DataErrorListObjectType = "BUDGET"
	PROJECT_SETTING_ProjectUpdateV30DataErrorListObjectType    ProjectUpdateV30DataErrorListObjectType = "PROJECT_SETTING"
	PROMOTION_BUDGET_ProjectUpdateV30DataErrorListObjectType   ProjectUpdateV30DataErrorListObjectType = "PROMOTION_BUDGET"
	PROMOTION_MATERIAL_ProjectUpdateV30DataErrorListObjectType ProjectUpdateV30DataErrorListObjectType = "PROMOTION_MATERIAL"
	PROMOTION_SETTING_ProjectUpdateV30DataErrorListObjectType  ProjectUpdateV30DataErrorListObjectType = "PROMOTION_SETTING"
)

// All allowed values of ProjectUpdateV30DataErrorListObjectType enum
var AllowedProjectUpdateV30DataErrorListObjectTypeEnumValues = []ProjectUpdateV30DataErrorListObjectType{
	"BUDGET",
	"PROJECT_SETTING",
	"PROMOTION_BUDGET",
	"PROMOTION_MATERIAL",
	"PROMOTION_SETTING",
}

// NewProjectUpdateV30DataErrorListObjectTypeFromValue returns a pointer to a valid ProjectUpdateV30DataErrorListObjectType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30DataErrorListObjectTypeFromValue(v string) (*ProjectUpdateV30DataErrorListObjectType, error) {
	ev := ProjectUpdateV30DataErrorListObjectType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30DataErrorListObjectType: valid values are %v", v, AllowedProjectUpdateV30DataErrorListObjectTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30DataErrorListObjectType) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30DataErrorListObjectTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_data_error_list_object_type value
func (v ProjectUpdateV30DataErrorListObjectType) Ptr() *ProjectUpdateV30DataErrorListObjectType {
	return &v
}
