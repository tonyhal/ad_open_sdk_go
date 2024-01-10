/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCreateV30PromotionType
type ProjectCreateV30PromotionType string

// List of project_create_v3.0_promotion_type
const (
	AWEME_HOME_PAGE_ProjectCreateV30PromotionType   ProjectCreateV30PromotionType = "AWEME_HOME_PAGE"
	LANDING_PAGE_LINK_ProjectCreateV30PromotionType ProjectCreateV30PromotionType = "LANDING_PAGE_LINK"
)

// All allowed values of ProjectCreateV30PromotionType enum
var AllowedProjectCreateV30PromotionTypeEnumValues = []ProjectCreateV30PromotionType{
	"AWEME_HOME_PAGE",
	"LANDING_PAGE_LINK",
}

// NewProjectCreateV30PromotionTypeFromValue returns a pointer to a valid ProjectCreateV30PromotionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30PromotionTypeFromValue(v string) (*ProjectCreateV30PromotionType, error) {
	ev := ProjectCreateV30PromotionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30PromotionType: valid values are %v", v, AllowedProjectCreateV30PromotionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30PromotionType) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30PromotionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_promotion_type value
func (v ProjectCreateV30PromotionType) Ptr() *ProjectCreateV30PromotionType {
	return &v
}
