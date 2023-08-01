/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectUpdateV30KeywordsBidType
type ProjectUpdateV30KeywordsBidType string

// List of project_update_v3.0_keywords_bid_type
const (
	FEED_TO_SEARCH_ProjectUpdateV30KeywordsBidType ProjectUpdateV30KeywordsBidType = "FEED_TO_SEARCH"
)

// All allowed values of ProjectUpdateV30KeywordsBidType enum
var AllowedProjectUpdateV30KeywordsBidTypeEnumValues = []ProjectUpdateV30KeywordsBidType{
	"FEED_TO_SEARCH",
}

// NewProjectUpdateV30KeywordsBidTypeFromValue returns a pointer to a valid ProjectUpdateV30KeywordsBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30KeywordsBidTypeFromValue(v string) (*ProjectUpdateV30KeywordsBidType, error) {
	ev := ProjectUpdateV30KeywordsBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30KeywordsBidType: valid values are %v", v, AllowedProjectUpdateV30KeywordsBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30KeywordsBidType) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30KeywordsBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_keywords_bid_type value
func (v ProjectUpdateV30KeywordsBidType) Ptr() *ProjectUpdateV30KeywordsBidType {
	return &v
}