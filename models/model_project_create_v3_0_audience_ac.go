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

// ProjectCreateV30AudienceAc
type ProjectCreateV30AudienceAc string

// List of project_create_v3.0_audience_ac
const (
	Enum_2_G_ProjectCreateV30AudienceAc ProjectCreateV30AudienceAc = "2G"
	Enum_3_G_ProjectCreateV30AudienceAc ProjectCreateV30AudienceAc = "3G"
	Enum_4_G_ProjectCreateV30AudienceAc ProjectCreateV30AudienceAc = "4G"
	Enum_5_G_ProjectCreateV30AudienceAc ProjectCreateV30AudienceAc = "5G"
	WIFI_ProjectCreateV30AudienceAc     ProjectCreateV30AudienceAc = "WIFI"
)

// All allowed values of ProjectCreateV30AudienceAc enum
var AllowedProjectCreateV30AudienceAcEnumValues = []ProjectCreateV30AudienceAc{
	"2G",
	"3G",
	"4G",
	"5G",
	"WIFI",
}

// NewProjectCreateV30AudienceAcFromValue returns a pointer to a valid ProjectCreateV30AudienceAc
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30AudienceAcFromValue(v string) (*ProjectCreateV30AudienceAc, error) {
	ev := ProjectCreateV30AudienceAc(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30AudienceAc: valid values are %v", v, AllowedProjectCreateV30AudienceAcEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30AudienceAc) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30AudienceAcEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_audience_ac value
func (v ProjectCreateV30AudienceAc) Ptr() *ProjectCreateV30AudienceAc {
	return &v
}
