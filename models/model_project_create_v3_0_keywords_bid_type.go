/*
API Title

巨量引擎开放平台

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCreateV30KeywordsBidType
type ProjectCreateV30KeywordsBidType string

// List of project_create_v3.0_keywords_bid_type
const (
	FEED_TO_SEARCH_ProjectCreateV30KeywordsBidType ProjectCreateV30KeywordsBidType = "FEED_TO_SEARCH"
)

// All allowed values of ProjectCreateV30KeywordsBidType enum
var AllowedProjectCreateV30KeywordsBidTypeEnumValues = []ProjectCreateV30KeywordsBidType{
	"FEED_TO_SEARCH",
}

// NewProjectCreateV30KeywordsBidTypeFromValue returns a pointer to a valid ProjectCreateV30KeywordsBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30KeywordsBidTypeFromValue(v string) (*ProjectCreateV30KeywordsBidType, error) {
	ev := ProjectCreateV30KeywordsBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30KeywordsBidType: valid values are %v", v, AllowedProjectCreateV30KeywordsBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30KeywordsBidType) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30KeywordsBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_keywords_bid_type value
func (v ProjectCreateV30KeywordsBidType) Ptr() *ProjectCreateV30KeywordsBidType {
	return &v
}
