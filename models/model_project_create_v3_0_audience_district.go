/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCreateV30AudienceDistrict
type ProjectCreateV30AudienceDistrict string

// List of project_create_v3.0_audience_district
const (
	BUSINESS_DISTRICT_ProjectCreateV30AudienceDistrict ProjectCreateV30AudienceDistrict = "BUSINESS_DISTRICT"
	NONE_ProjectCreateV30AudienceDistrict              ProjectCreateV30AudienceDistrict = "NONE"
	OVERSEA_ProjectCreateV30AudienceDistrict           ProjectCreateV30AudienceDistrict = "OVERSEA"
	REGION_ProjectCreateV30AudienceDistrict            ProjectCreateV30AudienceDistrict = "REGION"
)

// All allowed values of ProjectCreateV30AudienceDistrict enum
var AllowedProjectCreateV30AudienceDistrictEnumValues = []ProjectCreateV30AudienceDistrict{
	"BUSINESS_DISTRICT",
	"NONE",
	"OVERSEA",
	"REGION",
}

// NewProjectCreateV30AudienceDistrictFromValue returns a pointer to a valid ProjectCreateV30AudienceDistrict
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30AudienceDistrictFromValue(v string) (*ProjectCreateV30AudienceDistrict, error) {
	ev := ProjectCreateV30AudienceDistrict(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30AudienceDistrict: valid values are %v", v, AllowedProjectCreateV30AudienceDistrictEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30AudienceDistrict) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30AudienceDistrictEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_audience_district value
func (v ProjectCreateV30AudienceDistrict) Ptr() *ProjectCreateV30AudienceDistrict {
	return &v
}
