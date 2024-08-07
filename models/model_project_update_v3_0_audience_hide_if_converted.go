/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectUpdateV30AudienceHideIfConverted
type ProjectUpdateV30AudienceHideIfConverted string

// List of project_update_v3.0_audience_hide_if_converted
const (
	ADVERTISER_ProjectUpdateV30AudienceHideIfConverted   ProjectUpdateV30AudienceHideIfConverted = "ADVERTISER"
	APP_ProjectUpdateV30AudienceHideIfConverted          ProjectUpdateV30AudienceHideIfConverted = "APP"
	CUSTOMER_ProjectUpdateV30AudienceHideIfConverted     ProjectUpdateV30AudienceHideIfConverted = "CUSTOMER"
	NO_EXCLUDE_ProjectUpdateV30AudienceHideIfConverted   ProjectUpdateV30AudienceHideIfConverted = "NO_EXCLUDE"
	ORGANIZATION_ProjectUpdateV30AudienceHideIfConverted ProjectUpdateV30AudienceHideIfConverted = "ORGANIZATION"
	PROJECT_ProjectUpdateV30AudienceHideIfConverted      ProjectUpdateV30AudienceHideIfConverted = "PROJECT"
	PROMOTION_ProjectUpdateV30AudienceHideIfConverted    ProjectUpdateV30AudienceHideIfConverted = "PROMOTION"
)

// All allowed values of ProjectUpdateV30AudienceHideIfConverted enum
var AllowedProjectUpdateV30AudienceHideIfConvertedEnumValues = []ProjectUpdateV30AudienceHideIfConverted{
	"ADVERTISER",
	"APP",
	"CUSTOMER",
	"NO_EXCLUDE",
	"ORGANIZATION",
	"PROJECT",
	"PROMOTION",
}

// NewProjectUpdateV30AudienceHideIfConvertedFromValue returns a pointer to a valid ProjectUpdateV30AudienceHideIfConverted
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30AudienceHideIfConvertedFromValue(v string) (*ProjectUpdateV30AudienceHideIfConverted, error) {
	ev := ProjectUpdateV30AudienceHideIfConverted(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30AudienceHideIfConverted: valid values are %v", v, AllowedProjectUpdateV30AudienceHideIfConvertedEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30AudienceHideIfConverted) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30AudienceHideIfConvertedEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_audience_hide_if_converted value
func (v ProjectUpdateV30AudienceHideIfConverted) Ptr() *ProjectUpdateV30AudienceHideIfConverted {
	return &v
}
