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

// ProjectUpdateV30KeywordsMatchType
type ProjectUpdateV30KeywordsMatchType string

// List of project_update_v3.0_keywords_match_type
const (
	EXTENSIVE_ProjectUpdateV30KeywordsMatchType ProjectUpdateV30KeywordsMatchType = "EXTENSIVE"
	PHRASE_ProjectUpdateV30KeywordsMatchType    ProjectUpdateV30KeywordsMatchType = "PHRASE"
	PRECISION_ProjectUpdateV30KeywordsMatchType ProjectUpdateV30KeywordsMatchType = "PRECISION"
)

// All allowed values of ProjectUpdateV30KeywordsMatchType enum
var AllowedProjectUpdateV30KeywordsMatchTypeEnumValues = []ProjectUpdateV30KeywordsMatchType{
	"EXTENSIVE",
	"PHRASE",
	"PRECISION",
}

// NewProjectUpdateV30KeywordsMatchTypeFromValue returns a pointer to a valid ProjectUpdateV30KeywordsMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30KeywordsMatchTypeFromValue(v string) (*ProjectUpdateV30KeywordsMatchType, error) {
	ev := ProjectUpdateV30KeywordsMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30KeywordsMatchType: valid values are %v", v, AllowedProjectUpdateV30KeywordsMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30KeywordsMatchType) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30KeywordsMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_keywords_match_type value
func (v ProjectUpdateV30KeywordsMatchType) Ptr() *ProjectUpdateV30KeywordsMatchType {
	return &v
}
