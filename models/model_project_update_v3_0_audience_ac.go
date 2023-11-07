/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectUpdateV30AudienceAc
type ProjectUpdateV30AudienceAc string

// List of project_update_v3.0_audience_ac
const (
	Enum_2_G_ProjectUpdateV30AudienceAc ProjectUpdateV30AudienceAc = "2G"
	Enum_3_G_ProjectUpdateV30AudienceAc ProjectUpdateV30AudienceAc = "3G"
	Enum_4_G_ProjectUpdateV30AudienceAc ProjectUpdateV30AudienceAc = "4G"
	Enum_5_G_ProjectUpdateV30AudienceAc ProjectUpdateV30AudienceAc = "5G"
	WIFI_ProjectUpdateV30AudienceAc     ProjectUpdateV30AudienceAc = "WIFI"
)

// All allowed values of ProjectUpdateV30AudienceAc enum
var AllowedProjectUpdateV30AudienceAcEnumValues = []ProjectUpdateV30AudienceAc{
	"2G",
	"3G",
	"4G",
	"5G",
	"WIFI",
}

// NewProjectUpdateV30AudienceAcFromValue returns a pointer to a valid ProjectUpdateV30AudienceAc
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30AudienceAcFromValue(v string) (*ProjectUpdateV30AudienceAc, error) {
	ev := ProjectUpdateV30AudienceAc(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30AudienceAc: valid values are %v", v, AllowedProjectUpdateV30AudienceAcEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30AudienceAc) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30AudienceAcEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_audience_ac value
func (v ProjectUpdateV30AudienceAc) Ptr() *ProjectUpdateV30AudienceAc {
	return &v
}
